package tritslab

import (
	"fmt"
)

type TritsBanker struct {
	bank         map[string]uint64 // Bank is a map of account addresses and their current balances
	started_with uint64
	treshold byte
}

func NewTritsBanker(start_with uint64, treshold byte) *TritsBanker {
	banker := new(TritsBanker)
	banker.started_with = start_with
	banker.treshold = treshold
	banker.bank = make(map[string]uint64)
	banker.bank[BankAddr] = start_with
	return banker
}

// Check if the bank's still holds all the money
func (b *TritsBanker) healthcheck() bool {
	var sum uint64 = 0
	for _, balance := range b.bank {
		sum += balance
	}
	if sum != b.started_with {
		panic("The Bank is leaky or corrupted !")
	} else {
		return true
	}
}

func (b *TritsBanker) gamehealthcheck(game *TritsGame) bool {
	should_have := uint64(game.Middle + uint32(len(game.Trit.V1)+len(game.Trit.V2)+len(game.Trit.V3)))
	has := b.Tell(game.ThisGame)
	if should_have != has {
		l(LOG_PANIC, "Wrong game balance ", game.ThisGame.Human(), " has ", has, ", while it should have ", should_have, ".Exiting!")
	} else {
		return true
	}
	return false
}

// Tell balance for address
func (b *TritsBanker) Tell(who *TritsAddress) uint64 {
	w := who.Raw()
	if amount, ok := b.bank[w]; !ok {
		return 0
	} else {
		return amount
	}
}

// Move funds from one address/account to another
func (b *TritsBanker) MoveFunds(from *TritsAddress, to *TritsAddress, amount uint64) (bool, string) {
	// The from address must exist
	f := from.Raw()
	t := to.Raw()
	if _, ok := b.bank[f]; !ok {
		return false, "Sender " + LogName(from) + " doesn't exist"
	} else {
		if b.bank[f] < amount {
			return false, "Sender " + LogName(from) + " has got " + fmt.Sprint(b.bank[f]) + " only, reqesting " + fmt.Sprint(amount)
		}
	}
	if _, ok := b.bank[t]; !ok { // Create a new account if it doesn't exist
		b.bank[t] = 0
	}

	// Move the funds
	b.bank[f] -= amount
	b.bank[t] += amount
	b.healthcheck()
	return true, "OK, moved " + fmt.Sprint(amount) + " from sender " + from.Human() + " to " + to.Human()
}

// Determines bonus, moves funds to the game and tell croupier how much for the PlaceCoin
func (b *TritsBanker) PutBonus(game *TritsGame) uint64 {
	in_bank := b.Tell(NewTritsAddress(BankAddr))
	var bonus uint64 = 0
	if game.Nominal*4 > in_bank { // Can't give any bonus
		return 0
	}
	if game.Nominal*5 > in_bank {
		bonus = game.Nominal * 4
		b.MoveFunds(NewTritsAddress(BankAddr), game.ThisGame, bonus)
		return bonus
	}
	// 4 or 5, THAT is THE question
	r := RandByte() % 100

	if r >= byte(b.treshold) {
		bonus = 5 * game.Nominal
	} else {
		bonus = 4 * game.Nominal
	}
	b.MoveFunds(NewTritsAddress(BankAddr), game.ThisGame, bonus)
	//TODO: check if the money has really moved
	return bonus
}

// Dump all balances for logging purposes
func (b *TritsBanker) DumpBank() string {
	// TODO: Sort them alphabetically

	var out string = ""
	for addr, balance := range b.bank {
		out = out + LogName(NewTritsAddress(addr)) + ":" + fmt.Sprint(balance) + " "
	}
	return out
}
