package pen

import (
	"log"

	"github.com/gary23b/turtle/turtlemodel"
)

// Not sure if it is helpful or not to make it so a user doesn't have to cast things to float64 to use this package.
// Looking at the current examples, you only save a couple casts here and there.
// You save one mainly when you want to use the loop index as an x or y position.

// Use the normal pen, and overwrite some of the methods.
type penPermissive struct {
	turtlemodel.Turtle
}

var _ turtlemodel.TurtlePermissive = &penPermissive{} // Force the linter to tell us if the interface is implemented

// Create a new pen that implements the turtle interface.
func NewPenPermissive(can turtlemodel.Canvas) *penPermissive {
	var p penPermissive
	p.Turtle = NewPen(can)

	return &p
}

func toF64(a any) (float64, bool) {
	switch num := a.(type) {
	case uint8:
		return float64(num), true
	case int8:
		return float64(num), true
	case uint16:
		return float64(num), true
	case int16:
		return float64(num), true
	case uint32:
		return float64(num), true
	case int32:
		return float64(num), true
	case uint64:
		return float64(num), true
	case int64:
		return float64(num), true
	case int:
		return float64(num), true
	case float32:
		return float64(num), true
	case float64:
		return num, true
	}

	return 0, false
}

// Move forward in the current direction. Distance in units of pixels.
func (s *penPermissive) Forward(distance any) {
	d, ok := toF64(distance)
	if !ok {
		log.Println("Parameter is not a real number", distance)
	}
	s.Turtle.Forward(d)
}

// Alias of Forward
func (s *penPermissive) F(distance any) {
	s.Forward(distance)
}

// Move backward in the current direction. Distance in units of pixels.
func (s *penPermissive) Backward(distance any) {
	d, ok := toF64(distance)
	if !ok {
		log.Println("Parameter is not a real number", distance)
	}
	s.Turtle.Backward(d)
}

// Alias of Backward
func (s *penPermissive) B(distance any) {
	s.Backward(distance)
}

// While maintaining the current direction, move straight left. Crab walk left.
func (s *penPermissive) PanLeftward(distance any) {
	d, ok := toF64(distance)
	if !ok {
		log.Println("Parameter is not a real number", distance)
	}
	s.Turtle.PanLeftward(d)
}

// Alias of PanLeftward
func (s *penPermissive) PanL(distance any) {
	s.PanLeftward(distance)
}

// While maintaining the current direction, move straight right. Crab walk right.
func (s *penPermissive) PanRightward(distance any) {
	d, ok := toF64(distance)
	if !ok {
		log.Println("Parameter is not a real number", distance)
	}
	s.Turtle.PanRightward(d)
}

// Alias of PanRightward
func (s *penPermissive) PanR(distance any) {
	s.PanRightward(distance)
}

// From the current position, go to an absolute x, y position
func (s *penPermissive) GoTo(x, y any) {
	xF64, ok := toF64(x)
	yF64, ok2 := toF64(y)
	if !ok || !ok2 {
		log.Println("Parameter is not a real number", x, y)
	}
	s.Turtle.GoTo(xF64, yF64)
}

// Turn left by a given angle.
// The angle amount will be interpreted as radians or degrees based on the current angle mode.
func (s *penPermissive) Left(angle any) {
	a, ok := toF64(angle)
	if !ok {
		log.Println("Parameter is not a real number", angle)
	}
	s.Turtle.Left(a)
}

// Alias of Left
func (s *penPermissive) L(angle any) {
	s.Left(angle)
}

// Turn right by a given angle.
// The angle amount will be interpreted as radians or degrees based on the current angle mode.
func (s *penPermissive) Right(angle any) {
	a, ok := toF64(angle)
	if !ok {
		log.Println("Parameter is not a real number", angle)
	}
	s.Turtle.Right(a)
}

// Alias of Right
func (s *penPermissive) R(angle any) {
	s.Right(angle)
}

// Set the pen angle.
// The angle amount will be interpreted as radians or degrees based on the current angle mode.
func (s *penPermissive) Angle(angle any) {
	a, ok := toF64(angle)
	if !ok {
		log.Println("Parameter is not a real number", angle)
	}
	s.Turtle.Angle(a)
}

// Make the pen point toward an x,y location.
func (s *penPermissive) PointToward(x, y any) {
	xF64, ok := toF64(x)
	yF64, ok2 := toF64(y)
	if !ok || !ok2 {
		log.Println("Parameter is not a real number", x, y)
	}
	s.Turtle.PointToward(xF64, yF64)
}

// Set the pen's speed in pixels per second.
func (s *penPermissive) Speed(pixelsPerSecond any) {
	speed, ok := toF64(pixelsPerSecond)
	if !ok {
		log.Println("Parameter is not a real number", pixelsPerSecond)
	}
	s.Turtle.Speed(speed)
}

// Set the pen's draw size.
// As the pen moves, it will only draw the exact pixel it is on top of when the size is 0.
// For a size greater than 0, all pixels that are within the radius of size/2 will be colored.
func (s *penPermissive) Size(size any) {
	size2, ok := toF64(size)
	if !ok {
		log.Println("Parameter is not a real number", size)
	}
	s.Turtle.Size(size2)
}

// Draw the single pixel that the pen is currently on top of.
func (s *penPermissive) Dot(size any) {
	size2, ok := toF64(size)
	if !ok {
		log.Println("Parameter is not a real number", size)
	}
	s.Turtle.Dot(size2)
}

// Draw a circle with given radius. The center is radius units left of the turtle; angleAmountToDraw determines
// which part of the circle is drawn. If angleAmountToDraw is not a full circle, one endpoint of the arc is
// the current pen position. Draw the arc in counterclockwise direction if radius is positive,
// otherwise in clockwise direction. Finally the direction of the turtle is changed by the amount of angleAmountToDraw.
//
// As the circle is approximated by an inscribed regular polygon, steps determines the number of steps to use.
// May be used to draw regular polygons.
func (s *penPermissive) Circle(radius, angleAmountToDraw any, steps int) {
	rF64, ok := toF64(radius)
	aF64, ok2 := toF64(angleAmountToDraw)
	if !ok || !ok2 {
		log.Println("Parameter is not a real number", radius, angleAmountToDraw)
	}
	s.Turtle.Circle(rF64, aF64, steps)
}

// Set the turtle sprite bitmap scale
func (s *penPermissive) ShapeScale(scale any) {
	scale2, ok := toF64(scale)
	if !ok {
		log.Println("Parameter is not a real number", scale)
	}
	s.Turtle.ShapeScale(scale2)
}
