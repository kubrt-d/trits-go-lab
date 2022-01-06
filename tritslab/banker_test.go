package tritslab

import "testing"

func TestTritsBanker_MoveFunds(t *testing.T) {
	t.Run("Test banker constructor", func(t *testing.T) {
		b := NewTritsBanker(1000,50)
		got := b.Tell(NewTritsAddress(BankAddr))
		if got != 1000 {
			t.Errorf("TritsBanker.NewTritsBanker() got = %v, want %v", got, 1000)
		}
	})
	t.Run("Test banker move funds", func(t *testing.T) {
		b := NewTritsBanker(1000,50)
		bank := NewTritsAddress(BankAddr)
		neo := NewTritsAddress(NeoAddr)
		got, got1 := b.MoveFunds(bank, neo, 300)
		if !got {
			t.Errorf("TritsBanker.MoveFunds() got = %v, want %v", got1, true)
		}
	})
	t.Run("Test banker checks", func(t *testing.T) {
		b := NewTritsBanker(1000,50)
		bank := NewTritsAddress(BankAddr)
		neo := NewTritsAddress(NeoAddr)
		agent := NewTritsAddress(AgentAddr)
		b.MoveFunds(bank, neo, 300)
		b.MoveFunds(neo, agent, 200)
		trinity := NewTritsAddress(TrinityAddr)
		got, got1 := b.MoveFunds(neo, trinity, 150)
		if got {
			t.Errorf("TritsBanker.MoveFunds() got = %v, want %v", got1, "Insufficient funds")
		}
	})

	t.Run("Test sender doesn't exist", func(t *testing.T) {
		b := NewTritsBanker(0,50)
		neo := NewTritsAddress(NeoAddr)
		agent := NewTritsAddress(AgentAddr)
		got, got1 := b.MoveFunds(neo, agent, 1)
		if got {
			t.Errorf("TritsBanker.MoveFunds() got = %v, want %v", got1, "Sender does not exist")
		}
	})

	t.Run("Test banker tell", func(t *testing.T) {
		b := NewTritsBanker(uint64(500000000000000000),50)
		bank := NewTritsAddress(BankAddr)
		neo := NewTritsAddress(NeoAddr)
		b.MoveFunds(bank, neo, 400000000000000000)
		got := b.Tell(neo)
		if got != 400000000000000000 {
			t.Errorf("TritsBanker.MoveFunds() got = %v, want %v", got, 400000000000000000)
		}
	})
}
