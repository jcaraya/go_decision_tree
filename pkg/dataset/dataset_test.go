package dataset

import (
	"reflect"
	"testing"
)

func TestDataset_IsEmpty(t *testing.T) {
	type fields struct {
		x [][]float64
		y []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "NonEmpty",
			fields: fields{
				x: [][]float64{
					[]float64{0.0, 0.1, 0.2},
				},
				y: []int{0},
			},
			want: false,
		},
		{
			name: "Empty",
			fields: fields{
				x: [][]float64{},
				y: []int{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dataset{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := d.IsEmpty(); got != tt.want {
				t.Errorf("Dataset.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataset_IsAllSameClass(t *testing.T) {
	type fields struct {
		x [][]float64
		y []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  bool
	}{
		{
			name: "Empty",
			fields: fields{
				x: [][]float64{},
				y: []int{},
			},
			want:  0,
			want1: false,
		},
		{
			name: "OneItem",
			fields: fields{
				x: [][]float64{
					[]float64{0.0, 0.1},
				},
				y: []int{1},
			},
			want:  1,
			want1: true,
		},
		{
			name: "MultipleClass",
			fields: fields{
				x: [][]float64{
					[]float64{0.0, 0.1},
					[]float64{0.0, 0.1},
					[]float64{0.0, 0.1},
					[]float64{0.0, 0.1},
				},
				y: []int{0, 1, 2, 2},
			},
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dataset{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			got, got1 := d.IsAllSameClass()
			if got != tt.want {
				t.Errorf("Dataset.IsAllSameClass() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Dataset.IsAllSameClass() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDataset_Partition(t *testing.T) {
	type fields struct {
		x [][]float64
		y []int
	}
	type args struct {
		decisionFunction func(interface{}) bool
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantLeft  *Dataset
		wantRight *Dataset
	}{
		{
			name: "BasicPartition",
			fields: fields{
				x: [][]float64{
					[]float64{1.0, 0.0},
					[]float64{1.0, 0.0},
					[]float64{1.0, 10.0},
					[]float64{1.0, 10.0},
				},
				y: []int{0, 1, 2, 2},
			},
			args: args{
				decisionFunction: func(xVal interface{}) bool {
					x := xVal.([]float64)
					return x[1] <= 5.0
				},
			},
			wantRight: &Dataset{
				x: [][]float64{
					[]float64{1.0, 0.0},
					[]float64{1.0, 0.0},
				},
				y: []int{0, 1},
			},
			wantLeft: &Dataset{
				x: [][]float64{
					[]float64{1.0, 10.0},
					[]float64{1.0, 10.0},
				},
				y: []int{2, 2},
			},
		},
		{
			name: "AllLeft",
			fields: fields{
				x: [][]float64{
					[]float64{1.0, 10.0},
					[]float64{1.0, 10.0},
					[]float64{1.0, 10.0},
					[]float64{1.0, 10.0},
				},
				y: []int{0, 1, 2, 2},
			},
			args: args{
				decisionFunction: func(xVal interface{}) bool {
					x := xVal.([]float64)
					return x[1] <= 5.0
				},
			},
			wantRight: &Dataset{
				x: nil,
				y: nil,
			},
			wantLeft: &Dataset{
				x: [][]float64{
					[]float64{1.0, 10.0},
					[]float64{1.0, 10.0},
					[]float64{1.0, 10.0},
					[]float64{1.0, 10.0},
				},
				y: []int{0, 1, 2, 2},
			},
		},

		{
			name: "AllRight",
			fields: fields{
				x: [][]float64{
					[]float64{1.0, 1.0},
					[]float64{1.0, 1.0},
					[]float64{1.0, 1.0},
					[]float64{1.0, 1.0},
				},
				y: []int{0, 1, 2, 2},
			},
			args: args{
				decisionFunction: func(xVal interface{}) bool {
					x := xVal.([]float64)
					return x[1] <= 5.0
				},
			},
			wantLeft: &Dataset{
				x: nil,
				y: nil,
			},
			wantRight: &Dataset{
				x: [][]float64{
					[]float64{1.0, 1.0},
					[]float64{1.0, 1.0},
					[]float64{1.0, 1.0},
					[]float64{1.0, 1.0},
				},
				y: []int{0, 1, 2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dataset{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			gotLeft, gotRight := d.Partition(tt.args.decisionFunction)
			if !reflect.DeepEqual(gotLeft, tt.wantLeft) {
				t.Errorf("Dataset.Partition() gotLeft = %v, want %v", gotLeft, tt.wantLeft)
			}
			if !reflect.DeepEqual(gotRight, tt.wantRight) {
				t.Errorf("Dataset.Partition() gotRight = %v, want %v", gotRight, tt.wantRight)
			}
		})
	}
}

func TestDataset_MostCommonClass(t *testing.T) {
	type fields struct {
		x [][]float64
		y []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "MultipleClassEnd",
			fields: fields{
				x: [][]float64{
					[]float64{0.0, 0.1},
					[]float64{0.0, 0.1},
					[]float64{0.0, 0.1},
					[]float64{0.0, 0.1},
				},
				y: []int{0, 1, 2, 2},
			},
			want: 2,
		},
		{
			name: "MultipleClassMid",
			fields: fields{
				x: [][]float64{
					[]float64{0.0, 0.1},
					[]float64{0.0, 0.1},
					[]float64{0.0, 0.1},
					[]float64{0.0, 0.1},
				},
				y: []int{0, 1, 1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dataset{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := d.MostCommonClass(); got != tt.want {
				t.Errorf("Dataset.MostCommonClass() = %v, want %v", got, tt.want)
			}
		})
	}
}
