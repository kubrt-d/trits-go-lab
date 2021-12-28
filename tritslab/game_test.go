package tritslab

import "testing"

func TestTritsGame_PlaceCoin(t *testing.T) {
	
	dice := NewTritsSequence([]int8{1, 2, 3, 2, 1, 3, 1, 2, 1, 3, 2})
	
	var game_raw_addr rawaddr = []byte("e28533750bee16842a5cd4f533d235770e407367")
	game_addr := NewTritsAddress(game_raw_addr)

	game := NewTritsGame(game_addr, 10000, dice)

	var neo_raw_addr rawaddr = []byte("b35413750bee16842a5cd4f533d23577857412af")
	neo_addr := NewTritsAddress(neo_raw_addr)


	t.Run("Test sequence", func(t *testing.T) {
		got := game.PlaceCoin(neo_addr, 1000)
		if got != "yes" {
			t.Errorf("TritsSequence.Throw3Dice() = %v, want %v", got, "yes")
		}
	})

}
