package models

import "image/color"

type Canvas interface {
	SetCartesianPixel(x, y int, c color.RGBA)
	SetPixel(x, y int, c color.RGBA)
	Fill(x, y int, c color.RGBA) // Cartesian x,y
	ClearScreen(c color.RGBA)

	GetWidth() int
	GetHeight() int
	GetUserInput() UserInput
}
