package tritslab

import (
	"fmt"
)

const LOG_DEBUG int8 = -1
const LOG_NOTICE int8 = 0
const LOG_INFO int8 = 1
const LOG_WARN int8 = 2
const LOG_ERROR int8 = 3
const LOG_PANIC int8 = 4

// Log string if allowed in config
func l(level int8, a ...interface{}) {
	if level >= LOG_LEVEL {
		fmt.Printf("%s \n", fmt.Sprint(a...))
	}
}

// Game status formatter for logging
// Format: <GameName>: <Middle>(<Nominal>)/(<V1>,<V2>,<V3>)=<Inbalance>
func lgame(game *TritsGame) string {

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
func LogName(a *TritsAddress) string {
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
