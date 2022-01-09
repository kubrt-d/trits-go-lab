package tritslab

import (
	"fmt"
	"testing"
)

func TestTritsCroupier_Constructor(t *testing.T) {
	t.Run("Test croupier constructor", func(t *testing.T) {
		c := NewTritsCroupier(1000000, 50000, 50000)
		games := len(c.Table.Desk)
		if games != GAMES_ON_TABLE {
			t.Errorf("TritsCroupier.NewTritsCroupier().table.desk got = %v, want %v", games, GAMES_ON_TABLE)
		}

		players := len(c.Players.squad)
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
		c := NewTritsCroupier(1000000, 50000,50000)
		c.AskAround()
		var j = 0
		for j < 2000 {
			c.AskAround()
			j++
		}
		c.AskAround()
		res := c.Banker.DumpBank()
		fmt.Println(res)
	})

}
