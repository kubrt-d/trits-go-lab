package tritslab

import (
	"testing"
)

func TestTritsTable_GetCityName(t *testing.T) {
	type fields struct {
		Desk []*TritsGame
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Get Karachi",
			args: args{
				index: 10,
			},
			want: "Karachi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TritsTable{
				Desk: tt.fields.Desk,
			}
			if got := tr.GetCityName(tt.args.index); got != tt.want {
				t.Errorf("TritsTable.GetCityName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTritsTable_GetCityAddress(t *testing.T) {
	tr := NewTritsTable()
	got := tr.GetCityAddress(22)
	if !got.SameAs("1000022600000000100000001000000010000022") {
		t.Errorf("TritsTable.GetCityAddress() = %v, want %v", got,
			"10000022600000000100000001000000010000022")
	}

}
