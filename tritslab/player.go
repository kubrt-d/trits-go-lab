package tritslab

import (
	"math"
)

type TritsPlayerResponse struct {
	PlayerAddr *TritsAddress
	Game       *TritsGame
	Amount     uint64
}

func NewTritsPlayerResponse(game *TritsGame, amount uint64, my_addr *TritsAddress) *TritsPlayerResponse {
	r := new(TritsPlayerResponse)
	r.PlayerAddr = my_addr
	r.Game = game
	r.Amount = amount
	return r
}

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

// Composition of the base TritsPlayer
func NewTritsPlayerFactory(addr *TritsAddress, banker *TritsBanker, strategy string) *TritsPlayer {
	switch strategy {
	case "dumb":
		return NewTritsPlayer(addr, banker)
	}
	return nil
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
	return uint64(math.Round(float64(percent) / 100 * float64(pocket)))
}

func (p *TritsPlayer) Balance() uint64 {
	return p.banker.Tell(p.Addr)
}

// Bet or not - some strategy should be implemented here
func (p *TritsPlayer) Bet(desk []*TritsGame) []*TritsPlayerResponse {
	var responses []*TritsPlayerResponse
	pocket := p.Balance()
	if pocket == 0 { // No money no honey
		return nil
	}

	// Try to find an uninitiated game
	var i int = 0
	for i < GAMES_ON_TABLE && desk[i].Nominal != 0 {
		i++
	}
	if i < GAMES_ON_TABLE { // Found an uninitiated game
		amount := p.ChooseAmount()
		res := NewTritsPlayerResponse(desk[i], amount, p.Addr)
		responses = append(responses, res)
		return responses
	}

	// No uninitiated game found, try to find the inbalance of 2 which I can afford
	var j int = 0
	for j < GAMES_ON_TABLE && desk[j].GetInbalance() != 2 && desk[j].Nominal < pocket {
		j++
	}
	if j < GAMES_ON_TABLE { // Found a good inbalance, bet
		res := NewTritsPlayerResponse(desk[j], desk[j].Nominal, p.Addr)
		responses = append(responses, res)
		return responses
	}

	// No good inbalance found, bet on the first one I can afford
	var k int = 0
	for k < GAMES_ON_TABLE && desk[k].Nominal > pocket {
		k++
	}
	if k < GAMES_ON_TABLE { // Found something
		res := NewTritsPlayerResponse(desk[k], desk[k].Nominal, p.Addr)
		responses = append(responses, res)
		return responses
	}

	// I seem to be out of game
	return nil
}
