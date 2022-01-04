package tritslab

import (
	"fmt"
)

const LOG_NOTICE byte = 0
const LOG_INFO byte = 1
const LOG_WARN byte = 2
const LOG_ERROR byte = 3
const LOG_PANIC byte = 4

// Log string if allowed in config
func l(msg string, level byte) {
	if level >= LOG_LEVEL {
		fmt.Printf("%s \n", msg)
	}
}
