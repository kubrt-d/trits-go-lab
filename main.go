package main

import (
	"fmt"
	. "trits/tritslab"
)

func main() {
	fmt.Print("ROUND,Bank,")
	LogPlayersHeaders()
	playsome()
}

func playsome() {
	c := NewTritsCroupier(10000, 1000, 0)
	var j = 0
	for c.AskAround() && j < 1000000 {
		fmt.Print(j)
		fmt.Print(",")
		fmt.Print(c.Banker.Tell(NewTritsAddress(BankAddr)))
		fmt.Print(",")
		LogPlayersBalances(c.Banker)
		j++
	}
	/*
		res := c.Banker.Tell(NewTritsAddress(BankAddr))
		fmt.Println(fmt.Sprint(t, ",", res))
		fmt.Println(c.Banker.DumpBank())

		for _, game := range c.Table.Desk {
			fmt.Println(LGame(game))
		}
	*/
}
