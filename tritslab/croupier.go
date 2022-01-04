package tritslab

type TritsCroupier struct {
	table   *TritsTable  // Croupier's table
	players *TritsSquad  // Players
	banker  *TritsBanker // Banker
}

// Croupier starts his day, with a certain lot for players and something to give to the bank 
func NewTritsCroupier(lot uint64, bankstart uint64) *TritsCroupier {
	if lot%2 == 1 { // Man, we said even numbers only
		lot = lot - 1
	}
	bank := NewTritsAddress(BankAddr)
	croupier := new(TritsCroupier)                                          // Wash your hands
	croupier.banker = NewTritsBanker(lot + bankstart)                       // Hire a banker and put everything in the bank
	croupier.players = NewTritsSquad(croupier.banker)                       // Let in the players and tell them who is their banker
	size := croupier.players.SizeOf()                                       // How many players do we have
	amount_to_player := (lot - (lot % uint64(size))) / uint64(size) // Keep about a half for the bank
	var i int = 0
	for i < size { // and divide the rest evenly among all the players
		croupier.banker.MoveFunds(bank, croupier.players.squad[i].Addr, amount_to_player)
		i++
	}
	croupier.table = NewTritsTable() // Make a clean table with the required amount of trits
	return croupier
}

// Ask all players
func (p *TritsCroupier) AskAround() int8 {
	return 1
}
