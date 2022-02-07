package tritslab

type TritsPlayerAgent struct {
	TritsPlayerDefault
}

// Only initaites or throws 5 coins on inbalance 2

func (p *TritsPlayerAgent) Bet(desk []*TritsGame) []TritsPlayerResponse {
	var responses []TritsPlayerResponse
	pocket := p.Balance()
	if pocket == 0 { // No money no honey
		return nil
	}

	// Try to find an uninitiated game
	var i int = 0
	for i < GAMES_ON_TABLE && desk[i].Nominal != 0 {
		i++
	}
	if i < GAMES_ON_TABLE { // Found an uninitiated game
		amount := p.ChooseAmount()
		if amount == 0 { // This guy is out of money
			if TD {
				l(LOG_NOTICE, LogName(p.GetAddr()), " is broke !")
			}
			return nil
		}
		res := NewTritsPlayerResponse(desk[i], amount, p.Addr)
		responses = append(responses, res)
		if TD {
			l(LOG_NOTICE, LogName(p.Addr), " decides to init new game ", LogName(desk[i].ThisGame), ", with ", amount)
		}
		return responses
	}

	// No uninitiated game found, try to find the inbalance of 2 which I can afford
	var j int = 0
	for j < GAMES_ON_TABLE { // Loop through all games
		if (desk[j].GetInbalance() == 2) && (desk[j].Nominal < pocket) {
			break // If inbalance is 2 and the player can afford it
		}
		j++
	}
	if j < GAMES_ON_TABLE { // Found a good inbalance, bet
		x := 5 // Throw 5 coins in one go
		if pocket < uint64(x)*desk[j].Nominal {
			x = 1
		}
		res := NewTritsPlayerResponse(desk[j], uint64(x)*desk[j].Nominal, p.Addr)
		responses = append(responses, res)
		if TD {
			l(LOG_NOTICE, LogName(p.Addr), " has ", pocket, " and finds inbalance 2 ", LogName(desk[j].ThisGame), ", nominal ", desk[j].Nominal)
		}
		return responses
	}

	return nil
}
