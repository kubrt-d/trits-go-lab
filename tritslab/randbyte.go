package tritslab

import (
	cr "crypto/rand"
	mr "math/rand"
)

func RandByte() byte {
	b := []byte{0}
	if _, err := cr.Reader.Read(b); err != nil {
		panic(err)
	}
	return b[0]
}

func FastRandByte() byte {
	return byte(mr.Intn(255))
}
