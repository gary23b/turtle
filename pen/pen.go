package pen

import (
	"image"
	"image/color"
	"math"
	"time"

	"github.com/gary23b/turtle/turtlemodel"
)

type pen struct {
	can       turtlemodel.Canvas
	sprite    turtlemodel.Sprite
	visible   bool
	x         float64 // in pixels
	y         float64 // in pixels
	angle     float64 // in radians
	degreesEn bool
	compassEn bool
	speed     float64 // in pixels/second
	penDown   bool
	penColor  color.Color
	penSize   float64
}

var _ turtlemodel.Turtle = &pen{} // Force the linter to tell us if the interface is implemented

// Create a new pen that implements the turtle interface.
func NewPen(can turtlemodel.Canvas) *pen {
	black := color.RGBA{A: 0xFF}
	ret := &pen{
		can:       can,
		sprite:    can.CreateNewSprite(),
		penColor:  black,
		speed:     75,
		penSize:   0,
		degreesEn: true,
	}

	return ret
}

// Move forward in the current direction. Distance in units of pixels.
func (s *pen) Forward(distance float64) {
	s.goAngle(s.angle, distance)
}

// Alias of Forward
func (s *pen) F(distance float64) {
	s.Forward(distance)
}

// Move backward in the current direction. Distance in units of pixels.
func (s *pen) Backward(distance float64) {
	s.goAngle(s.angle, -distance)
}

// Alias of Backward
func (s *pen) B(distance float64) {
	s.Backward(distance)
}

// While maintaining the current direction, move straight left. Crab walk left.
func (s *pen) PanLeftward(distance float64) {
	s.goAngle(s.angle+math.Pi/2.0, distance)
}

// Alias of PanLeftward
func (s *pen) PanL(distance float64) {
	s.PanLeftward(distance)
}

// While maintaining the current direction, move straight right. Crab walk right.
func (s *pen) PanRightward(distance float64) {
	s.goAngle(s.angle-math.Pi/2.0, distance)
}

// Alias of PanRightward
func (s *pen) PanR(distance float64) {
	s.PanRightward(distance)
}

// From the current position, go a given distance at the given angle.
func (s *pen) goAngle(angle, distance float64) {
	sin, cos := math.Sincos(angle)
	x2 := s.x + distance*cos
	y2 := s.y + distance*sin
	s.GoTo(x2, y2)
}

// From the current position, go to an absolute x, y position
func (s *pen) GoTo(x, y float64) {
	if s.penDown {
		s.drawLine(s.x, s.y, x, y, s.penColor)
	}
	s.x = x
	s.y = y

	s.sprite.Set(s.visible, s.x, s.y, s.angle)
}

// Get the current pen position
func (s *pen) GetPos() (x, y float64) {
	return s.x, s.y
}

// Turn left by a given angle.
// The angle amount will be interpreted as radians or degrees based on the current angle mode.
func (s *pen) Left(angle float64) {
	if s.degreesEn {
		angle *= (math.Pi / 180.0)
	}
	s.angle += angle

	// Normalize the angle
	if s.angle > 2*math.Pi {
		s.angle -= 2 * math.Pi
	} else if s.angle < -2*math.Pi {
		s.angle += 2 * math.Pi
	}
	s.sprite.SetRotation(s.angle)
}

// Alias of Left
func (s *pen) L(angle float64) {
	s.Left(angle)
}

// Turn right by a given angle.
// The angle amount will be interpreted as radians or degrees based on the current angle mode.
func (s *pen) Right(angle float64) {
	s.Left(-angle)
}

// Alias of Right
func (s *pen) R(angle float64) {
	s.Right(angle)
}

// Set the pen angle.
// The angle amount will be interpreted as radians or degrees based on the current angle mode.
func (s *pen) Angle(angle float64) {
	angle = s.absoluteAngleToRad(angle)
	s.angle = angle
	s.sprite.SetRotation(s.angle)
}

// Get the pen angle.
// The angle amount will be radians or degrees based on the current angle mode.
func (s *pen) GetAngle() float64 {
	return s.radToAbsoluteAngle(s.angle)
}

// Make the pen point toward an x,y location.
func (s *pen) PointToward(x, y float64) {
	deltaX := x - s.x
	deltaY := y - s.y
	s.angle = math.Atan2(deltaY, deltaX)
	s.sprite.SetRotation(s.angle)
}

// Set the turtle to degrees mode. East is 0, North is 90, West is 180, South is 270
// This is the default mode
func (s *pen) DegreesMode() {
	s.degreesEn = true
	s.compassEn = false
}

// Set the turtle to Radians mode. East is 0, North is pi/2, West is pi, South is 3/2*pi
func (s *pen) RadiansMode() {
	s.degreesEn = false
	s.compassEn = false
}

// Set the turtle to Compass mode. North is 0 degrees, and East is 90, South is 180, West is 270.
func (s *pen) CompassMode() {
	s.degreesEn = true
	s.compassEn = true
}

// Returns the pen's current angle mode
func (s *pen) GetAngleMode() turtlemodel.AngleMode {
	switch {
	case !s.degreesEn:
		return turtlemodel.RadiansMode
	case s.compassEn:
		return turtlemodel.CompassMode
	default:
		return turtlemodel.DegreesMode
	}
}

// Set the pen's speed in pixels per second.
func (s *pen) Speed(pixelsPerSecond float64) {
	if pixelsPerSecond < 1 {
		return
	}
	s.speed = pixelsPerSecond
}

// Get the pen's speed in pixels per second.
func (s *pen) GetSpeed() float64 {
	return s.speed
}

// Stop drawing. As the pen moves, it will no longer draw on the canvas.
func (s *pen) PenUp() {
	s.penDown = false
}

// Alias of PenUp
func (s *pen) PU() {
	s.PenUp()
}

// Alias of PenUp
func (s *pen) Off() {
	s.PenUp()
}

// Start drawing. As the pen moves, it will draw on the canvas.
func (s *pen) PenDown() {
	s.penDown = true
}

// Alias of PenDown
func (s *pen) PD() {
	s.PenDown()
}

// Alias of PenDown
func (s *pen) On() {
	s.PenDown()
}

// Return true if the pen is currently down.
func (s *pen) IsPenDown() bool {
	return s.penDown
}

// Set the pen's current color
func (s *pen) Color(c color.Color) {
	s.penColor = c
}

// Get the pen's current color
func (s *pen) GetColor() color.Color {
	return s.penColor
}

// Set the pen's draw size.
// As the pen moves, it will only draw the exact pixel it is on top of when the size is 0.
// For a size greater than 0, all pixels that are within the radius of size/2 will be colored.
func (s *pen) Size(size float64) {
	if size < 0 {
		return
	}
	s.penSize = size
}

// Get the currently set pen size
func (s *pen) GetSize() float64 {
	return s.penSize
}

// Draw the single pixel that the pen is currently on top of.
func (s *pen) Dot(size float64) {
	if size <= 0 {
		s.paintPixel(s.x, s.y, s.penColor)
		return
	}

	s.drawFilledCircle(s.x, s.y, size, s.penColor)
}

// Perform a fill starting at the turtles current location. All pixels that are connected and the same color will be changed to the new color.
func (s *pen) Fill(c color.Color) {
	x, y := floatPosToPixel(s.x, s.y)
	s.can.Fill(x, y, c)
}

// Draw a circle with given radius. The center is radius units left of the turtle; angleAmountToDraw determines
// which part of the circle is drawn. If angleAmountToDraw is not a full circle, one endpoint of the arc is
// the current pen position. Draw the arc in counterclockwise direction if radius is positive,
// otherwise in clockwise direction. Finally the direction of the turtle is changed by the amount of angleAmountToDraw.
//
// As the circle is approximated by an inscribed regular polygon, steps determines the number of steps to use.
// May be used to draw regular polygons.
func (s *pen) Circle(radius, angleAmountToDraw float64, steps int) {
	// Convert to radians
	if s.degreesEn {
		angleAmountToDraw *= (math.Pi / 180.0)
	}
	if angleAmountToDraw > math.Pi*2.0 {
		angleAmountToDraw = math.Pi * 2.0
	}
	if angleAmountToDraw < -math.Pi*2.0 {
		angleAmountToDraw = -math.Pi * 2.0
	}
	if radius < 0 {
		angleAmountToDraw *= -1
	}
	angleStepSize := angleAmountToDraw / float64(steps)
	endTurtleAngle := s.angle + angleAmountToDraw

	// Get center of Circle
	sin, cos := math.Sincos(s.angle + math.Pi/2.0)
	xCenter := s.x + radius*cos
	yCenter := s.y + radius*sin
	radius = math.Abs(radius)

	// Get the start of the circle
	deltaX := s.x - xCenter
	deltaY := s.y - yCenter
	startAngle := math.Atan2(deltaY, deltaX)

	for step := 1; step <= steps; step++ {
		currentAngle := startAngle + float64(step)*angleStepSize
		sin, cos := math.Sincos(currentAngle)
		x := xCenter + radius*cos
		y := yCenter + radius*sin
		s.angle += angleStepSize
		s.GoTo(x, y)
	}
	s.angle = endTurtleAngle
}

// Show the turtle sprite
func (s *pen) ShowTurtle() {
	s.visible = true
	s.sprite.SetVisible(true)
}

// Hide the turtle sprite
func (s *pen) HideTurtle() {
	s.visible = false
	s.sprite.SetVisible(false)
}

// Set the turtle sprite bitmap to be the built in turtle
func (s *pen) ShapeAsTurtle() {
	s.sprite.SetSpriteImageTurtle()
}

// Set the turtle sprite bitmap to be the built in arrow
func (s *pen) ShapeAsArrow() {
	s.sprite.SetSpriteImageArrow()
}

// Set the turtle sprite bitmap to the provided image
func (s *pen) ShapeAsImage(in image.Image) {
	s.sprite.SetSpriteImage(in)
}

// Set the turtle sprite bitmap scale
func (s *pen) ShapeScale(scale float64) {
	s.sprite.SetScale(scale)
}

// //////////////////////
func (s *pen) absoluteAngleToRad(angle float64) float64 {
	if s.degreesEn {
		angle *= (math.Pi / 180.0)

		if s.compassEn {
			angle = -angle + math.Pi/2
		}
	}

	return angle
}

func (s *pen) radToAbsoluteAngle(angle float64) float64 {
	if s.degreesEn {
		if s.compassEn {
			angle = -angle + math.Pi/2
		}
		angle *= (180.0 / math.Pi)
	}

	return angle
}

// This is what splits cartesian space into discrete pixels.
// This includes moving (0,0) be be centered in the middle of the (0,0) pixel. The center of the (0,0) pixel is at (.5, .5)
func floatPosToPixel(x, y float64) (int, int) {
	retX := int(math.Floor(x + .5))
	retY := int(math.Floor(y + .5))

	return retX, retY
}

func (s *pen) paintPixel(x, y float64, c color.Color) {
	pixX, pixY := floatPosToPixel(x, y)
	s.can.SetCartesianPixel(pixX, pixY, c)
}

// The concept of this line draw function is to determine if X or Y have a larger number of pixels to cover,
// and the larger one is chosen. Then we step
func (s *pen) drawLine(x1, y1, x2, y2 float64, c color.Color) {
	xDelta := x2 - x1
	yDelta := y2 - y1
	largerDelta := math.Max(math.Abs(xDelta), math.Abs(yDelta))

	loopSteps := int(math.Ceil(largerDelta))
	xStep := xDelta / float64(loopSteps)
	yStep := yDelta / float64(loopSteps)

	x := x1
	y := y1

	distance := math.Sqrt(xDelta*xDelta + yDelta*yDelta)
	timeToDraw := distance / s.speed
	timePerPixel := timeToDraw / largerDelta

	sleepTime := time.Duration(timePerPixel * 1.0e9)
	tNow := time.Now()

	for i := 0; i <= loopSteps; i++ {
		pixX, pixY := floatPosToPixel(x, y)
		s.can.SetCartesianPixel(pixX, pixY, c)
		if s.penSize > 0 {
			s.drawFilledCircle(x, y, s.penSize, c)
		}

		x += xStep
		y += yStep
		s.sprite.SetPosition(x, y)
		if s.speed < turtlemodel.MaxSpeed {
			tNow = tNow.Add(sleepTime)
			time.Sleep(time.Until(tNow))
		}
	}
}

func (s *pen) drawFilledCircle(xIn, yIn, size float64, c color.Color) {
	halfSize := size / 2
	halfSizeSquared := halfSize * halfSize
	xMax := int(math.Floor(xIn + halfSize))
	xMin := int(math.Floor(xIn - halfSize))
	yMax := int(math.Floor(yIn + halfSize))
	yMin := int(math.Floor(yIn - halfSize))

	for yInt := yMin; yInt <= yMax; yInt++ {
		for xInt := xMin; xInt <= xMax; xInt++ {
			deltaX := (float64(xInt) - xIn)
			deltaY := (float64(yInt) - yIn)
			distanceSquared := deltaX*deltaX + deltaY*deltaY
			if distanceSquared <= halfSizeSquared {
				s.can.SetCartesianPixel(xInt, yInt, c)
			}
		}
	}
}
