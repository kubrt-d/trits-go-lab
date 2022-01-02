package tritslab

const outweight int8 = 3

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

func (t *TritsTriangle) HitVertice(vertice int8, by *TritsAddress) bool {
	switch vertice {
	case 1:
		t.V1 = append(t.V1, by)
		if len(t.V1)-len(t.V2) >= int(outweight) {
			return true
		}
		if len(t.V1)-len(t.V3) >= int(outweight) {
			return true
		}
	case 2:
		t.V2 = append(t.V2, by)
		if len(t.V2)-len(t.V1) >= int(outweight) {
			return true
		}
		if len(t.V2)-len(t.V3) >= int(outweight) {
			return true
		}
	case 3:
		t.V3 = append(t.V3, by)
		if len(t.V3)-len(t.V1) >= int(outweight) {
			return true
		}
		if len(t.V3)-len(t.V2) >= int(outweight) {
			return true
		}
	}
	return false
}
