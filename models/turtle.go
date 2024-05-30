package models

import (
	"image"
	"image/color"
)

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
	PaintDot(size float64)
	Fill(c color.RGBA)

	// Draw a circle with given radius. The center is radius units left of the turtle; angleAmountToDraw determines
	// which part of the circle is drawn. If angleAmountToDraw is not a full circle, one
	// endpoint of the arc is the current pen position. Draw the arc in counterclockwise direction if radius is positive,
	// otherwise in clockwise direction. Finally the direction of the turtle is changed by the amount of angleAmountToDraw.
	//
	// As the circle is approximated by an inscribed regular polygon, steps determines the number of steps to use.
	// May be used to draw regular polygons.
	Circle(radius, angleAmountToDraw float64, steps int)

	SetVisible(isVisible bool)
	SetShapeAsTurtle()
	SetShapeAsArrow()
	SetShapeAsImage(in image.Image)
	SetShapeScale(scale float64)
}
