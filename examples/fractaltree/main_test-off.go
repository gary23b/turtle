package main

import "testing"

func Test_dualRate(t *testing.T) {
	type args struct {
		x1 float64
		x2 float64
		y1 float64
		y2 float64
		in float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{args: args{x1: 0, x2: 1, y1: 0, y2: 1, in: 0}, want: 0},
		{args: args{x1: 0, x2: 1, y1: 0, y2: 1, in: 1}, want: 1},
		{args: args{x1: 0, x2: 1, y1: 0, y2: 1, in: -1}, want: -1},
		{args: args{x1: 10, x2: 9, y1: 10, y2: 0, in: 10}, want: 10},
		{args: args{x1: 10, x2: 9, y1: 10, y2: 0, in: 9}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dualRate(tt.args.x1, tt.args.x2, tt.args.y1, tt.args.y2, tt.args.in); got != tt.want {
				t.Errorf("dualRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
