package ward

import (
	"reflect"
	"testing"
)

func TestWard(t *testing.T) {
	type args struct {
		X [][]float64
	}
	tests := []struct {
		name string
		args args
		want Tree
	}{
		{
			name: "test for Ward",
			args: args{
				[][]float64{
					{0, 0},
					{1, 0},
					{5, 0},
				},
			},
			want: Tree{
				Group{
					left:    -1,
					right:   -1,
					N:       1,
					vec:     []float64{0, 0},
					dist:    0,
					visited: true,
				},
				Group{
					left:    -1,
					right:   -1,
					N:       1,
					vec:     []float64{1, 0},
					dist:    0,
					visited: true,
				},
				Group{
					left:    -1,
					right:   -1,
					N:       1,
					vec:     []float64{5, 0},
					dist:    0,
					visited: true,
				},
				Group{
					left:    0,
					right:   1,
					N:       2,
					vec:     []float64{0.5, 0},
					dist:    0.5,
					visited: true,
				},
				Group{
					left:    2,
					right:   3,
					N:       3,
					vec:     []float64{2, 0},
					dist:    13.5,
					visited: false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ward(tt.args.X); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkage(t *testing.T) {
	type args struct {
		T Tree
	}
	tests := []struct {
		name string
		args args
		want Tree
	}{
		{
			name: "test for linkage of two Groups",
			args: args{
				Tree{
					Group{
						left:    -1,
						right:   -1,
						N:       1,
						vec:     []float64{0, 0},
						dist:    0,
						visited: false,
					},
					Group{
						left:    -1,
						right:   -1,
						N:       1,
						vec:     []float64{2, 0},
						dist:    0,
						visited: false,
					},
				},
			},
			want: Tree{
				Group{
					left:    -1,
					right:   -1,
					N:       1,
					vec:     []float64{0, 0},
					dist:    0,
					visited: true,
				},
				Group{
					left:    -1,
					right:   -1,
					N:       1,
					vec:     []float64{2, 0},
					dist:    0,
					visited: true,
				},
				Group{
					left:    0,
					right:   1,
					N:       2,
					vec:     []float64{1, 0},
					dist:    2,
					visited: false,
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := linkage(tt.args.T); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("linkage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_centerOfGravity(t *testing.T) {
	type args struct {
		g1 Group
		g2 Group
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "test for centerOfGravity (0, 0) and (2, 0)",
			args: args{
				Group{
					left:    -1,
					right:   -1,
					N:       1,
					vec:     []float64{0, 0},
					dist:    0,
					visited: false,
				},
				Group{
					left:    -1,
					right:   -1,
					N:       1,
					vec:     []float64{2, 0},
					dist:    0,
					visited: false,
				},
			},
			want: []float64{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := centerOfGravity(tt.args.g1, tt.args.g2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("centerOfGravity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initialize(t *testing.T) {
	type args struct {
		X [][]float64
	}
	tests := []struct {
		name string
		args args
		want Tree
	}{
		{
			name: "test for initialize",
			args: args{
				[][]float64{
					{0, 0},
				},
			},
			want: Tree{Group{
				left:    -1,
				right:   -1,
				N:       1,
				vec:     []float64{0, 0},
				dist:    0,
				visited: false,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initialize(tt.args.X); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dist(t *testing.T) {
	type args struct {
		g1 Group
		g2 Group
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test for dist from (0, 0) to (3, 4)",
			args: args{
				Group{
					left:    -1,
					right:   -1,
					N:       1,
					vec:     []float64{0, 0},
					dist:    0,
					visited: false,
				},
				Group{
					left:    -1,
					right:   -1,
					N:       1,
					vec:     []float64{3, 4},
					dist:    0,
					visited: false,
				},
			},
			want: 12.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dist(tt.args.g1, tt.args.g2); got != tt.want {
				t.Errorf("dist() = %v, want %v", got, tt.want)
			}
		})
	}
}
