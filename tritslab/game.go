package tritslab

import (
	"time"
)

const ACTION_TRANSFER = 1
const ACTION_ASK_BONUS = 2

type TritsGameResponse struct {
	Action     int8          // what to do, one of ACTION_TRANSFER ACTION_ASK_BONUS
	Funds_from *TritsAddress //	(optional) if funds transfer is involved, this tells from which address the funds should be taken
	Funds_to   *TritsAddress //	(optional) if funds transfer is involved, this tells to which address the funds should be sent
	Amount     uint64        //	(optonal) how much to send
}

func NewGameResponse() *TritsGameResponse { // Constructor
	gr := new(TritsGameResponse)
	gr.Action = ACTION_TRANSFER
	gr.Funds_from = nil
	gr.Funds_to = nil
	gr.Amount = 0
	return gr
}

type TritsGame struct {
	Trit     *TritsTriangle // It's the triangle that holds the status
	ThisGame *TritsAddress  // Address of this game
	Owner    *TritsAddress  // Whoever starts the game owns the game
	Nominal  uint64         // Game nominal (set by the first PlaceCoin call), 0 means the game has not yet started
	Middle   uint32         // Number of coins in the middle
	updated  int64          // Nanotime last updated
	ll       int64          // Time to expire
	rand     Random3Dice    // Random number generator of choice
}

func (game *TritsGame) ResetGame() {
	game.Trit = NewTritsTriangle()
	game.Nominal = 0
	game.Middle = 0
	game.Owner = nil
	game.updated = time.Now().UnixNano()
}

func NewTritsGame(addr *TritsAddress, lifelength int64, rand Random3Dice) *TritsGame {
	game := new(TritsGame)
	game.ResetGame()
	game.ThisGame = addr
	game.ll = lifelength
	game.rand = rand
	return game
}

/* Place coin on the trit - by the time this function is called, the amount has already been added to the game address */
func (game *TritsGame) PlaceCoin(from *TritsAddress, amount uint64) []*TritsGameResponse {

	var r []*TritsGameResponse = nil

	// NON ZERO POSITIVE AMOUNTS ONLY
	if amount <= 0 {
		return nil // Nothing happens
	}

	// EXPIRE GAME
	if time.Now().UnixNano() > game.updated+game.ll {
		if TD {
			l(LOG_NOTICE, "GAME: ", LogName(game.ThisGame), " has expired.")
		}
		total := game.GetTotal()
		if total > 0 {
			r = game.addResponse(r, ACTION_TRANSFER,
				game.ThisGame,             // From this game
				NewTritsAddress(BankAddr), // The bank takes it all
				game.GetTotal())           // Refund everything on the table
			// Add the response and!!! CONTINUE !!!
		}
		game.ResetGame()
	}

	// BANK BONUS - if this is Bank sending the bonus, just place it in the middle and return an empty response
	if from.SameAs(BankAddr) {
		if TD {
			l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " received bonus from the bank.")
		}
		if game.Nominal == 0 || amount%game.Nominal > 0 { // This should never happen
			if TD {
				l(LOG_WARN, "GAME: ", LogName(game.ThisGame), " received incorrect bonus, asking croupier to send it back.")
			}
			r = game.addResponse(r, ACTION_TRANSFER,
				game.ThisGame,             // From this game
				NewTritsAddress(BankAddr), // Back to the Bank
				amount)                    // The amount bank has sent
			return r
		}
		coins := uint32(amount / game.Nominal) // How many coins did the banks sent ?
		game.Middle = game.Middle + coins      // Update the amount of coins in the midle
		if TD {
			l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " adding ", coins, " coins on the middle")
		}
		game.updated = time.Now().UnixNano()
		return r
	}

	// WRONG NOMINAL - the game has started and has got a nominal set, but the incoming amount is not equal (a mistake or client out of sync) */
	if game.Nominal > 0 && amount != game.Nominal {
		if game.Nominal > amount { // Not enough money, send it back
			if TD {
				l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " recieved an incorrect (smaller) amount ", amount, " from ", LogName(from), ", sending it back")
			}
			r = game.addResponse(r, ACTION_TRANSFER,
				game.ThisGame, // From this game
				from,          // Back to whoever has sent it to us
				amount)        // Refund the whole amount
			return r
		} else { // Too much money, keep the nominal and send the rest back
			if TD {
				l(LOG_NOTICE, "GAME: ", LogName(game.ThisGame), " recieved an incorrect (higher) amount ", amount, " from ", LogName(from), ", using the nominal and sending back the rest")
			}
			r = game.addResponse(r, ACTION_TRANSFER,
				game.ThisGame,       // From this game
				from,                // Back to whoever has sent it to us
				amount-game.Nominal) // Refund whatever is over the nominal
			amount = game.Nominal // Adjust the amount and !!! CONTINUE !!!
			// !!! CONTINUE !!!
		}
	}

	// START GAME - The first coin that has arrived to this game
	if game.Nominal == 0 {
		if TD {
			l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " started with nominal ", amount, " by ", LogName(from))
		}
		game.updated = time.Now().UnixNano() // It's an update
		game.Nominal = amount                // Set the nominal
		game.Middle = 1                      // First coin in the middle
		game.Owner = from                    // Set the game owner
		if TD {
			l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " is asking bank for the bonus")
		}
		r = game.addResponse(r, ACTION_ASK_BONUS, // Ask for the bonus
			NewTritsAddress(BankAddr), // From the bank
			game.ThisGame,             // To this game
			0)                         // We don't know how much
		return r
	}

	// THROW COIN - place coin randomly on the trit
	if game.Nominal > 0 && game.Middle > 0 {
		destiny := game.rand.Throw3Dice() // Choose random arm
		if TD {
			l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " coin accepted with the destiny ", destiny)
		}
		win := game.Trit.HitVertice(destiny, from) // Place the coin on it
		if win {
			if TD {
				l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " finished by a throw from ", LogName(from), " on arm ", destiny)
			}
			if TD {
				l(LOG_DEBUG, "GAME: ", LGame(game))
			}
			evil := game.rand.Throw3Dice() // Randomly choose the evil arm
			if TD {
				l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " the evil arm number is ", evil)
			}
			if destiny == evil { // Ouch, you hit the bank's arm
				if TD {
					l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " bad luck, hit the banks arm. ")
				}
				var rew uint32 = 0
				if game.GetTotal() >= 2*game.Nominal {
					rew = 2
					if TD {
						l(LOG_DEBUG, "GAME: can afford to double the owner's money, ", game.Nominal*uint64(rew), " goes to ", LogName(game.Owner))
					}

				} else {
					rew = 1
					if TD {
						l(LOG_DEBUG, "GAME: can only affor a money return, ", game.Nominal*uint64(rew), " goes to ", LogName(game.Owner))
					}
				}
				amount_to_owner := game.Nominal * uint64(rew)
				amount_to_bank := game.GetTotal() - amount_to_owner
				r = game.addResponse(r, ACTION_TRANSFER,
					game.ThisGame,   // From this game
					game.Owner,      // To the game owner
					amount_to_owner) // 1-2  coins from the middle
				game.Middle -= rew // Should be 0 after this

				r = game.addResponse(r, ACTION_TRANSFER,
					game.ThisGame,             // From this game
					NewTritsAddress(BankAddr), // To the bank
					amount_to_bank)            // Bank takes it all except 1-2 coins
				game.ResetGame() // Reset game
				return r

			} else { // Hooray, you deserve it, what a game !
				game_total := game.GetTotal()
				game_nominal := game.Nominal
				game_tripple := uint64(3) * game.Nominal
				if game_total%game_nominal != 0 {
					panic("Something went terribly wrong with the game total versus nominal")
				}
				if game.Owner.Equals(from) { // Game owner and the winner are the same, send it all in one go
					if TD {
						l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " winner same as owner, sending ", game_total, " to ", LogName(from))
					}
					r = game.addResponse(r, ACTION_TRANSFER,
						game.ThisGame, // From this game
						from,          // To the winner
						game_total)    // Everything
					game.ResetGame()
					return r
				} else {
					if game.Middle > 1 { // Game with bonus
						if TD {
							l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " rewarding the owner, sending ", game_tripple, " to ", LogName(game.Owner))
						}
						r = game.addResponse(r, ACTION_TRANSFER, // Reward the game owner
							game.ThisGame, // From this game
							game.Owner,    // To the game owner
							game_tripple)  // All except the owner's reward
						game_total -= game_tripple // Lower the total
						game.Middle -= 3           // Lower the middle counter
					} else { // Game without bonus
						if TD {
							l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " the 1 middle coin back the owner, sending ", game.Nominal, " to ", LogName(game.Owner))
						}
						r = game.addResponse(r, ACTION_TRANSFER, // Reward the game owner
							game.ThisGame, // From this game
							game.Owner,    // To the game owner
							game.Nominal)  // Just one coin
						game_total -= game.Nominal // Lower the total
						game.Middle -= 1           // Lower the middle counter
					}
					if TD {
						l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " rewarding the winner, sending ", game_total-game_tripple, " to ", LogName(from))
					}
					r = game.addResponse(r, ACTION_TRANSFER, // Send (the rest) of the money to the winner
						game.ThisGame, // From this game
						from,          // To the winner
						game_total)    // All (except the owner's reward)
					game.ResetGame()
					return r
				}
			}
		} else {
			return r
		}
	}
	return r // Nothing to do - TODO: should this be an error as it should never happen ?
}

// Add response to teh responses slice
func (game *TritsGame) addResponse(r []*TritsGameResponse, action int8, from *TritsAddress, to *TritsAddress, amount uint64) []*TritsGameResponse {
	response := NewGameResponse()
	response.Action = action
	response.Funds_from = from
	response.Funds_to = to
	response.Amount = amount
	r = append(r, response)
	return r
}

// Get the total on the table
func (game *TritsGame) GetTotal() uint64 {
	return game.Nominal * uint64((len(game.Trit.V1) + len(game.Trit.V2) + len(game.Trit.V3) + int(game.Middle)))
}

// Get the Trit inbalance
func (game *TritsGame) GetInbalance() int8 {
	return game.Trit.Inbalance()
}
