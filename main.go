package main

import (
	"fmt"
	mr "math/rand"
	"time"
	tl "trits/tritslab"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

func main() {
	// Initialize the random number generator
	mr.Seed(time.Now().UnixNano())
	// Initialize INFLUX connection
	client := influxdb2.NewClient("http://"+tl.INFLUX_HOST+":"+tl.INFLUX_PORT, tl.INFLUX_API_KEY)
	influxWriteAPI := client.WriteAPI("Trits", "tritslab")

	defer client.Close()
	/*
		// open log file
		logFile, err := os.OpenFile(tl.LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Panic(err)
		}
		defer logFile.Close()
	*/
	fmt.Print("ROUND,Bank,Zion," + tl.LogPlayersHeaders())
	playsome(influxWriteAPI)
}

func playsome(infWAPI api.WriteAPI) {
	c := tl.NewTritsCroupier(140000, 10000, 10000)
	var j = 0
	var cont = true
	for cont && j < 1000000 {
		//tl.L(tl.LOG_DEBUG, "Round:", j)
		bank_balance := c.Banker.Tell(tl.NewTritsAddress(tl.BankAddr))
		zion_balance := tl.LogZionTotal(c.Banker)
		now := time.Unix(0, time.Now().UnixNano())
		p := influxdb2.NewPoint("balance",
			map[string]string{"unit": "balance"},
			map[string]interface{}{"bank": bank_balance, "zion": zion_balance},
			now,
		)
		infWAPI.WritePoint(p)
		fmt.Print(fmt.Sprint(j, ",", bank_balance, ",", zion_balance, ",", tl.LogPlayersBalances(c.Banker)))
		j++
		cont = c.AskAround()

	}
	infWAPI.Flush()
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
