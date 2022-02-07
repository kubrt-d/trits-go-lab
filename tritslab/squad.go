package tritslab

import (
	"math/rand"
	"time"
)

const LenderAddr string = "0000000000000000000000000000000000000000"
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

const NoAddr string = "0101010101010101010101010101010101010101"

// TritsSquad is a slice of *TritsPlayer
type TritsSquad struct {
	squad []TritsPlayer // Our dear players
}

func NewTritsSquad(banker TritsBanker) TritsSquad {
	s := TritsSquad{}
	var all [14]TritsPlayer
	all[0] = NewTritsPlayerFactory(NewTritsAddress(NeoAddr), banker, "sharp")
	all[1] = NewTritsPlayerFactory(NewTritsAddress(TrinityAddr), banker, "zion")
	all[2] = NewTritsPlayerFactory(NewTritsAddress(KeymakerAddr), banker, "zion")
	all[3] = NewTritsPlayerFactory(NewTritsAddress(MorpheusAddr), banker, "zion")
	all[4] = NewTritsPlayerFactory(NewTritsAddress(OracleAddr), banker, "zion")
	all[5] = NewTritsPlayerFactory(NewTritsAddress(SeraphAddr), banker, "zion")
	all[6] = NewTritsPlayerFactory(NewTritsAddress(BaneAddr), banker, "zion")
	all[7] = NewTritsPlayerFactory(NewTritsAddress(AgentAddr), banker, "agent")
	all[8] = NewTritsPlayerFactory(NewTritsAddress(NiobeAddr), banker, "dumb")
	all[9] = NewTritsPlayerFactory(NewTritsAddress(PersephoneAddr), banker, "dumb")
	all[10] = NewTritsPlayerFactory(NewTritsAddress(TwinsAddr), banker, "dumb")
	all[11] = NewTritsPlayerFactory(NewTritsAddress(BugsAddr), banker, "dumb")
	all[12] = NewTritsPlayerFactory(NewTritsAddress(AnalystAddr), banker, "dumb")
	all[13] = NewTritsPlayerFactory(NewTritsAddress(ArchitectAddr), banker, "dumb")

	s.squad = all[0:PLAYERS_IN_SQUAD]
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
