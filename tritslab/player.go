package tritslab

type TritsPlayerResponse struct {
	PlayerAddr TritsAddress
	Game       *TritsGame
	Amount     uint64
}

func NewTritsPlayerResponse(game *TritsGame, amount uint64, my_addr TritsAddress) TritsPlayerResponse {
	r := TritsPlayerResponse{}
	r.PlayerAddr = my_addr
	r.Game = game
	r.Amount = amount
	return r
}

// Composition of the base TritsPlayer
func NewTritsPlayerFactory(addr TritsAddress, banker TritsBanker, strategy string) TritsPlayer {
	var out TritsPlayer
	switch strategy {
	case "zion":
		player := TritsPlayerZion{}
		player.Addr = addr
		player.Player_type = "zion"
		player.banker = banker
		out = &player
	case "sharp":
		player := TritsPlayerZionSharp{}
		player.Addr = addr
		player.Player_type = "zion"
		player.banker = banker
		out = &player
	case "agent":
		player := TritsPlayerAgent{}
		player.Addr = addr
		player.Player_type = "agent"
		player.banker = banker
		out = &player
	default:
		player := TritsPlayerDumb{}
		player.Addr = addr
		player.Player_type = "dumb"
		player.banker = banker
		out = &player
	}
	return out
}

type TritsPlayer interface {
	SetStartedWith(started_with uint64)
	ChooseTable() int8
	ChooseAmount() uint64
	Balance() uint64
	Borrow(max_borrow uint64) uint64
	TakeProfit() uint64
	Bet(desk []*TritsGame) []TritsPlayerResponse
	Name() string
	GetAddr() TritsAddress
	GetPlayerType() string
	Recharge() uint64
}

// Helper function to get player name by address statically
func PlayerName(addr TritsAddress) string {
	p := TritsPlayerDumb{}
	p.Addr = addr
	return p.Name()
}
