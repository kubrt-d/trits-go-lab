package tritslab

import (
	"crypto/rand"
)

/* Used for generating random throws of a three-sided dice */

type Random3Dice interface {
	Throw3Dice() int8
}

type TritsDice struct {
}

func NewTritsDice() *TritsDice {
	dice := new(TritsDice)
	return dice
}

func (s *TritsDice) Throw3Dice() int8 {
	b := []byte{0}
	if _, err := rand.Reader.Read(b); err != nil {
		panic(err)
	}
	return int8(b[0]%3 + 1)
}
