package turtlemodel

import (
	"image"
	"image/color"
)

// Not sure if it is helpful or not to make it so a user doesn't have to cast things to float64 to use this package.
// Looking at the current examples, you only save a couple casts here and there.
// You save one mainly when you want to use the loop index as an x or y position.

// The same as the turtle interface, except all number inputs are in the form of "any". This makes it so you don't have cast all your numbers to float64.
// Every turtle created for the user of this package can do this exact list of actions.
// This mirrors the python turtle graphics command set.
// Each turtle is completely independent, and can be commanded in the same go routine or in different go routines.
type TurtlePermissive interface {
	Forward(distance any)
	F(distance any) // Forward alias
	Backward(distance any)
	B(distance any) // Backward alias
	PanRightward(distance any)
	PanR(distance any) // PanRightward alias
	PanLeftward(distance any)
	PanL(distance any) // PanLeftward alias

	GoTo(x, y any)          // Cartesian (x,y). Center in the middle of the window
	GetPos() (x, y float64) // Cartesian (x,y). Center in the middle of the window

	Left(angle any)
	L(angle any) // Turn Left alias
	Right(angle any)
	R(angle any) // Turn Right alias
	Angle(angle any)
	GetAngle() float64
	PointToward(x, y any)

	DegreesMode() // Default is degrees mode.
	RadiansMode()
	CompassMode() // Make it so North is 0 degrees, East is 90...
	GetAngleMode() AngleMode

	Speed(PixelsPerSecond any)
	GetSpeed() float64

	PenUp()
	PU()  // Pen Up alias
	Off() // Pen Up alias
	PenDown()
	PD() // Pen Down alias
	On() // Pen Down alias
	IsPenDown() bool
	Color(c color.Color)
	GetColor() color.Color
	Size(size any)
	GetSize() float64
	Dot(size any)
	Fill(c color.Color)

	// Draw a circle with given radius. The center is radius units left of the turtle; angleAmountToDraw determines
	// which part of the circle is drawn. If angleAmountToDraw is not a full circle, one endpoint of the arc is
	// the current pen position. Draw the arc in counterclockwise direction if radius is positive,
	// otherwise in clockwise direction. Finally the direction of the turtle is changed by the amount of angleAmountToDraw.
	//
	// As the circle is approximated by an inscribed regular polygon, steps determines the number of steps to use.
	// May be used to draw regular polygons.
	Circle(radius, angleAmountToDraw any, steps int)

	ShowTurtle()
	HideTurtle()    // Default
	ShapeAsTurtle() // Default
	ShapeAsArrow()
	ShapeAsImage(in image.Image)
	ShapeScale(scale any) // Default = 0.35
}
