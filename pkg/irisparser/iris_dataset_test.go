package irisparser

import (
	"reflect"
	"testing"
)

func TestNewIrisDataset(t *testing.T) {
	type args struct {
		lines [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    *IrisDataset
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				lines: [][]string{
					[]string{"0.0", "0.1", "0.2", "0.3", "class1"},
				},
			},
			want: &IrisDataset{
				x: [][]float64{
					[]float64{0.0, 0.1, 0.2, 0.3},
				},
				y: []int{0},
				categorySet: map[string]int{
					"class1": 0,
				},
				categoryCounter: 1,
			},
			wantErr: false,
		},
		{
			name: "MultipleLinesSingleClass",
			args: args{
				lines: [][]string{
					[]string{"0.0", "0.1", "0.2", "0.3", "class1"},
					[]string{"0.0", "0.2", "0.4", "0.6", "class1"},
				},
			},
			want: &IrisDataset{
				x: [][]float64{
					[]float64{0.0, 0.1, 0.2, 0.3},
					[]float64{0.0, 0.2, 0.4, 0.6},
				},
				y: []int{0, 0},
				categorySet: map[string]int{
					"class1": 0,
				},
				categoryCounter: 1,
			},
			wantErr: false,
		},
		{
			name: "MultipleLineSingleClass",
			args: args{
				lines: [][]string{
					[]string{"0.0", "0.1", "0.2", "0.3", "class1"},
					[]string{"0.0", "0.2", "0.4", "0.6", "class2"},
					[]string{"0.0", "0.3", "0.6", "0.9", "class1"},
				},
			},
			want: &IrisDataset{
				x: [][]float64{
					[]float64{0.0, 0.1, 0.2, 0.3},
					[]float64{0.0, 0.2, 0.4, 0.6},
					[]float64{0.0, 0.3, 0.6, 0.9},
				},
				y: []int{0, 1, 0},
				categorySet: map[string]int{
					"class1": 0,
					"class2": 1,
				},
				categoryCounter: 2,
			},
			wantErr: false,
		},
		{
			name: "ParseFailure",
			args: args{
				lines: [][]string{
					[]string{"0.0", "0.1", "0.2", "0.3A", "class1"},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIrisDataset(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIrisDataset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIrisDataset() = %v, want %v", got, tt.want)
			}
		})
	}
}
