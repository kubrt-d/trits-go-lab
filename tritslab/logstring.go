package tritslab

import (
	"fmt"
	"log"
)

const LOG_DEBUG int8 = -1
const LOG_NOTICE int8 = 0
const LOG_INFO int8 = 1
const LOG_WARN int8 = 2
const LOG_ERROR int8 = 3
const LOG_PANIC int8 = 4

// Export function l
func L(level int8, a ...interface{}) {
	l(level, a...)
}

// Log string if allowed in config
func l(level int8, a ...interface{}) {
	if level >= LOG_LEVEL {
		log.Println(fmt.Sprint(a...))
	}
}

// Game status formatter for logging
// Format: <GameName>: <Middle>(<Nominal>)/(<V1>,<V2>,<V3>)=<Inbalance>
func LGame(game *TritsGame) string {

	//Example: "Kairo: 6(333)/(1,2,4)=2"
	return fmt.Sprint(GameName(game.ThisGame),
		": ", game.Middle,
		"(", game.Nominal, ")/(",
		len(game.Trit.V1), ",",
		len(game.Trit.V2), ",",
		len(game.Trit.V3), ")=",
		game.GetInbalance())
}

// Helper function for translating adresses to names (player or game)
func LogName(a TritsAddress) string {
	addr := a.Raw()
	var name string = ""
	// Is this a game address or a player address ?
	if GameName(NewTritsAddress(addr)) != NewTritsAddress(addr).Human() {
		name = GameName(NewTritsAddress(addr))
	} else {
		name = PlayerName(NewTritsAddress(addr))
	}
	return name
}

// CSV column headers - players
func LogPlayersHeaders() string {
	var out string = ""
	s := NewTritsSquad(NewTritsBanker(0))
	l := len(s.squad) - 1
	for k, p := range s.squad {
		out += fmt.Sprint(LogName(p.GetAddr()))
		if k < l {
			out += ","
		} else {
			out += fmt.Sprintln("")
		}
	}
	return out
}

// CSV column headers - players
func LogPlayersBalances(banker TritsBanker) string {
	var out string = ""
	s := NewTritsSquad(banker)
	l := len(s.squad) - 1
	for k, p := range s.squad {
		out += fmt.Sprint(banker.Tell(p.GetAddr()))
		if k < l {
			out += ","
		} else {
			out += fmt.Sprintln("")
		}
	}
	return out
}

func LogZionTotal(banker TritsBanker) uint64 {
	total := banker.Tell(NewTritsAddress(BankAddr))
	s := NewTritsSquad(banker)
	for _, p := range s.squad {
		if p.GetPlayerType() == "zion" {
			total += banker.Tell(p.GetAddr())
		}
	}
	return total
}


// Human readable game names fo logging purposes
func GameName(a TritsAddress) string {
	t := TritsTable{}
	var i int = 0
	for i < GAMES_ON_TABLE {
		if a.Equals(t.GetCityAddress(i)) {
			return t.GetCityName(i)
		}
		i++
	}
	return a.Human()
}
