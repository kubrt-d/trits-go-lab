package tritslab

import (
	"fmt"
)

type TritsBanker struct {
	bank         map[string]uint64 // Bank is a map of account addresses and their current balances
	started_with uint64
}

func NewTritsBanker(start_with uint64) *TritsBanker {
	banker := new(TritsBanker)
	banker.started_with = start_with
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
		return false, "Sender " + from.Human() + " doesn't exist"
	} else {
		if b.bank[f] < amount {
			return false, "Insufficient funds, sender " + from.Human() + " has got " + fmt.Sprint(b.bank[f]) + " only"
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
