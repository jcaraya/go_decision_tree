package dataset

import (
	"reflect"
	"sort"
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
			if got := makeRange(tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeShuffledRange(t *testing.T) {
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
			args: args{5},
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeShuffledRange(tt.args.max)
			sort.Ints(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeShuffledRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleDataset2(t *testing.T) {
	type args struct {
		x [][]float64
		y []int
	}
	tests := []struct {
		name  string
		args  args
		want  [][]float64
		want1 []int
	}{
		{
			name: "SingleLine",
			args: args{
				x: [][]float64{
					{0.0, 0.1, 0.2, 0.3},
				},
				y: []int{0},
			},
			want: [][]float64{
				{0.0, 0.1, 0.2, 0.3},
			},
			want1: []int{0},
		},
		{
			name: "SingleLine",
			args: args{
				x: [][]float64{
					{0.0, 0.0, 0.0, 0.0},
					{1.0, 0.0, 0.0, 0.0},
					{2.0, 0.0, 0.0, 0.0},
				},
				y: []int{0, 1, 2},
			},
			want: [][]float64{
				{0.0, 0.0, 0.0, 0.0},
				{1.0, 0.0, 0.0, 0.0},
				{2.0, 0.0, 0.0, 0.0},
			},
			want1: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ShuffleDataset(tt.args.x, tt.args.y)

			sort.Ints(got1)
			sort.SliceStable(got, func(i, j int) bool {
				return got[i][0] < got[j][0]
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleDataset() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ShuffleDataset() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
