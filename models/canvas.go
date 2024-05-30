package models

import "image/color"

type Canvas interface {
	CreateNewSprite() Sprite
	SetCartesianPixel(x, y int, c color.Color) // Cartesian (x,y). Center in the middle of the window
	SetPixel(x, y int, c color.Color)          // Computer graphics (x,y). So y=0 is the top of the window, positive down.
	Fill(x, y int, c color.Color)              // Cartesian (x,y). Center in the middle of the window
	ClearScreen(c color.Color)

	GetWidth() int
	GetHeight() int

	PressedUserInput() *UserInput
	SubscribeToJustPressedUserInput() chan *UserInput
	UnSubscribeToJustPressedUserInput(in chan *UserInput)

	Exit()
}
