package tritslab

import (
	"log"
	"time"
)

const RawBankAddress string = "ffffffffffffffffffffffffffffffffffffffff"

type TritsGame struct {
	Trit     *TritsTriangle // It's the triangle that holds the status
	ThisGame *TritsAddress  // Address of this game
	Seeder   *TritsAddress  // Address of the Seeder
	Winner   *TritsAddress  // Address of the winner
	Nominal  uint32         // Game nominal (set by the first PlaceCoin call)
	created  int64          // Nanotime created
	updated  int64          // Nanotime last updated
	is_new   bool           // Is it a freshly created game ?
	is_over  bool           // Is the game over
	ll       uint64         // Time to expire
	i_count  uint32         // Number of coins in the middle
	rand     Random3Dice    // Random number generator of choice
}

func NewTritsGame(addr *TritsAddress, lifelength uint64, rand Random3Dice) *TritsGame {
	game := new(TritsGame)
	game.is_new = true
	game.is_over = false
	game.Nominal = 0
	game.ll = lifelength
	game.rand = rand
	game.Trit = NewTritsTriangle(0, 0, 0)
	game.created = time.Now().UnixNano()
	game.updated = time.Now().UnixNano()
	return game
}

/* Place coin on the trit */
func (game *TritsGame) PlaceCoin(from *TritsAddress, amount uint32) string {

	x := string(game.rand.Throw3Dice())

	log.Fatal(x)
	// If this is Bank sending the bonus, just place it in the middle
	if from.SameAs(RawBankAddress) {
		coins := amount / game.Nominal
		game.i_count = game.i_count + coins
		return ""
	}

	// Adjust amount to the nominal/multiple of the nominal (client out of sync with the game state)
	if !game.is_new && !game.is_over && game.Nominal > 0 && amount%game.Nominal > 0 {
		reminder := amount % game.Nominal
		amount := amount - game.Nominal
		coins := amount / game.Nominal
		for coins > 0 {
			coins--
			game.PlaceCoin(from, game.Nominal) // Recursively place all the coins

		}
		if reminder > 0 {
			return "BookKeeperDo(ThisGame,Sender,reminder)" // Return the reminder
		} else {
			return "" //TODO: Those recursive calls may/should/will return tangible actions
		}
	}

	// Previous game has ended, start a new one and continue with this coin
	if game.is_over && !game.is_new {
		game.is_new = true
		game.is_over = false
		game.Nominal = 0
		game.Seeder = nil
		game.Winner = nil
		game.created = time.Now().UnixNano()
		game.updated = time.Now().UnixNano()
	}

	// The first coin that has arrived to this game
	if game.is_new && !game.is_over {
		game.updated = time.Now().UnixNano()          // It's an update
		game.Seeder = from                            // Assign a seeder (address of the player who started this game)
		game.Winner = nil                             // Just to be sure
		game.Nominal = amount                         // Set the nominal
		game.i_count = 1                              // First coin in the middle
		return "BankGiveBonus(ThisGame,game.Nominal)" // Ask bank to give the bonus
	}

	return "" // Nothing to do
}
