package tritslab

import (
	"testing"
)

func TestTritsCroupier_AskAround(t *testing.T) {
	t.Run("Test croupier constructor", func(t *testing.T) {
		b := NewTritsCroupier(1000000, 50000)
		got := b
		if got != nil {

		}
		/*
			if got != 1000 {
				t.Errorf("TritsBanker.NewTritsBanker() got = %v, want %v", got, 1000)
			}
		*/
	})
}
