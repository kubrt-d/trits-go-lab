package tritslab

//"math/rand"
//"time"

type TritsCroupier struct {
	table   *TritsTable  // Croupier's table
	players *TritsSquad  // Players
	banker  *TritsBanker // Banker
}

// Croupier starts his day, with a certain lot for players and something to give to the bank
func NewTritsCroupier(lot uint64, bank_start uint64) *TritsCroupier {
	if lot%2 == 1 { // Man, we said even numbers only
		lot = lot - 1
	}
	bank := NewTritsAddress(BankAddr)
	croupier := new(TritsCroupier)                                  // Wash your hands
	croupier.banker = NewTritsBanker(lot + bank_start)              // Hire a banker and put everything in the bank
	croupier.players = NewTritsSquad(croupier.banker)               // Let in the players and tell them who is their banker
	size := croupier.players.SizeOf()                               // How many players do we have
	amount_to_player := (lot - (lot % uint64(size))) / uint64(size) // Keep about a half for the bank
	var i int = 0
	for i < size { // and divide the rest evenly among all the players
		croupier.banker.MoveFunds(bank, croupier.players.squad[i].Addr, amount_to_player)
		i++
	}
	croupier.table = NewTritsTable() // Make a clean table with the required amount of trits
	var j int = 0
	for j < GAMES_ON_TABLE { // Create a bank account for each city on the table
		croupier.banker.MoveFunds(bank, croupier.table.GetCityAddress(j), 0)
		j++
	}
	return croupier
}

// Ask all players and do what they say
func (c *TritsCroupier) AskAround() {
	//var responses []*TritsPlayerResponse
	c.players.Shuffle()               // Ask players in random order
	num_players := c.players.SizeOf() // How many players do we have
	var i int = 0
	for i < num_players {
		player_responses := c.players.squad[i].Bet(c.table.desk)
		for _, r := range player_responses {
			// Place bet as instructed
			if ok, err := c.banker.MoveFunds(r.PlayerAddr, r.Game.ThisGame, r.Amount); ok {
				// Reflect it in the game, which may yield some todos for banker
				todos := r.Game.PlaceCoin(r.PlayerAddr, r.Amount)
				for _, do := range todos {
					switch do.Action {
					case ACTION_TRANSFER:
						if ok, err := c.banker.MoveFunds(do.Funds_from, do.Funds_to, do.Amount); !ok {
							panic(err)
						}
					case ACTION_ASK_BONUS:
						bonus := c.banker.PutBonus(r.Game)
						if bonus > 0 {
							r.Game.PlaceCoin(NewTritsAddress(BankAddr), bonus)
						}
					}
				}
			} else {
				panic(err)
			}
		}
		//responses = append(responses, player_responses...)
		i++
	}
	// Shuffle the responses
	//rand.Seed(time.Now().UnixNano())
	//rand.Shuffle(len(responses), func(i, j int) { responses[i], responses[j] = responses[j], responses[i] })
	// Process all responses one by one

}
