package turtleutil

import "image/color"

var (
	Water color.RGBA = color.RGBA{0x23, 0x89, 0xDA, 0xFF} //2389DA
	Black color.RGBA = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	White color.RGBA = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	Red   color.RGBA = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	Green color.RGBA = color.RGBA{0x00, 0xFF, 0x00, 0xFF}
	Blue  color.RGBA = color.RGBA{0x00, 0x00, 0xFF, 0xFF}
)

func Lerp(a, b, ratio float64) float64 {
	y := a + (b-a)*ratio
	return y
}

func LerpUint8(a, b uint8, ratio float64) uint8 {
	aF := float64(a)
	bF := float64(b)
	y := aF + (bF-aF)*ratio
	return uint8(y)
}

func LerpColor(a, b color.RGBA, ratio float64) color.RGBA {
	ret := color.RGBA{
		R: LerpUint8(a.R, b.R, ratio),
		G: LerpUint8(a.G, b.G, ratio),
		B: LerpUint8(a.B, b.B, ratio),
		A: LerpUint8(a.A, b.A, ratio),
	}

	return ret
}
