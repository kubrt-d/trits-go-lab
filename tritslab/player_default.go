package tritslab

import (
	"math"
)

type TritsPlayerDumb struct {
	TritsPlayerDefault
}

type TritsPlayerDefault struct {
	Addr         TritsAddress // Player's address
	Player_type  string       // Player type ("dumb", "zion" etc...)
	banker       TritsBanker  // Connection to the banker
	started_with uint64       // Initail pocket before joining the game

}

// Set started_with
func (p *TritsPlayerDefault) SetStartedWith(started_with uint64) {
	p.started_with = started_with
}

// Choose the table
func (p *TritsPlayerDefault) ChooseTable() int8 {
	b := RandByte()
	return int8(b % 23)
}

// Choose amount
func (p *TritsPlayerDefault) ChooseAmount() uint64 {
	pocket := p.Balance()
	if NOMINAL > 0 && NOMINAL < pocket {
		return NOMINAL
	} else {
		b := RandByte()
		pocket := p.Balance()
		percent := int8(b%3 + 1) // Bet no more between 1 - 3 percent of the pocket
		return uint64(math.Round(float64(percent) / 100 * float64(pocket)))
	}
}

// Get the balance
func (p *TritsPlayerDefault) Balance() uint64 {
	return p.banker.Tell(p.Addr)
}

//Get the Address
func (p *TritsPlayerDefault) GetAddr() TritsAddress {
	return p.Addr
}

//Get Player type
func (p *TritsPlayerDefault) GetPlayerType() string {
	return p.Player_type
}

//Recharge if necessary (only zion players can do this)
func (p *TritsPlayerDefault) Recharge() uint64 {
	return 0
}


// Borrow money if necessary
func (p *TritsPlayerDefault) Borrow(max_borrow uint64) uint64 {
	pocket := p.Balance()
	if max_borrow >= 100*pocket {
		return max_borrow
	} else {
		return 0
	}
}

// Take profit as pleased
func (p *TritsPlayerDefault) TakeProfit() uint64 {
	pocket := p.Balance()
	if pocket < p.started_with*PROFIT_TRESHOLD { // Have not won enough money yet
		return 0
	} else {
		r := RandByte()
		if r%100 < LEAVE_GAME_PROB { // 0 means never leave the game
			return pocket - p.started_with
		}
	}
	return 0
}

// Bet or not - some strategy should be implemented here
func (p *TritsPlayerDefault) Bet(desk []*TritsGame) []TritsPlayerResponse {
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
		b := RandByte()
		x := int8(b%3 + 1) // Throw 1-3 coins in one go
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

	// No good inbalance found, bet on the first one I can afford
	var k int = 0
	for k < GAMES_ON_TABLE {
		if desk[k].Nominal < pocket {
			break
		}
		k++
	}
	if k < GAMES_ON_TABLE { // Found something to bet on
		res := NewTritsPlayerResponse(desk[k], desk[k].Nominal, p.Addr)
		responses = append(responses, res)
		if TD {
			l(LOG_NOTICE, LogName(p.Addr), " has ", pocket, " and decides to play on ", LogName(desk[k].ThisGame), ", with ", desk[k].Nominal)
		}
		return responses
	}

	if TD {
		l(LOG_NOTICE, LogName(p.Addr), " has got ", pocket, " only and can't afford any game")
	}

	return nil
}

// Human name for logging
func (p *TritsPlayerDefault) Name() string {
	var name string = ""
	switch p.Addr.Raw() {
	case LenderAddr:
		name = "Lender"
	case BankAddr:
		name = "Bank"
	case NeoAddr:
		name = "Neo"
	case TrinityAddr:
		name = "Trinity"
	case AgentAddr:
		name = "Agent"
	case KeymakerAddr:
		name = "Keymaker"
	case MorpheusAddr:
		name = "Morpheus"
	case NiobeAddr:
		name = "Niobe"
	case OracleAddr:
		name = "Oracle"
	case PersephoneAddr:
		name = "Persephone"
	case TwinsAddr:
		name = "Twins"
	case BugsAddr:
		name = "Bugs"
	case AnalystAddr:
		name = "Analyst"
	case SeraphAddr:
		name = "Seraph"
	case ArchitectAddr:
		name = "Architect"
	case BaneAddr:
		name = "Bane"
	default:
		name = p.Addr.Human()
	}
	return name
}
