package tritslab

import (
)

type TritsCroupier struct {
	table *TritsTable // Croupier's table
}

func NewTritsCroupier() *TritsCroupier {
	croupier := new(TritsCroupier)
	return croupier
}


// Choose the table
func (p *TritsCroupier) AskAround() int8 {
	return 1
}

