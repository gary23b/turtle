package turtle

import "image/color"

// The ratio is capped between 0 and 1
func Lerp(a, b, ratio float64) float64 {
	if ratio > 1 {
		ratio = 1
	} else if ratio < 0 {
		ratio = 0
	}
	y := a + (b-a)*ratio
	return y
}

// The ratio is capped between 0 and 1
// Currently the function floors the number instead of rounding to nearest.
func LerpUint8(a, b uint8, ratio float64) uint8 {
	if ratio > 1 {
		ratio = 1
	} else if ratio < 0 {
		ratio = 0
	}

	aF := float64(a)
	bF := float64(b)
	y := aF + (bF-aF)*ratio
	return uint8(y)
}

// Creates a color between the given a and b. 0 means a is given, 1 means b is given, .5 is a color half way between.
// The ratio is capped between 0 and 1
// Currently the function floors the number instead of rounding to nearest.
func LerpColor(a, b color.RGBA, ratio float64) color.RGBA {
	ret := color.RGBA{
		R: LerpUint8(a.R, b.R, ratio),
		G: LerpUint8(a.G, b.G, ratio),
		B: LerpUint8(a.B, b.B, ratio),
		A: LerpUint8(a.A, b.A, ratio),
	}

	return ret
}
