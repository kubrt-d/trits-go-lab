package tritslab

import (
	"crypto/rand"
)

func RandByte() byte {
	b := []byte{0}
	if _, err := rand.Reader.Read(b); err != nil {
		panic(err)
	}
	return b[0]
}