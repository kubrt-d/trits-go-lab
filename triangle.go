package tritslab

const outweight int8 = 3

type TritsTriangle struct {
	V1, V2, V3 int
}

func NewTritsTriangle(v1 int, v2 int, v3 int) *TritsTriangle {
	triangle := new(TritsTriangle)
	triangle.V1 = v1
	triangle.V2 = v2
	triangle.V3 = v3
	return triangle
}

/*
* HitVertice (vertice) - place coin on one of the vertices (1-3), return true if its a victory, false otherwise.
* The victory occurs when the number of coins on the vertice that has just been hit outnumbers the
* number of coins on any other vertice by the value of "outweight" const (default = 3)
 */

func (t *TritsTriangle) HitVertice(vertice int8) bool {
	switch vertice {
	case 1:
		t.V1++
		if t.V1-t.V2 >= int(outweight) {
			return true
		}
		if t.V1-t.V3 >= int(outweight) {
			return true
		}
	case 2:
		t.V2++
		if t.V2-t.V1 >= int(outweight) {
			return true
		}
		if t.V2-t.V3 >= int(outweight) {
			return true
		}
	case 3:
		t.V3++
		if t.V3-t.V1 >= int(outweight) {
			return true
		}
		if t.V3-t.V2 >= int(outweight) {
			return true
		}
	}
	return false
}
