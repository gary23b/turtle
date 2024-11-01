package turtlemodel

import (
	"image"
	"image/color"
)

// Type for the current turtle angle mode.
type AngleMode string

const (
	// Degrees mode. East is 0, North is 90, West is 180, South is 270
	DegreesMode AngleMode = "Degrees Angle Mode"

	// Radians mode. East is 0, North is pi/2, West is pi, South is 3/2*pi
	RadiansMode AngleMode = "Radians Angle Mode"

	// Compass mode. North is 0 degrees, and East is 90, South is 180, West is 270.
	CompassMode AngleMode = "Compass Angle Mode"

	// The max pen speed. No delay will be used.
	MaxSpeed float64 = 1.0e12
)

// The turtle interface.
// Every turtle created for the user of this package can do this exact list of actions.
// This mirrors the python turtle graphics command set.
// Each turtle is completely independent, and can be commanded in the same go routine or in different go routines.
type Turtle interface {
	Forward(distance float64)
	F(distance float64) // Forward alias
	Backward(distance float64)
	B(distance float64) // Backward alias
	PanRightward(distance float64)
	PanR(distance float64) // PanRightward alias
	PanLeftward(distance float64)
	PanL(distance float64) // PanLeftward alias

	GoTo(x, y float64)      // Cartesian (x,y). Center in the middle of the window
	GetPos() (x, y float64) // Cartesian (x,y). Center in the middle of the window

	// instantly move to a position.
	// Cartesian (x,y). Center in the middle of the window
	Teleport(x, y float64)

	Left(angle float64)
	L(angle float64) // Turn Left alias
	Right(angle float64)
	R(angle float64) // Turn Right alias
	Angle(angle float64)
	GetAngle() float64
	PointToward(x, y float64)

	DegreesMode() // Default is degrees mode.
	RadiansMode()
	CompassMode() // Make it so North is 0 degrees, East is 90...
	GetAngleMode() AngleMode

	Speed(PixelsPerSecond float64)
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
	Size(size float64)
	GetSize() float64
	Dot(size float64)
	Fill(c color.Color)

	// Draw a circle with given radius. The center is radius units left of the turtle; angleAmountToDraw determines
	// which part of the circle is drawn. If angleAmountToDraw is not a full circle, one endpoint of the arc is
	// the current pen position. Draw the arc in counterclockwise direction if radius is positive,
	// otherwise in clockwise direction. Finally the direction of the turtle is changed by the amount of angleAmountToDraw.
	//
	// As the circle is approximated by an inscribed regular polygon, steps determines the number of steps to use.
	// May be used to draw regular polygons.
	Circle(radius, angleAmountToDraw float64, steps int)

	ShowTurtle()
	HideTurtle()    // Default
	ShapeAsTurtle() // Default
	ShapeAsArrow()
	ShapeAsImage(in image.Image)
	ShapeScale(scale float64) // Default = 0.35

	Clone() Turtle // Create a clone of the turtle
}
