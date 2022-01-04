package tritslab

import (
	"math"
)

type TritsPlayer struct {
	Addr   *TritsAddress // Player's address
	banker *TritsBanker  // Connection to the banker
}

func NewTritsPlayer(addr *TritsAddress, banker *TritsBanker) *TritsPlayer {
	player := new(TritsPlayer)
	player.banker = banker
	player.Addr = addr
	return player
}

// Choose the table
func (p *TritsPlayer) ChooseTable() int8 {
	b := RandByte()
	return int8(b % 23)
}

// Choose amount
func (p *TritsPlayer) ChooseAmount() uint64 {
	b := RandByte()
	pocket := p.Balance()
	percent := int8(b%10 + 1) // Bet no more between 1 - 10 percent of the pocket
	return uint64(math.Round(float64(percent) / 10000 * float64(pocket)))
}

func (p *TritsPlayer) Balance() uint64 {
	return p.banker.Tell(p.Addr)
}

// Play or not ?
func (p *TritsPlayer) PlayOrNot(nominal uint64) bool {
	pocket := p.Balance()
	b := RandByte()
	percent := int8(b%30 + 1) // 1-30 percent of my pocket is ok
	if pocket > nominal && float64(nominal) < (float64(percent)/100*float64(pocket)) {
		return true
	} else {
		return false
	}
}
