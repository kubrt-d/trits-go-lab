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

type rawaddr []byte
type humanaddr []byte

type TritsAddress struct {
	addr rawaddr
}

func NewTritsAddress(a rawaddr) *TritsAddress {
	address := new(TritsAddress)

	/* Validate the input string */
	dst := make([]byte, hex.DecodedLen(len(a)))
	n, err := hex.Decode(dst, a)
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

func (a *TritsAddress) Raw() rawaddr {
	return a.addr
}

func (a *TritsAddress) Human() humanaddr {
	return humanaddr(a.addr[32:40])
}

func (a *TritsAddress) SameAs(as string) bool {
	if strings.Compare(string(a.addr), as) == 0 {
		return true
	} else {
		return false
	}
}
