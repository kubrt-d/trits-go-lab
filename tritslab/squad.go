package tritslab

import (
	"math/rand"
	"time"
)

const NoAddr string = "0000000000000000000000000000000000000000"
const BankAddr string = "ffffffffffffffffffffffffffffffffffffffff"
const NeoAddr string = "1111111111111111111111111111111111111111"
const TrinityAddr string = "2222222222222222222222222222222222222222"
const AgentAddr string = "3333333333333333333333333333333333333333"
const KeymakerAddr string = "4444444444444444444444444444444444444444"
const MorpheusAddr string = "5555555555555555555555555555555555555555"
const NiobeAddr string = "6666666666666666666666666666666666666666"
const OracleAddr = "7777777777777777777777777777777777777777"
const PersephoneAddr = "8888888888888888888888888888888888888888"
const TwinsAddr = "9999999999999999999999999999999999999999"
const BugsAddr = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
const AnalystAddr = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
const SeraphAddr = "cccccccccccccccccccccccccccccccccccccccc"
const ArchitectAddr = "dddddddddddddddddddddddddddddddddddddddd"
const BaneAddr = "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"

type TritsSquad struct {
	squad [14]*TritsPlayer // Our dear players
}

func NewTritsSquad(banker *TritsBanker) *TritsSquad {
	s := new(TritsSquad)
	s.squad[0] = NewTritsPlayer(NewTritsAddress(NeoAddr), banker)
	s.squad[1] = NewTritsPlayer(NewTritsAddress(TrinityAddr), banker)
	s.squad[2] = NewTritsPlayer(NewTritsAddress(AgentAddr), banker)
	s.squad[3] = NewTritsPlayer(NewTritsAddress(KeymakerAddr), banker)
	s.squad[4] = NewTritsPlayer(NewTritsAddress(MorpheusAddr), banker)
	s.squad[5] = NewTritsPlayer(NewTritsAddress(NiobeAddr), banker)
	s.squad[6] = NewTritsPlayer(NewTritsAddress(OracleAddr), banker)
	s.squad[7] = NewTritsPlayer(NewTritsAddress(PersephoneAddr), banker)
	s.squad[8] = NewTritsPlayer(NewTritsAddress(TwinsAddr), banker)
	s.squad[9] = NewTritsPlayer(NewTritsAddress(BugsAddr), banker)
	s.squad[10] = NewTritsPlayer(NewTritsAddress(BugsAddr), banker)
	s.squad[11] = NewTritsPlayer(NewTritsAddress(AnalystAddr), banker)
	s.squad[12] = NewTritsPlayer(NewTritsAddress(SeraphAddr), banker)
	s.squad[13] = NewTritsPlayer(NewTritsAddress(ArchitectAddr), banker)
	return s
}

// Randomly shuffle the players
func (s *TritsSquad) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s.squad), func(i, j int) { s.squad[i], s.squad[j] = s.squad[j], s.squad[i] })
}

func (s *TritsSquad) SizeOf() int {
	return len(s.squad)
}
