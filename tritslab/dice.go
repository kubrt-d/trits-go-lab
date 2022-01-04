package tritslab

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
	b := RandByte()
	return int8(b%3 + 1)
}
