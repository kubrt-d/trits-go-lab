package tritslab

import (
	"math/rand"
	"testing"
	"time"
)

func TestTritsTriangle_HitVertice(t *testing.T) {

	var foo_addr = NewTritsAddress("1111111111111111111111111111111111111111")

	t.Run("Try to win", func(t *testing.T) {
		tr := NewTritsTriangle()
		tr.HitVertice(1, foo_addr)
		tr.HitVertice(2, foo_addr)
		tr.HitVertice(3, foo_addr)
		tr.HitVertice(1, foo_addr)
		tr.HitVertice(1, foo_addr)
		got := tr.HitVertice(1, foo_addr)
		if got != true {
			t.Errorf("TritsTriangle.HitVertice() = %v, want %v", got, true)
		}
	})

	t.Run("Try not to win", func(t *testing.T) {
		tr := NewTritsTriangle()
		tr.HitVertice(1, foo_addr)
		tr.HitVertice(2, foo_addr)
		tr.HitVertice(3, foo_addr)
		tr.HitVertice(1, foo_addr)
		tr.HitVertice(4, foo_addr)
		got := tr.HitVertice(1, foo_addr)
		if got != false {
			t.Errorf("TritsTriangle.HitVertice() = %v, want %v", got, false)
		}
	})

	t.Run("Should not run too many rounds", func(t *testing.T) {
		tr := NewTritsTriangle()
		counter := 0
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		for counter < 100 { // 100 rounds would be extremely rare
			counter++
			v := byte(r1.Intn(3)) + 1
			if tr.HitVertice(v, foo_addr) {
				counter = 111
			}
		}
		if counter != 111 {
			t.Errorf("TritsTriangle.HitVertice() = %v, want %v", counter, 111)
		}
	})

}

func BenchmarkTritsTriangle_HitVertice(b *testing.B) {
	var foo_addr = NewTritsAddress("1111111111111111111111111111111111111111")
	tr := NewTritsTriangle()
	for n := 0; n < b.N; n++ {
		tr.HitVertice(1, foo_addr)
	}
}

func BenchmarkTritsTriangle_Inbalance(b *testing.B) {
	var foo_addr = NewTritsAddress("1111111111111111111111111111111111111111")
	tr := NewTritsTriangle()
	tr.HitVertice(1, foo_addr)
	tr.HitVertice(1, foo_addr)
	tr.HitVertice(2, foo_addr)
	tr.HitVertice(3, foo_addr)
	tr.HitVertice(3, foo_addr)
	for n := 0; n < b.N; n++ {
		tr.Inbalance()
	}
}
