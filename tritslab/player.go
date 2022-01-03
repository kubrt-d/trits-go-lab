package tritslab

import (
	"crypto/rand"
	"math"
)

type TritsPlayer struct {
	pocket   uint64         // Amount at player's disposal
	croupier *TritsCroupier // Connection to the croupier
	banker   *TritsBanker   // Connection to the banker
}

func NewTritsPlayer(addr *TritsAddress, banker *TritsBanker, croupier *TritsCroupier) *TritsPlayer {
	player := new(TritsPlayer)
	player.banker = banker
	player.croupier = croupier
	return player
}

// Choose the table
func (p *TritsPlayer) ChooseTable() int8 {
	b := p.RandByte()
	return int8(b % 23)
}

// Choose amount
func (p *TritsPlayer) ChooseAmount() uint64 {
	b := p.RandByte()
	percent := int8(b%10 + 1) // Bet no more between 1 - 10 percent of the pocket
	return uint64(math.Round(float64(percent) / 10000 * float64(p.pocket)))
}

// Play or not ?
func (p *TritsPlayer) PlayOrNot(nominal uint64) bool {
	b := p.RandByte()
	percent := int8(b%30 + 1) // 1-30 percent of my pocket is ok
	if p.pocket > nominal && float64(nominal) < (float64(percent)/100*float64(p.pocket)) {
		return true
	} else {
		return false
	}
}

// Return a random byte
func (p *TritsPlayer) RandByte() byte {
	b := []byte{0}
	if _, err := rand.Reader.Read(b); err != nil {
		panic(err)
	}
	return b[0]
}
