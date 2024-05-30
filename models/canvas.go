package models

import (
	"image"
	"image/color"
)

type Canvas interface {
	CreateNewSprite() Sprite
	SetCartesianPixel(x, y int, c color.Color) // Cartesian (x,y). Center in the middle of the window
	SetPixel(x, y int, c color.Color)          // Computer graphics (x,y). So x=0, y=0 is the top-left of the window, positive down-right.
	Fill(x, y int, c color.Color)              // Cartesian (x,y). Center in the middle of the window
	ClearScreen(c color.Color)
	GetScreenshot() image.Image

	GetWidth() int
	GetHeight() int

	PressedUserInput() *UserInput
	SubscribeToJustPressedUserInput() chan *UserInput
	UnSubscribeToJustPressedUserInput(in chan *UserInput)

	Exit()
}
