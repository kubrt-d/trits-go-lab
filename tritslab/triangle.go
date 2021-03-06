package tritslab

const OUTWEIGHT uint8 = 3

type TritsTriangle struct {
	V1, V2, V3 []TritsAddress
}

func NewTritsTriangle() *TritsTriangle {
	triangle := new(TritsTriangle)
	triangle.V1 = make([]TritsAddress, 0, 15)
	triangle.V2 = make([]TritsAddress, 0, 15)
	triangle.V3 = make([]TritsAddress, 0, 15)
	return triangle
}

/*
* HitVertice (vertice) - place coin on one of the vertices (1-3), return true if its a victory, false otherwise.
* The victory occurs when the number of coins on the vertice that has just been hit outnumbers the
* number of coins on any other vertice by the value of "outweight" const (default = 3)
 */
func (t *TritsTriangle) HitVertice(vertice byte, by TritsAddress) bool {
	switch vertice {
	case 1:
		t.V1 = append(t.V1, by)
	case 2:
		t.V2 = append(t.V2, by)
	case 3:
		t.V3 = append(t.V3, by)
	}
	return t.Inbalance() == OUTWEIGHT
}

// Returns the max difference between arms [0,1 or 2] (or even more)
func (t *TritsTriangle) Inbalance() uint8 {
	v1 := uint8(len(t.V1))
	v2 := uint8(len(t.V2))
	v3 := uint8(len(t.V3))
	min := v1
	max := v1
	if v2 < min {
		min = v2
	}
	if v2 > max {
		max = v2
	}
	if v3 < min {
		min = v3
	}
	if v2 > max {
		max = v3
	}
	return max - min
}
