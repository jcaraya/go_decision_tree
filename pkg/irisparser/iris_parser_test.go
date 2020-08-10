package irisparser

import (
	"reflect"
	"testing"
)

// Validates all the test cases related to parsing CSV lines.
func TestParseCsvLines(t *testing.T) {
	type args struct {
		csvLines [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]float64
		want1   []int
		want2   map[string]int
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				csvLines: [][]string{
					[]string{"0.0", "0.1", "0.2", "0.3", "class1"},
				},
			},
			want: [][]float64{
				[]float64{0.0, 0.1, 0.2, 0.3},
			},
			want1: []int{0},
			want2: map[string]int{
				"class1": 0,
			},
			wantErr: false,
		},
		{
			name: "MultipleLinesSingleClass",
			args: args{
				csvLines: [][]string{
					[]string{"0.0", "0.1", "0.2", "0.3", "class1"},
					[]string{"0.0", "0.2", "0.4", "0.6", "class1"},
				},
			},
			want: [][]float64{
				[]float64{0.0, 0.1, 0.2, 0.3},
				[]float64{0.0, 0.2, 0.4, 0.6},
			},
			want1: []int{0, 0},
			want2: map[string]int{
				"class1": 0,
			},
			wantErr: false,
		},
		{
			name: "MultipleLineSingleClass",
			args: args{
				csvLines: [][]string{
					[]string{"0.0", "0.1", "0.2", "0.3", "class1"},
					[]string{"0.0", "0.2", "0.4", "0.6", "class2"},
					[]string{"0.0", "0.3", "0.6", "0.9", "class1"},
				},
			},
			want: [][]float64{
				[]float64{0.0, 0.1, 0.2, 0.3},
				[]float64{0.0, 0.2, 0.4, 0.6},
				[]float64{0.0, 0.3, 0.6, 0.9},
			},
			want1: []int{0, 1, 0},
			want2: map[string]int{
				"class1": 0,
				"class2": 1,
			},
			wantErr: false,
		},
		{
			name: "ParseFailure",
			args: args{
				csvLines: [][]string{
					[]string{"0.0", "0.1", "0.2", "0.3A", "class1"},
				},
			},
			want:    nil,
			want1:   nil,
			want2:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := ParseCsvLines(tt.args.csvLines)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLines() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ParseLines() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("ParseLines() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
