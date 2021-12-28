package tritslab

import (
	"math/rand"
	"testing"
	"time"
)

func TestTritsTriangle_HitVertice(t *testing.T) {
	type fields struct {
		V1 int
		V2 int
		V3 int
	}
	type args struct {
		vertice int8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Hit Vertice 1",
			args: args{
				vertice: 1,
			},
			want: false,
		},
		{
			name: "Hit Vertice 2",
			args: args{
				vertice: 2,
			},
			want: false,
		},
		{
			name: "Hit Vertice 3",
			args: args{
				vertice: 3,
			},
			want: false,
		},
		{
			name: "Hit Vertice 127",
			args: args{
				vertice: 127,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TritsTriangle{
				V1: tt.fields.V1,
				V2: tt.fields.V2,
				V3: tt.fields.V3,
			}
			if got := tr.HitVertice(tt.args.vertice); got != tt.want {
				t.Errorf("TritsTriangle.HitVertice() = %v, want %v", got, tt.want)
			}
		})
	}
	
	t.Run("Try to win", func(t *testing.T) {
		tr := NewTritsTriangle(3, 3, 3)
		tr.HitVertice(1)
		tr.HitVertice(2)
		tr.HitVertice(3)
		tr.HitVertice(1)
		tr.HitVertice(1)
		got := tr.HitVertice(1)
		if got != true {
			t.Errorf("TritsTriangle.HitVertice() = %v, want %v", got, true)
		}
	})
	
	t.Run("Try not to win", func(t *testing.T) {
		tr := NewTritsTriangle(3, 3, 3)
		tr.HitVertice(1)
		tr.HitVertice(2)
		tr.HitVertice(3)
		tr.HitVertice(1)
		tr.HitVertice(4)
		got := tr.HitVertice(1)
		if got != false {
			t.Errorf("TritsTriangle.HitVertice() = %v, want %v", got, false)
		}
	})
	
	t.Run("Should not run too many rounds", func(t *testing.T) {
		tr := NewTritsTriangle(0, 0, 0)
		counter := 0
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		for counter < 100 {  // 100 rounds would be extremely rare 
			counter++
			v := int8(r1.Intn(3)) + 1
			if tr.HitVertice(v) {
				counter = 111
			}
		}
		if counter != 111 {
			t.Errorf("TritsTriangle.HitVertice() = %v, want %v", counter, 111)
		}
	})
	
}
