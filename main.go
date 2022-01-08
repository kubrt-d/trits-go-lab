package main

import (
	"fmt"
	. "trits/tritslab"
)

func main() {

	var t byte = 0
	for t < 1 {
		t++
		playsome(t)
	}
}

func playsome(t byte) {
	c := NewTritsCroupier(1000000, 0, t)
	var j = 0
	for j < 10000 {
		c.AskAround()
		fmt.Println(c.Banker.Tell(NewTritsAddress(BankAddr)))
		/*
			if c.Banker.Tell(NewTritsAddress(NeoAddr)) <= 1 {
				fmt.Println(fmt.Sprint("Neo broke after ", j, " rounds"))
				break
			}
		*/
		j++
	}
	res := c.Banker.Tell(NewTritsAddress(BankAddr))
	fmt.Println(fmt.Sprint(t, ",", res))
	fmt.Println(c.Banker.DumpBank())

	for _, game := range c.Table.Desk {
		fmt.Println(LGame(game))
	}

}
