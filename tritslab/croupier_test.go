package tritslab

import (
	"fmt"
	"testing"
)

func TestTritsCroupier_Constructor(t *testing.T) {
	t.Run("Test croupier constructor", func(t *testing.T) {
		c := NewTritsCroupier(1000000, 50000)
		games := len(c.table.desk)
		if games != GAMES_ON_TABLE {
			t.Errorf("TritsCroupier.NewTritsCroupier().table.desk got = %v, want %v", games, GAMES_ON_TABLE)
		}
		players := len(c.players.squad)
		if players != PLAYERS_IN_SQUAD {
			t.Errorf("TritsCroupier.NewTritsCroupier().players.squad got = %v, want %v", players, PLAYERS_IN_SQUAD)
		}
		/*
			if got != 1000 {
				t.Errorf("TritsBanker.NewTritsBanker() got = %v, want %v", got, 1000)
			}
		*/
	})
}

func TestTritsCroupier_AskAround(t *testing.T) {
	t.Run("Test croupier AskAround", func(t *testing.T) {
		c := NewTritsCroupier(1000000, 50000)
		c.AskAround()
		var j = 0
		for j < 20 {
			c.AskAround()
			l(fmt.Sprint(
				c.table.desk[0].Middle,",",
				len(c.table.desk[0].Trit.V1),",",
				len(c.table.desk[0].Trit.V2),",",
				len(c.table.desk[0].Trit.V3)),LOG_INFO)
			j++
		}
		c.AskAround()
	})
}
