package tritslab

import (
	"testing"
	"time"
)

func test(t *testing.T, r TritsGameResponse, a TritsGameResponse) bool {
	var ok bool = true
	if r.Action != a.Action {
		ok = false
		t.Errorf("TritsGame.PlaceCoin().Action = %v, want %v", r.Action, a.Action)
	}
	if r.Amount != a.Amount {
		ok = false
		t.Errorf("TritsGame.PlaceCoin().Amount = %v, want %v", r.Amount, a.Amount)
	}
	if !r.Funds_from.Equals(a.Funds_from) {
		ok = false
		t.Errorf("TritsGame.PlaceCoin().Funds_from = %v, want %v", r.Funds_from.Human(), a.Funds_from.Human())
	}
	if !r.Funds_to.Equals(a.Funds_to) {
		ok = false
		t.Errorf("TritsGame.PlaceCoin().Funds_to = %v, want %v", r.Funds_to.Human(), a.Funds_to.Human())
	}
	return ok
}

/* Assert template
assert := NewGameResponse()response0 :=
		response1 := res[1]
assert.Amount = 0
assert.Funds_from = game_addr
assert.Funds_to = NewTritsAddress(BankAddr)
assert.Payload = ""
test(t, response,assert)
*/

func TestTritsGame_PlaceCoin(t *testing.T) {

	dice := NewTritsSequence([]byte{1, 2, 3, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 1, 2, 3, 1, 1})
	game_addr := NewTritsAddress("e28533750bee16842a5cd4f533d235770e407367")
	neo_addr := NewTritsAddress(NeoAddr)
	bank_addr := NewTritsAddress(BankAddr)
	trinity_addr := NewTritsAddress(TrinityAddr)
	morpheus_addr := NewTritsAddress(MorpheusAddr)
	agent_addr := NewTritsAddress(AgentAddr)
	keymaker_addr := NewTritsAddress(KeymakerAddr)

	game := NewTritsGame(game_addr, 60000000000, dice)

	// BANK GOES MAD
	t.Run("Bank goes mad", func(t *testing.T) {
		res := game.PlaceCoin(bank_addr, 1000)

		response := res[0]
		assert := NewGameResponse()
		assert.Action = ACTION_TRANSFER
		assert.Amount = 1000
		assert.Funds_from = game_addr
		assert.Funds_to = NewTritsAddress(BankAddr)
		test(t, response, assert)
	})
	// BANK SENDS 0
	t.Run("Bank sends 0", func(t *testing.T) {
		res := game.PlaceCoin(bank_addr, 0)

		if res != nil {
			t.Errorf("TritsGame.PlaceCoin().PlaceCoin = %v, want %v", "some response", nil)
		}
	})
	// NEO STARTS THE GAME
	t.Run("Neo sends 300 and starts the game", func(t *testing.T) {
		res := game.PlaceCoin(neo_addr, 300)

		assert := NewGameResponse()
		assert.Action = ACTION_ASK_BONUS
		assert.Funds_from = bank_addr
		assert.Funds_to = game_addr
		test(t, res[0], assert)
		total := game.GetTotal()
		if total != 300 {
			t.Errorf("TritsGame.PlaceCoin().GetTotal() = %v, want %v", total, 300)
		}
	})
	// BANK SEND BONUS
	t.Run("Bank sends bonus", func(t *testing.T) {
		game.PlaceCoin(bank_addr, 4*300)

		total := game.GetTotal()
		if total != 5*300 {
			t.Errorf("TritsGame.PlaceCoin().GetTotal() = %v, want %v", total, 5*300)
		}
		middle := game.Middle
		if middle != 5 {
			t.Errorf("TritsGame.PlaceCoin().Middle = %v, want %v", middle, 5)
		}
	})
	// TRINITY SENDS TOO MUCH - 2 coins are accepted and 200 is returned
	// 1,2  trinity 800 => 5(1,1,0)
	t.Run("Trinity sends too much", func(t *testing.T) {
		res := game.PlaceCoin(trinity_addr, 800)

		response := res[0]

		assert := NewGameResponse()
		assert.Action = ACTION_TRANSFER
		assert.Funds_from = game_addr
		assert.Funds_to = trinity_addr
		assert.Amount = 200
		test(t, response, assert)
		total := game.GetTotal()
		v10_addr_h := game.Trit.V1[0].Human()
		v20_addr_h := game.Trit.V2[0].Human()
		if total != 7*300 {
			t.Errorf("TritsGame.PlaceCoin().GetTotal() = %v, want %v", total, 7*300)
		}
		middle := game.Middle
		if middle != 5 {
			t.Errorf("TritsGame.PlaceCoin().Middle = %v, want %v", middle, 5)
		}
		if v10_addr_h != trinity_addr.Human() {
			t.Errorf("TritsGame.PlaceCoin().V1[0] = %v, want %v", v10_addr_h, trinity_addr.Human())
		}
		if v20_addr_h != trinity_addr.Human() {
			t.Errorf("TritsGame.PlaceCoin().V1[0] = %v, want %v", v20_addr_h, trinity_addr.Human())
		}

	})
	// MORPHEUS SENDS 2 COINS - V2, V3
	// 3, morpheus 300 => 5(1,1,1)
	// 2, morpheus 300 => 5(1,2,1)
	t.Run("Morpheus sends two coins", func(t *testing.T) {
		game.PlaceCoin(morpheus_addr, 300)
		game.PlaceCoin(morpheus_addr, 300)
		total := game.GetTotal()
		v21_addr_h := game.Trit.V2[1].Human()
		v30_addr_h := game.Trit.V3[0].Human()
		if total != 9*300 {
			t.Errorf("TritsGame.PlaceCoin().GetTotal() = %v, want %v", total, 9*300)
		}

		if v21_addr_h != morpheus_addr.Human() {
			t.Errorf("TritsGame.PlaceCoin().V1[0] = %v, want %v", v21_addr_h, morpheus_addr.Human())
		}
		if v30_addr_h != morpheus_addr.Human() {
			t.Errorf("TritsGame.PlaceCoin().V1[0] = %v, want %v", v30_addr_h, morpheus_addr.Human())
		}
	})
	// AGENT SENDS 2 COINS - V2,V2
	// 2, agent 300 => 5(1,3,1)
	// 3, agent 300 => 5(1,3,2)
	t.Run("Agent sends two coins", func(t *testing.T) {
		game.PlaceCoin(agent_addr, 300)
		game.PlaceCoin(agent_addr, 300)
		total := game.GetTotal()
		v22_addr_h := game.Trit.V2[2].Human()
		v31_addr_h := game.Trit.V3[1].Human()
		if total != 11*300 {
			t.Errorf("TritsGame.PlaceCoin().GetTotal() = %v, want %v", total, 11*300)
		}

		if v22_addr_h != agent_addr.Human() {
			t.Errorf("TritsGame.PlaceCoin().V1[0] = %v, want %v", v22_addr_h, agent_addr.Human())
		}
		if v31_addr_h != agent_addr.Human() {
			t.Errorf("TritsGame.PlaceCoin().V1[0] = %v, want %v", v31_addr_h, agent_addr.Human())
		}
	})
	// KEYMAKER SENDS TOO LITTLE, THEN TOO MUCH
	// 1,2 keymaker 666 =>5(2,4,2)
	t.Run("Keymaker tries 3 times", func(t *testing.T) {
		res := game.PlaceCoin(keymaker_addr, 133)
		response := res[0]
		assert := NewGameResponse()
		assert.Action = ACTION_TRANSFER
		assert.Funds_from = game_addr
		assert.Funds_to = keymaker_addr
		assert.Amount = 133
		test(t, response, assert)
		total := game.GetTotal()
		if total != 11*300 {
			t.Errorf("TritsGame.PlaceCoin().GetTotal() = %v, want %v", total, 11*300)
		}

		res = game.PlaceCoin(keymaker_addr, 666)
		response = res[0]
		assert = NewGameResponse()
		assert.Action = ACTION_TRANSFER
		assert.Funds_from = game_addr
		assert.Funds_to = keymaker_addr
		assert.Amount = 66
		test(t, response, assert)
		total = game.GetTotal()
		if total != 13*300 {
			t.Errorf("TritsGame.PlaceCoin().GetTotal() = %v, want %v", total, 13*300)
		}

		v23_addr_h := game.Trit.V2[3].Human()
		if v23_addr_h != keymaker_addr.Human() {
			t.Errorf("TritsGame.PlaceCoin().V3[1] = %v, want %v", v23_addr_h, keymaker_addr.Human())
		}

	})
	// NEO WINS THE GAME (with a slightly wrong amount)
	t.Run("Neo sends 610, wins the game and starts a new one", func(t *testing.T) {

		responses := game.PlaceCoin(neo_addr, 610)
		res0 := responses[0]
		assert0 := NewGameResponse()
		assert0.Action = ACTION_TRANSFER
		assert0.Funds_from = game_addr
		assert0.Funds_to = neo_addr
		assert0.Amount = 14 * 300
		test(t, res0, assert0)

		res1 := responses[1]
		assert1 := NewGameResponse()
		assert1.Action = ACTION_ASK_BONUS
		assert1.Funds_from = bank_addr
		assert1.Funds_to = game_addr
		assert1.Amount = 0
		test(t, res1, assert1)
		total := game.GetTotal()
		if total != 310 {
			t.Errorf("TritsGame.PlaceCoin().GetTotal() = %v, want %v", total, 310)
		}

	})
	// Reward the owner
	//TODO: Re write this test to take into teh acccount that Neo may not win
	random_dice := NewTritsDice()
	/*
		random_game := NewTritsGame(game_addr, 60000000000, random_dice)
		t.Run("Trinity starts then gets a reward from a random game", func(t *testing.T) {
			random_game.PlaceCoin(trinity_addr, 1)
			random_game.PlaceCoin(bank_addr, 4)
			var last_responses []*TritsGameResponse
			for random_game.Nominal != 0 { // Neo play random until he wins eventually
				last_responses = random_game.PlaceCoin(neo_addr, 1)
			}
			res0 := last_responses[0]
			assert0 := NewGameResponse()
			assert0.Action = ACTION_TRANSFER
			assert0.Funds_from = game_addr
			assert0.Funds_to = trinity_addr
			assert0.Amount = 3
			test(t, res0, assert0)
		})
	*/
	// Reward the owner against the evil
	//TODO: Write this test
	// Test expiration (0.1s)
	t.Run("Test expiration", func(t *testing.T) {
		ephemeral_game := NewTritsGame(game_addr, 1000000, random_dice)
		time_start := time.Now().UnixNano()
		ephemeral_game.PlaceCoin(trinity_addr, 333)
		ephemeral_game.PlaceCoin(bank_addr, 5*333)
		ephemeral_game.PlaceCoin(keymaker_addr, 333)
		for time.Now().UnixNano() <= time_start+10000000 {
		}
		response := ephemeral_game.PlaceCoin(neo_addr, 10)
		res0 := response[0]
		assert0 := NewGameResponse()
		assert0.Action = ACTION_TRANSFER
		assert0.Funds_from = game_addr
		assert0.Funds_to = trinity_addr
		assert0.Amount = 333
		test(t, res0, assert0)
		res1 := response[1]
		assert1 := NewGameResponse()
		assert1.Action = ACTION_TRANSFER
		assert1.Funds_from = game_addr
		assert1.Funds_to = bank_addr
		assert1.Amount = 6 * 333
		test(t, res1, assert1)
		total := ephemeral_game.GetTotal()
		if total != 10 {
			t.Errorf("TritsGame.PlaceCoin().GetTotal() = %v, want %v", total, 10)
		}
	})

}

func BenchmarkTritsGame_AddResponse(b *testing.B) {
	game_addr := NewTritsAddress("e28533750bee16842a5cd4f533d235770e407367")
	random_dice := NewTritsDice()
	neo_addr := NewTritsAddress(NeoAddr)
	bank_addr := NewTritsAddress(BankAddr)
	game := NewTritsGame(game_addr, 1000000, random_dice)
	var r []TritsGameResponse = make([]TritsGameResponse, 0,5)

	for n := 0; n < b.N; n++ {
		game.addResponse(r,ACTION_TRANSFER,neo_addr,bank_addr,1000)
	}
}