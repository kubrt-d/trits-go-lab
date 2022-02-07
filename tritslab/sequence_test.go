package tritslab

import "testing"

func TestTritsSequence_Throw3Dice(t *testing.T) {
	sequence := NewTritsSequence([]byte{1, 2, 3, 2, 1, 3})

	t.Run("Seq 1", func(t *testing.T) {
		got := sequence.Throw3Dice()
		if got != 1 {
			t.Errorf("TritsSequence.Throw3Dice() = %v, want %v", got, 1)
		}
	})

	t.Run("Seq 4", func(t *testing.T) {
		sequence.Throw3Dice()
		sequence.Throw3Dice()
		got := sequence.Throw3Dice()
		if got != 2 {
			t.Errorf("TritsSequence.Throw3Dice() = %v, want %v", got, 2)
		}
	})

	t.Run("Seq 7", func(t *testing.T) {
		sequence.Throw3Dice()
		sequence.Throw3Dice()
		got := sequence.Throw3Dice()
		if got != 1 {
			t.Errorf("TritsSequence.Throw3Dice() = %v, want %v", got, 1)
		}
	})
}
