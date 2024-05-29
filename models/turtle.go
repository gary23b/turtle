package models

import "image/color"

type Turtle interface {
	Forward(distance float64)
	F(distance float64) // Forward alias
	Backward(distance float64)
	B(distance float64) // Backward alias
	PanRightward(distance float64)
	PanR(distance float64) // PanRightward alias
	PanLeftward(distance float64)
	PanL(distance float64) // PanLeftward alias

	GoTo(x, y float64)
	GetPos() (x, y float64)

	Left(angle float64)
	L(angle float64) // Turn Left alias
	Right(angle float64)
	R(angle float64) // Turn Right alias
	SetAngle(angle float64)
	PointToward(x, y float64)
	GetAngle() float64

	SetDegreesMode() // Default is degree mode.
	SetRadianMode()
	EnableCompassAngleMode(in bool) // Make it so North is 0 degrees, East is 90...

	SetSpeed(PixelsPerSecond float64)

	PenUp()
	PU()  // Pen Up alias
	Off() // Pen Up alias
	PenDown()
	PD() // Pen Down alias
	On() // Pen Down alias
	PenColor(c color.RGBA)
	PenSize(size float64)
}
