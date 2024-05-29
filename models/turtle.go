package models

import "image/color"

type Turtle interface {
	Forward(distance float64)
	Backward(distance float64)
	PanRightward(distance float64)
	PanLeftward(distance float64)

	GoTo(x, y float64)
	GetPos() (x, y float64)

	Left(angle float64)
	Right(angle float64)
	SetAngle(angle float64)
	PointToward(x, y float64)
	GetAngle() float64

	SetDegreesMode()
	SetRadianMode()
	EnableCompassAngleMode(in bool)

	SetSpeed(PixelsPerSecond float64)

	PenUp()
	PenDown()
	PenColor(c color.RGBA)
	PenSize(size float64)
}
