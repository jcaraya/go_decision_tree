package dataset

import (
	"reflect"
	"testing"
)

func TestMakeRange(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Negative",
			args: args{-1},
			want: []int{},
		},
		{
			name: "Zero",
			args: args{0},
			want: []int{},
		},
		{
			name: "Success",
			args: args{3},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeRange(tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
