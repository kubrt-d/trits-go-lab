package main

import (
	"fmt"
	. "trits/tritslab"
)

func main() {
	var t byte = 0
	for t < 100 {
		t++
		c := NewTritsCroupier(140000, 10000, t)
		c.AskAround()
		var j = 0
		for j < 1000000 {
			c.AskAround()
			j++
		}
		res := c.Banker.Tell(NewTritsAddress(BankAddr))
		fmt.Println(fmt.Sprint(t, res))
	}
}
