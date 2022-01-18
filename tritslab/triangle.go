package tritslab

import (
	"sort"
)

const OUTWEIGHT int8 = 3

type TritsTriangle struct {
	V1, V2, V3 []*TritsAddress
}

func NewTritsTriangle() *TritsTriangle {
	triangle := new(TritsTriangle)
	triangle.V1 = nil
	triangle.V2 = nil
	triangle.V3 = nil
	return triangle
}

/*
* HitVertice (vertice) - place coin on one of the vertices (1-3), return true if its a victory, false otherwise.
* The victory occurs when the number of coins on the vertice that has just been hit outnumbers the
* number of coins on any other vertice by the value of "outweight" const (default = 3)
 */
func (t *TritsTriangle) HitVertice(vertice int8, by TritsAddress) bool {
	switch vertice {
	case 1:
		t.V1 = append(t.V1, &by)
	case 2:
		t.V2 = append(t.V2, &by)
	case 3:
		t.V3 = append(t.V3, &by)
	}
	return t.Inbalance() == OUTWEIGHT
}

// Returns the max difference between arms [0,1 or 2] (or even more)
func (t *TritsTriangle) Inbalance() int8 {
	s := []int{len(t.V1), len(t.V2), len(t.V3)}
	sort.Ints(s)
	return int8(s[2] - s[0])
}
