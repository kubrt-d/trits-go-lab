package tritslab

/* Used for generating random throws of a three-sided dice */

type Random3Dice interface {
	Throw3Dice() byte
}

type TritsDice struct {
}

func NewTritsDice() *TritsDice {
	dice := new(TritsDice)
	return dice
}

func (s *TritsDice) Throw3Dice() byte {
	//b := RandByte() -- use this for cryptosecure generator
	b := FastRandByte()
	return byte(b%3 + 1)
}
