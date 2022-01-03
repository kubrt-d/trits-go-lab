package tritslab

import "testing"

func TestTritsTable_GetCityName(t *testing.T) {
	type fields struct {
		table []*TritsGame
	}
	type args struct {
		index int8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Get Karachi",
			args: args {
				index: 10,
			},
			want: "Karachi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TritsTable{
				table: tt.fields.table,
			}
			if got := tr.GetCityName(tt.args.index); got != tt.want {
				t.Errorf("TritsTable.GetCityName() = %v, want %v", got, tt.want)
			}
		})
	}
}
