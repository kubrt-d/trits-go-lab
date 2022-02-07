package tritslab

import "testing"

func TestTritsDice_Throw3Dice(t *testing.T) {
	dice := NewTritsDice()

	t.Run("Throw 3dice", func(t *testing.T) {
		got := dice.Throw3Dice()
		if got > 3 || got < 1 {
			t.Errorf("TritsSequence.Throw3Dice() = %v, want %v", got, " one of [1,2,3]")
		}
	})
}

func BenchmarkTritsDice_Throw3Dice(b *testing.B) {
	dice := NewTritsDice()
	for n := 0; n < b.N; n++ {
		dice.Throw3Dice()
	}
}

