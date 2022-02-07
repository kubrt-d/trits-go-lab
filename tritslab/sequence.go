package tritslab

/* Used for generating predefined sequence of numbers instead of
 * random numbers, useful for running tests
 */

type TritsSequence struct {
	pos int
	seq []byte
	len int
}

func NewTritsSequence(seq []byte) *TritsSequence {
	sequence := new(TritsSequence)
	sequence.pos = 0
	sequence.seq = seq
	sequence.len = len(seq)
	return sequence
}

func (s *TritsSequence) Throw3Dice() byte {
	if s.pos >= s.len { // Loop through the sequence if necessary (the sequence length is shorther than the number of calls)
		s.pos = 0
	}
	var out byte = s.seq[s.pos]
	s.pos++
	return out
}
