package tritslab

/* Used for generating predefined sequence of numbers instead of
 * random numbers, useful for running tests
 */

type TritsSequence struct {
	pos int
	seq []int8
	len int
}

func NewTritsSequence(seq []int8) *TritsSequence {
	sequence := new(TritsSequence)
	sequence.pos = 0
	sequence.seq = seq
	sequence.len = len(seq)
	return sequence
}

func (s *TritsSequence) Throw3Dice() int8 {
	if s.pos >= s.len { // Loop through the sequence if necessary (the sequence length is shorther than the number of calls)
		s.pos = 0
	}
	var out int8 = s.seq[s.pos]
	s.pos++
	return out
}
