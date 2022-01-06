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

	var responses []*TritsGameResponse
	response := NewGameResponse() // Default response is

	// EXPIRE GAME
	if time.Now().UnixNano() > game.updated+game.ll {
		if TD {
			l(LOG_NOTICE, "GAME: ", LogName(game.ThisGame), " has expired.")
		}
		total := game.GetTotal()
		if total > 0 {
			response.Funds_from = game.ThisGame           // From this game
			response.Funds_to = NewTritsAddress(BankAddr) // The bank takes it all
			response.Amount = game.GetTotal()             // Refund everything on the table
			responses = append(responses, response)       // Add the response and!!! CONTINUE !!!
			response = NewGameResponse()                  // Clear the response variable
		}
		game.ResetGame()
	}

	// BANK BONUS - if this is Bank sending the bonus, just place it in the middle and return an empty response
	if from.SameAs(BankAddr) {
		if TD {
			l(LOG_NOTICE, "GAME: ", LogName(game.ThisGame), " received bonus from the bank.")
		}
		if game.Nominal == 0 || amount%game.Nominal > 0 { // This should never happen
			if TD {
				l(LOG_WARN, "GAME: ", LogName(game.ThisGame), " received incorrect bonus, asking croupier to send it back.")
			}
			response.Funds_from = game.ThisGame           // From this game
			response.Funds_to = NewTritsAddress(BankAddr) // Back to the Bank
			response.Amount = amount                      // the amount it sent
			responses = append(responses, response)       // Add the response
			return responses
		}
		coins := uint32(amount / game.Nominal) // How many coins did the banks sent ?
		game.Middle = game.Middle + coins      // Update the amount of coins in the midle
		if TD {
			l(LOG_NOTICE, "GAME: ", LogName(game.ThisGame), " adding ", coins, " coins on the middle")
		}
		game.updated = time.Now().UnixNano()
		return responses
	}

	// WRONG NOMINAL - the game has started and has got a nominal set, but the incoming amount is not equal (a mistake or client out of sync) */
	if game.Nominal > 0 && amount != game.Nominal {
		if game.Nominal > amount { // Not enough money, send it back
			if TD {
				l(LOG_NOTICE, "GAME: ", LogName(game.ThisGame), " recieved an incorrect (smaller) amount ", amount, " from ", LogName(from), ", sending it back")
			}
			response.Funds_from = game.ThisGame     // From this game
			response.Funds_to = from                // Back to whoever has sent it to us
			response.Amount = amount                // Refund the whole amount
			responses = append(responses, response) // Add the response (would be the deafult update or refund transfer)
			return responses
		} else { // Too much money, keep the nominal and send the rest back
			if TD {
				l(LOG_NOTICE, "GAME: ", LogName(game.ThisGame), " recieved an incorrect (higher) amount ", amount, " from ", LogName(from), ", using the nominal and sending back the rest")
			}
			response.Funds_from = game.ThisGame     // From this game
			response.Funds_to = from                // Back to whoever has sent it to us
			response.Amount = amount - game.Nominal // Refund whatever is over the nominal
			responses = append(responses, response) // Add the response (would be the deafult update or refund transfer)
			response = NewGameResponse()            // Clear the response variable
			amount = game.Nominal                   // Adjust the amount and !!! CONTINUE !!!
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
		response.Action = ACTION_ASK_BONUS // Ask for the bonus
		response.Funds_to = game.ThisGame  // To this game
		response.Funds_from = NewTritsAddress(BankAddr)
		responses = append(responses, response) // Add the response
		return responses
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
				l(LOG_DEBUG, "GAME: ", lgame(game))
			}
			evil := game.rand.Throw3Dice() // Randomly choose the evil arm
			if TD {
				l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " the evil arm's number is ", evil)
			}
			if destiny == evil { // Ouch, you hit the bank's arm
				if TD {
					l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " bad luck, hit the banks arm. ", game.GetTotal(), " goes to the bank.")
				}
				response.Funds_from = game.ThisGame           // From this game
				response.Funds_to = NewTritsAddress(BankAddr) // To the bank
				response.Amount = game.GetTotal()             // Bank takes it all
				responses = append(responses, response)       // Add the response
				game.ResetGame()                              // Reset game
				return responses
			} else { // Hooray, you deserve it, what a game !
				game_total := game.GetTotal()
				game_nominal := game.Nominal
				game_tripple := uint64(3) * game.Nominal
				if game_total%game_nominal != 0 {
					panic("Something went terribly wrong with the game total versus nominal")
				}
				if game.Owner.Equals(from) { // Game owner and the winner are the same, send it in one go
					if TD {
						l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " winner same as owner, sending ", game_total, " to ", LogName(from))
					}
					response.Funds_from = game.ThisGame     // From this game
					response.Funds_to = from                // To the winner
					response.Amount = game_total            // Everything
					responses = append(responses, response) // Add the response
					game.ResetGame()
					return responses
				} else {
					if game.Middle > 1 { // Game with bonus
						if TD {
							l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " rewarding the owner, sending ", game_tripple, " to ", LogName(game.Owner))
						}
						response0 := NewGameResponse()           // Reward the game owner
						response0.Funds_from = game.ThisGame     // From this game
						response0.Funds_to = game.Owner          // To the game owner
						response0.Amount = game_tripple          // All except the owner's reward
						game.Middle -= 3                         // Lower the total
						responses = append(responses, response0) // Add the response
					}
					if TD {
						l(LOG_DEBUG, "GAME: ", LogName(game.ThisGame), " rewarding the winner, sending ", game_total-game_tripple, " to ", LogName(from))
					}
					response1 := NewGameResponse()               // Send money to the winner
					response1.Funds_from = game.ThisGame         // From this game
					response1.Funds_to = from                    // To the winner
					response1.Amount = game_total - game_tripple // All except the owner's reward
					responses = append(responses, response1)     // Add the response
					game.ResetGame()
					return responses
				}
			}
		} else {
			return responses
		}
	}

	return responses // Nothing to do - TODO: should this be an error as it should never happen ?
}

// Get the total on the table
func (game *TritsGame) GetTotal() uint64 {
	return game.Nominal * uint64((len(game.Trit.V1) + len(game.Trit.V2) + len(game.Trit.V3) + int(game.Middle)))
}

func (game *TritsGame) GetInbalance() int8 {
	return game.Trit.Inbalance()
}
