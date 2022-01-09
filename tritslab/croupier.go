package tritslab

//"math/rand"
//"time"

type TritsCroupier struct {
	Table      *TritsTable  // Croupier's table
	Players    *TritsSquad  // Players
	Banker     *TritsBanker // Banker
	max_borrow uint64       // Max amount which can be borrowed
}

// Croupier starts his day, with a certain lot for players and something to give to the bank
func NewTritsCroupier(lot uint64, bank_start uint64, max_borrow uint64) *TritsCroupier {

	if lot%2 == 1 { // Man, we said even numbers only
		lot = lot - 1
	}
	bank := NewTritsAddress(BankAddr)
	croupier := new(TritsCroupier)                     // Wash your hands
	croupier.Banker = NewTritsBanker(lot + bank_start) // Hire a banker and put everything in the bank
	if TD {
		l(LOG_INFO, "CROUPIER invites BANKER and creates a bank account with ", lot+bank_start)
	}
	croupier.max_borrow = max_borrow
	croupier.Players = NewTritsSquad(croupier.Banker)               // Let in the players and tell them who is their banker
	size := croupier.Players.SizeOf()                               // How many players do we have
	amount_to_player := (lot - (lot % uint64(size))) / uint64(size) // Distribute evenly
	var i int = 0
	for i < size { // and divide the rest evenly among all the players
		croupier.Banker.MoveFunds(bank, croupier.Players.squad[i].Addr, amount_to_player)
		if TD {
			l(LOG_INFO, "BANKER sends ", amount_to_player, " to ", croupier.Players.squad[i].Name())
		}
		i++
	}
	croupier.Table = NewTritsTable() // Make a clean table with the required amount of trits
	var j int = 0
	for j < GAMES_ON_TABLE { // Create a bank account for each city on the table
		croupier.Banker.MoveFunds(bank, croupier.Table.GetCityAddress(j), 0)
		if TD {
			l(LOG_INFO, "BANKER creates an empty account for a new game ", GameName(croupier.Table.GetCityAddress(j)))
		}
		j++
	}
	if TD {
		l(LOG_DEBUG, "BANK: ", croupier.Banker.DumpBank())
	}
	return croupier
}

// Ask all players and do what they say
func (c *TritsCroupier) AskAround() bool {
	//var responses []*TritsPlayerResponse
	c.Players.Shuffle()               // Ask players in random order
	num_players := c.Players.SizeOf() // How many players do we have
	some_player_alive := false
	var i int = 0
	for i < num_players {
		if TD {
			c.Banker.gameshealthcheck(c.Table.Desk)
			l(LOG_DEBUG, "CROUPIER asks ", c.Players.squad[i].Name(), " to play...")
		}
		player_responses := c.Players.squad[i].Bet(c.Table.Desk)
		if player_responses != nil {
			some_player_alive = true
		}
		for _, r := range player_responses {
			// Place bet as instructed
			if TD {
				l(LOG_DEBUG, PlayerName(r.PlayerAddr), " says put ", r.Amount, " on ", GameName(r.Game.ThisGame))
			}
			if ok, err := c.Banker.MoveFunds(r.PlayerAddr, r.Game.ThisGame, r.Amount); ok {
				// Reflect it in the game, which may yield some todos for banker
				if TD {
					l(LOG_INFO, "BANKER sends ", r.Amount, " from ", PlayerName(r.PlayerAddr), " to ", GameName(r.Game.ThisGame))
					l(LOG_DEBUG, "BANK: ", c.Banker.DumpBank())
				}
				todos := r.Game.PlaceCoin(r.PlayerAddr, r.Amount)
				if len(todos) == 0 {
					if TD {
						l(LOG_DEBUG, "CROUPIER has got nothing to do")
						l(LOG_DEBUG, "GAME: ", LGame(r.Game))
						l(LOG_DEBUG, "BANK: ", c.Banker.DumpBank())
					}
				}
				if TD {
					l(LOG_DEBUG, "CROUPIER receives ", len(todos), " TODOs")
				}
				for _, do := range todos {
					switch do.Action {
					case ACTION_TRANSFER:
						if TD {
							l(LOG_DEBUG, "CROUPIER passes ACTION_TRANSFER to the BANKER")
						}
						if ok, err := c.Banker.MoveFunds(do.Funds_from, do.Funds_to, do.Amount); !ok {
							panic(err)
						}
					case ACTION_ASK_BONUS:
						if TD {
							l(LOG_DEBUG, "CROUPIER passes ACTION_ASK_BONUS to the BANKER")
						}
						bonus := c.Banker.PutBonus(r.Game)
						if bonus > 0 {
							r.Game.PlaceCoin(NewTritsAddress(BankAddr), bonus)
							if TD {
								l(LOG_DEBUG, "GAME: ", GameName(r.Game.ThisGame), " received the bonus")
								l(LOG_DEBUG, "GAME: ", LGame(r.Game))
							}
						} else {
							if TD {
								l(LOG_DEBUG, "BANK can't afford to put any bonus to ", GameName(r.Game.ThisGame))
								l(LOG_DEBUG, "GAME: ", LGame(r.Game))
							}
						}
					}
				}
			} else {
				panic(err)
			}
		}
		i++ // Next player
	}
	var j int = 0
	for j < num_players {
		if TD {
			c.Banker.gameshealthcheck(c.Table.Desk)
			l(LOG_DEBUG, "CROUPIER asks ", c.Players.squad[j].Name(), " if she want's to borrow some meoney ...")
		}
		amount := c.Players.squad[j].Borrow(c.max_borrow)

		if amount == 0 {
			j++
			continue
		} else {
			if TD {
				l(LOG_DEBUG, c.Players.squad[j].Name(), " borrows ", amount, " from ", LogName(NewTritsAddress(LenderAddr)))
			}
			c.Banker.MoveFunds(NewTritsAddress(LenderAddr), c.Players.squad[j].Addr, amount)
		}
		j++
	}
	return some_player_alive
}
