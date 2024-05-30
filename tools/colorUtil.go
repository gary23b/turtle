package tools

import "image/color"

// https://www.w3schools.com/colors/colors_names.asp
var (
	Black color.RGBA = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	White color.RGBA = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	Red   color.RGBA = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	Lime  color.RGBA = color.RGBA{0x00, 0xFF, 0x00, 0xFF}
	Blue  color.RGBA = color.RGBA{0x00, 0x00, 0xFF, 0xFF}

	Yellow  color.RGBA = color.RGBA{0xFF, 0xFF, 0x00, 0xFF}
	Aqua    color.RGBA = color.RGBA{0x00, 0xFF, 0xFF, 0xFF}
	Magenta color.RGBA = color.RGBA{0xFF, 0x00, 0xFF, 0xFF}

	Orange color.RGBA = color.RGBA{0xFF, 0xA5, 0x00, 0xFF}
	Green  color.RGBA = color.RGBA{0x00, 0x80, 0x00, 0xFF}
	Purple color.RGBA = color.RGBA{0x80, 0x00, 0x80, 0xFF}
	Indigo color.RGBA = color.RGBA{0x4B, 0x00, 0x82, 0xFF}
	Violet color.RGBA = color.RGBA{0xEE, 0x82, 0xEE, 0xFF}
)

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
