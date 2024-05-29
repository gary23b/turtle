package turtle

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_turtle_angleToRad(t *testing.T) {
	type args struct {
		angle float64
	}
	tests := []struct {
		name string
		s    *turtle
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "North",
			s:    &turtle{degreesEn: true, compassEn: true},
			args: args{angle: 0},
			want: math.Pi / 2,
		},
		{
			name: "East",
			s:    &turtle{degreesEn: true, compassEn: true},
			args: args{angle: 90},
			want: 0,
		},
		{
			name: "South",
			s:    &turtle{degreesEn: true, compassEn: true},
			args: args{angle: 180},
			want: -math.Pi / 2,
		},
		{
			name: "West",
			s:    &turtle{degreesEn: true, compassEn: true},
			args: args{angle: 270},
			want: -math.Pi,
		},
		{
			name: "North-West",
			s:    &turtle{degreesEn: true, compassEn: true},
			args: args{angle: -45},
			want: 3.0 / 4.0 * math.Pi,
		},
		//
		{
			name: "0 degrees",
			s:    &turtle{degreesEn: true},
			args: args{angle: 0},
			want: 0,
		},
		{
			name: "45 degrees",
			s:    &turtle{degreesEn: true},
			args: args{angle: 45},
			want: math.Pi / 4,
		},
		{
			name: "-100 degrees",
			s:    &turtle{degreesEn: true},
			args: args{angle: -100},
			want: -100.0 * math.Pi / 180.0,
		},
		{
			name: "360+45 degrees",
			s:    &turtle{degreesEn: true},
			args: args{angle: 360 + 45},
			want: (360 + 45) * math.Pi / 180.0,
		},
		//
		{
			name: "360+45 degrees",
			s:    &turtle{},
			args: args{angle: 1234.12345},
			want: 1234.12345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// require.InDelta
			got := tt.s.angleToRad(tt.args.angle)
			require.Equal(t, tt.want, got)
			got2 := tt.s.radToAngle(got)
			require.Equal(t, tt.args.angle, got2)
		})
	}
}
