package turtletools

import (
	"image/color"
	"reflect"
	"testing"
)

func TestLerp(t *testing.T) {
	type args struct {
		a     float64
		b     float64
		ratio float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{name: "ratio 0", args: args{a: 0, b: 100, ratio: 0}, want: 0},
		{name: "ratio .5", args: args{a: 0, b: 100, ratio: .5}, want: 50},
		{name: "ratio 1", args: args{a: 0, b: 100, ratio: 1}, want: 100},
		// Test negative direction
		{name: "ratio 0", args: args{a: 100, b: -100, ratio: 0}, want: 100},
		{name: "ratio .5", args: args{a: 100, b: -100, ratio: .5}, want: 0},
		{name: "ratio 1", args: args{a: 100, b: -100, ratio: 1}, want: -100},
		// Now test out of bounds behavior
		{name: "ratio 3", args: args{a: 0, b: 100, ratio: 3}, want: 100},
		{name: "ratio -1", args: args{a: 0, b: 100, ratio: -1}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lerp(tt.args.a, tt.args.b, tt.args.ratio); got != tt.want {
				t.Errorf("Lerp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLerpUint8(t *testing.T) {
	type args struct {
		a     uint8
		b     uint8
		ratio float64
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		// TODO: Add test cases.
		{name: "ratio 0", args: args{a: 0, b: 255, ratio: 0}, want: 0},
		{name: "ratio 0", args: args{a: 0, b: 255, ratio: .5}, want: 127},
		{name: "ratio 0", args: args{a: 0, b: 255, ratio: 1}, want: 255},
		{name: "ratio 0", args: args{a: 255, b: 128, ratio: 0}, want: 255},
		{name: "ratio 0", args: args{a: 255, b: 128, ratio: .5}, want: 191},
		{name: "ratio 0", args: args{a: 255, b: 128, ratio: 1}, want: 128},
		// Now test out of bounds behavior
		{name: "ratio 0", args: args{a: 0, b: 255, ratio: 2.5}, want: 255},
		{name: "ratio 0", args: args{a: 0, b: 255, ratio: -1}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lerp(tt.args.a, tt.args.b, tt.args.ratio); got != tt.want {
				t.Errorf("LerpUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLerpColor(t *testing.T) {
	var (
		black color.RGBA = color.RGBA{0x00, 0x00, 0x00, 0xFF} // #000000
		white color.RGBA = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF} // #FFFFFF
	)

	type args struct {
		a     color.RGBA
		b     color.RGBA
		ratio float64
	}
	tests := []struct {
		name string
		args args
		want color.RGBA
	}{
		// TODO: Add test cases.
		{name: "ratio 0", args: args{a: black, b: white, ratio: 0}, want: black},
		{name: "ratio 1", args: args{a: black, b: white, ratio: 1}, want: white},
		{
			name: "ratio 1",
			args: args{
				a:     color.RGBA{0x00, 0x00, 0x00, 0xFF},
				b:     color.RGBA{0x10, 0x10, 0x10, 0x00},
				ratio: .5,
			},
			want: color.RGBA{0x08, 0x08, 0x08, 0x7F},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LerpColor(tt.args.a, tt.args.b, tt.args.ratio); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LerpColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
