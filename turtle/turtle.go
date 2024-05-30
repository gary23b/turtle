package turtle

import (
	"image/color"
	"log"
	"math"
	"time"

	"github.com/gary23b/turtle/models"
	"github.com/gary23b/turtle/turtleutil"
)

type turtle struct {
	can       models.Canvas
	x         float64 // in pixels
	y         float64 // in pixels
	angle     float64 // in radians
	degreesEn bool
	compassEn bool
	speed     float64 // in pixels/second
	penDown   bool
	penColor  color.RGBA
	penSize   float64
}

var _ models.Turtle = &turtle{} // Force the linter to tell us if the interface is implemented

func NewTurtle(can models.Canvas) *turtle {
	ret := &turtle{
		can:       can,
		penColor:  turtleutil.Black,
		speed:     75,
		penSize:   0,
		degreesEn: true,
	}

	return ret
}

func (s *turtle) Forward(distance float64) {
	s.goAngle(s.angle, distance)
}

func (s *turtle) F(distance float64) {
	s.Forward(distance)
}

func (s *turtle) Backward(distance float64) {
	s.goAngle(s.angle, -distance)
}

func (s *turtle) B(distance float64) {
	s.Backward(distance)
}

func (s *turtle) PanLeftward(distance float64) {
	s.goAngle(s.angle+math.Pi/2.0, distance)
}

func (s *turtle) PanL(distance float64) {
	s.PanLeftward(distance)
}

func (s *turtle) PanRightward(distance float64) {
	s.goAngle(s.angle-math.Pi/2.0, distance)
}

func (s *turtle) PanR(distance float64) {
	s.PanRightward(distance)
}

func (s *turtle) goAngle(angle, distance float64) {
	sin, cos := math.Sincos(angle)
	x2 := s.x + distance*cos
	y2 := s.y + distance*sin
	s.GoTo(x2, y2)
}

func (s *turtle) GoTo(x, y float64) {
	if s.penDown {
		s.drawLine(s.x, s.y, x, y, s.penColor)
	}
	s.x = x
	s.y = y
}

func (s *turtle) GetPos() (x, y float64) {
	return s.x, s.y
}

func (s *turtle) Left(angle float64) {
	if s.degreesEn {
		angle *= (math.Pi / 180.0)
	}
	s.angle += angle

	// Normaize the angle
	if s.angle > 2*math.Pi {
		s.angle -= 2 * math.Pi
	} else if s.angle < -2*math.Pi {
		s.angle += 2 * math.Pi
	}
}

func (s *turtle) L(angle float64) {
	s.Left(angle)
}

func (s *turtle) Right(angle float64) {
	s.Left(-angle)
}

func (s *turtle) R(angle float64) {
	s.Right(angle)
}

func (s *turtle) SetAngle(angle float64) {
	angle = s.absoluteAngleToRad(angle)
	s.angle = angle
}

func (s *turtle) PointToward(x, y float64) {
	deltaX := x - s.x
	deltaY := y - s.y
	s.angle = math.Atan2(deltaY, deltaX)
}

func (s *turtle) GetAngle() float64 {
	return s.radToAbsoluteAngle(s.angle)
}

func (s *turtle) SetDegreesMode() {
	s.degreesEn = true
}

func (s *turtle) SetRadianMode() {
	if s.compassEn {
		log.Println("ERROR: Radian mode cannot be enabled while compass mode is active.")
		return
	}
	s.degreesEn = false
}

// In compass mode, North is 0 degrees, and East is 90, West is -90. Also forces Degrees mode on.
func (s *turtle) EnableCompassAngleMode(in bool) {
	if in {
		s.SetDegreesMode()
	}

	s.compassEn = in
}

func (s *turtle) SetSpeed(pixelsPerSecond float64) {
	if pixelsPerSecond < 1 {
		return
	}
	s.speed = pixelsPerSecond
}

func (s *turtle) PenUp() {
	s.penDown = false
}

func (s *turtle) PU() {
	s.PenUp()
}

func (s *turtle) Off() {
	s.PenUp()
}

func (s *turtle) PenDown() {
	s.penDown = true
}

func (s *turtle) PD() {
	s.PenDown()
}

func (s *turtle) On() {
	s.PenDown()
}

func (s *turtle) PenColor(c color.RGBA) {
	s.penColor = c
}

func (s *turtle) PenSize(size float64) {
	if size < 0 {
		return
	}
	s.penSize = size
}

func (s *turtle) PaintDot(size float64) {
	if size <= 0 {
		s.paintPixel(s.x, s.y, s.penColor)
		return
	}

	s.drawFilledCircle(s.x, s.y, size, s.penColor)
}

func (s *turtle) Fill(c color.RGBA) {
	x, y := floatPosToPixel(s.x, s.y)
	s.can.Fill(x, y, c)
}

func (s *turtle) Circle(radius, angleAmountToDraw float64, steps int) {
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
		sin, cos := math.Sincos(startAngle + float64(step)*angleStepSize)
		x := xCenter + radius*cos
		y := yCenter + radius*sin
		s.GoTo(x, y)
	}
	s.angle += angleAmountToDraw
}

// //////////////////////
func (s *turtle) absoluteAngleToRad(angle float64) float64 {
	if s.degreesEn {
		angle *= (math.Pi / 180.0)
	}
	if s.compassEn {
		angle = -angle + math.Pi/2
	}

	return angle
}

func (s *turtle) radToAbsoluteAngle(angle float64) float64 {
	if s.compassEn {
		angle = -angle + math.Pi/2
	}

	if s.degreesEn {
		angle *= (180.0 / math.Pi)
	}

	return angle
}

// This is what splits cartesian space into discrete pixels.
// This includes moveing (0,0) be be centered in the middle of the (0,0) pixel. The center of the (0,0) pixel is at (.5, .5)
func floatPosToPixel(x, y float64) (int, int) {
	retX := int(math.Floor(x + .5))
	retY := int(math.Floor(y + .5))

	return retX, retY
}

func (s *turtle) paintPixel(x, y float64, c color.RGBA) {
	pixX, pixY := floatPosToPixel(x, y)
	s.can.SetCartesianPixel(pixX, pixY, c)
}

// The concept of this line draw function is to determine if X or Y have a larger number of pixels to cover,
// and the larger one is chosen. Then we step
func (s *turtle) drawLine(x1, y1, x2, y2 float64, c color.RGBA) {
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
		tNow = tNow.Add(sleepTime)
		time.Sleep(time.Until(tNow))
	}
}

func (s *turtle) drawFilledCircle(xIn, yIn, size float64, c color.RGBA) {
	halfSize := size / 2
	halfSizeSqrd := halfSize * halfSize
	xMax := int(math.Floor(xIn + halfSize))
	xMin := int(math.Floor(xIn - halfSize))
	yMax := int(math.Floor(yIn + halfSize))
	yMin := int(math.Floor(yIn - halfSize))

	for yInt := yMin; yInt <= yMax; yInt++ {
		for xInt := xMin; xInt <= xMax; xInt++ {
			deltaX := (float64(xInt) - xIn)
			deltaY := (float64(yInt) - yIn)
			distanceSqrd := deltaX*deltaX + deltaY*deltaY
			if distanceSqrd <= halfSizeSqrd {
				s.can.SetCartesianPixel(xInt, yInt, c)
			}
		}
	}
}
