package main

import (
	"fmt"
	tl "trits/tritslab"
)

func main() {

	/*
		// open log file
		logFile, err := os.OpenFile(tl.LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Panic(err)
		}
		defer logFile.Close()
	*/
	fmt.Print("ROUND,Bank,Zion," + tl.LogPlayersHeaders())
	playsome()
}

func playsome() {
	c := tl.NewTritsCroupier(140000, 10000, 10000)
	var j = 0
	var cont = true
	for cont && j < 1000000 {
		//tl.L(tl.LOG_DEBUG, "Round:", j)
		fmt.Print(fmt.Sprint(j, ",", c.Banker.Tell(tl.NewTritsAddress(tl.BankAddr)), ",", tl.LogZionTotal(c.Banker), ",", tl.LogPlayersBalances(c.Banker)))
		j++
		cont = c.AskAround()

	}
	fmt.Println(c.Banker.DumpBank())
	/*
		res := c.Banker.Tell(NewTritsAddress(BankAddr))
		fmt.Println(fmt.Sprint(t, ",", res))
		fmt.Println(c.Banker.DumpBank())

		for _, game := range c.Table.Desk {
			fmt.Println(LGame(game))
		}
	*/
}
