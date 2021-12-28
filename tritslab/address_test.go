package tritslab

import (
	"strings"
	"testing"
)

func TestTritsAddress_Raw(t *testing.T) {

	t.Run("Valid raw address", func(t *testing.T) {
		var valid rawaddr = []byte("e28533750bee16842a5cd4f533d235770e407367")
		a := NewTritsAddress(valid)
		got := a.Raw()
		if strings.Compare(string(got), string(valid)) != 0 {
			t.Errorf("TritsTriangle.Raw() = %v, want %v", got, valid)
		}
	})

	t.Run("Valid human address", func(t *testing.T) {
		var valid rawaddr = []byte("e28533750bee16842a5cd4f533d235770e407367")
		a := NewTritsAddress(valid)
		got := a.Human()
		if strings.Compare(string(got), "0e407367") != 0 {
			t.Errorf("TritsTriangle.Human() = %v, want %v", string(got), "0e407367")
		}
	})

	t.Run("Compare raw string", func(t *testing.T) {
		var valid rawaddr = []byte("e28533750bee16842a5cd4f533d235770e407367")
		var x string = "e28533750bee16842a5cd4f533d235770e407367"
		a := NewTritsAddress(valid)
		got := a.SameAs(x)
		if got {
			t.Errorf("TritsTriangle.SameAs() = %v, want %v", got, true)
		}
	})

}
