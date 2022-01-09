package tritslab

import (
	"encoding/hex"
	"log"
	"strings"
)

// TODO: Implement real IOTA address

/* For the purposes of the lab, we use string of hex 40 bytes long
	i.e. e28533750bee16842a5cd4f533d235770
   and
	e4073670
   for human readable
*/

type TritsAddress struct {
	addr string
}

// Constructor - validate string and create address struct
func NewTritsAddress(a string) *TritsAddress {
	address := new(TritsAddress)

	/* Validate the input string */
	dst := make([]byte, hex.DecodedLen(len(a)))
	n, err := hex.Decode(dst, []byte(a))
	//TODO:  Maybe doesn't have to be fatal
	if err != nil {
		log.Fatal(err)
	}
	if n != 20 {
		log.Fatal("Address too short")
	} else {
		address.addr = a
	}
	return address
}

// Return address as a raw string
func (a *TritsAddress) Raw() string {
	if a != nil {
		return a.addr
	} else {
		return "n/a"
	}
}

// Return address is a human readable form
func (a *TritsAddress) Human() string {
	if a != nil {
		return string(a.addr[32:40])
	} else {
		return "n/a"
	}
}

// Compare two addresses 
func (a *TritsAddress) Equals(to *TritsAddress) bool {
	if to == nil && a == nil {
		return true
	}
	if to == nil || a == nil {
		return false
	}
	if strings.Compare(a.addr, to.addr) == 0 {
		return true
	} else {
		return false
	}
}

// Compare address with a raw string 
func (a *TritsAddress) SameAs(as string) bool {
	if strings.Compare(string(a.addr), as) == 0 {
		return true
	} else {
		return false
	}
}
