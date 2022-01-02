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
	Amount     uint32        //	(optonal) how much to send
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
	Trit       *TritsTriangle // It's the triangle that holds the status
	ThisGame   *TritsAddress  // Address of this game
	Owner      *TritsAddress  // Whoever starts the game owns the game
	Nominal    uint32         // Game nominal (set by the first PlaceCoin call), 0 means the game has not yet started
	Middle     uint32         // Number of coins in the middle
	with_bonus bool           // Did bank send the bonus (always unless it runs out of money)
	updated    int64          // Nanotime last updated
	ll         int64          // Time to expire
	rand       Random3Dice    // Random number generator of choice
}

func (game *TritsGame) ResetGame() {
	game.Trit = NewTritsTriangle()
	game.Nominal = 0
	game.with_bonus = false
	game.updated = time.Now().UnixNano()
}

func NewTritsGame(addr *TritsAddress, lifelength int64, rand Random3Dice) *TritsGame {
	game := new(TritsGame)
	game.ResetGame()
	game.ThisGame = addr
	game.ll = lifelength
	game.with_bonus = false
	game.rand = rand
	return game
}

/* Place coin on the trit - by the time this function is called, the amount has already been added to the game address */
func (game *TritsGame) PlaceCoin(from *TritsAddress, amount uint32) []*TritsGameResponse {

	var responses []*TritsGameResponse
	response := NewGameResponse() // Default response is

	// EXPIRE GAME
	if time.Now().UnixNano() > game.updated+game.ll {
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
		if game.Nominal == 0 || amount%game.Nominal > 0 { // TODO: Log this as it should never happen
			response.Funds_from = game.ThisGame           // From this game
			response.Funds_to = NewTritsAddress(BankAddr) // Back to the Bank
			response.Amount = amount                      // the amount it sent
			responses = append(responses, response)       // Add the response
			return responses
		}
		coins := amount / game.Nominal    // How many coins did the banks sent ?
		game.Middle = game.Middle + coins // Update the amount of coins in the midle
		game.updated = time.Now().UnixNano()
		game.with_bonus = true
		return responses
	}

	// WRONG NOMINAL - the game has started and has got a nominal set, but the incoming amount is not equal (a mistake or client out of sync) */
	if game.Nominal > 0 && amount != game.Nominal {
		if game.Nominal > amount { // Not enough money, send it back
			response.Funds_from = game.ThisGame     // From this game
			response.Funds_to = from                // Back to whoever has sent it to us
			response.Amount = amount                // Refund the whole amount
			responses = append(responses, response) // Add the response (would be the deafult update or refund transfer)
			return responses
		} else { // Too much money, keep the nominal and send the rest back
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
		game.updated = time.Now().UnixNano() // It's an update
		game.Nominal = amount                // Set the nominal
		game.Middle = 1                      // First coin in the middle
		game.Owner = from                    // Set the game owner
		response.Action = ACTION_ASK_BONUS   // Ask for the bonus
		response.Funds_to = game.ThisGame    // To this game
		response.Funds_from = NewTritsAddress(BankAddr)
		responses = append(responses, response) // Add the response
		return responses
	}

	// THROW COIN - place coin randomly on the trit
	if game.Nominal > 0 && game.Middle > 0 {
		destiny := game.rand.Throw3Dice()          // Choose random arm
		win := game.Trit.HitVertice(destiny, from) // Place the coin on it
		if win {
			evil := game.rand.Throw3Dice() // Randomly choose the evil arm
			if destiny == evil {           // Ouch, you hit the bank's arm
				response.Funds_from = game.ThisGame           // From this game
				response.Funds_to = NewTritsAddress(BankAddr) // To the bank
				response.Amount = game.GetTotal()             // Bank takes it all
				responses = append(responses, response)       // Add the response
				return responses
			} else { // Hooray, you deserve it, what a game !
				if game.Owner.Equals(from) { // Game owner and the winner are the same, send it in one go
					response.Funds_from = game.ThisGame     // From this game
					response.Funds_to = from                // To the winner
					response.Amount = game.GetTotal()       // Refund everything
					responses = append(responses, response) // Add the response
					game.ResetGame()
					return responses
				} else {
					if game.with_bonus {
						response0 := NewGameResponse()           // Reward the owner
						response0.Funds_from = game.ThisGame     // From this game
						response0.Funds_to = game.Owner          // To the winner
						response0.Amount = 3 * game.Nominal      // All except the owner's reward
						game.Nominal -= response0.Amount         // Lower the total
						responses = append(responses, response0) // Add the response
					}
					response1 := NewGameResponse()                      // Send money to the winner
					response1.Funds_from = game.ThisGame                // From this game
					response1.Funds_to = from                           // To the winner
					response1.Amount = game.GetTotal() - 3*game.Nominal // All except the owner's reward
					responses = append(responses, response1)            // Add the response

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

/* Get the total on the table */
func (game *TritsGame) GetTotal() uint32 {
	return game.Nominal * (uint32(len(game.Trit.V1)) + uint32(len(game.Trit.V2)) + uint32(len(game.Trit.V3)) + game.Middle)
}
