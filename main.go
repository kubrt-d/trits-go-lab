package main

import (
	"fmt"
	. "trits/tritslab"
)

func main() {

	var t byte = 0
	for t < 100 {
		t++
		playsome(t)
	}
}

func playsome(t byte) {
	c := NewTritsCroupier(1000, 1000, t)
	c.AskAround()
	var j = 0
	for j < 100000 {
		c.AskAround()
		if c.Banker.Tell(NewTritsAddress(NeoAddr)) <= 1 {
			fmt.Println(fmt.Sprint("Neo broke after ", j, " rounds"))
			break
		}
		j++
	}
	res := c.Banker.Tell(NewTritsAddress(BankAddr))
	fmt.Println(fmt.Sprint(t, ",", res))
	//fmt.Println(c.Banker.DumpBank())
}
