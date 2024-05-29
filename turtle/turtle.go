package turtle

import (
	"image/color"
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
		can:      can,
		penColor: turtleutil.Black,
		speed:    75,
		penSize:  0,
	}

	return ret
}

func (s *turtle) Forward(distance float64) {
	s.goAngle(s.angle, distance)
}

func (s *turtle) Backward(distance float64) {
	s.goAngle(s.angle, -distance)
}

func (s *turtle) PanLeftward(distance float64) {
	s.goAngle(s.angle+math.Pi/2.0, distance)
}

func (s *turtle) PanRightward(distance float64) {
	s.goAngle(s.angle-math.Pi/2.0, distance)
}

func (s *turtle) goAngle(angle, distance float64) {
	x2 := s.x + distance*math.Cos(angle)
	y2 := s.y + distance*math.Sin(angle)
	s.GoTo(x2, y2)
}

func (s *turtle) GoTo(x, y float64) {
	if s.penDown {
		s.drawLine(int(s.x), int(s.y), int(x), int(y), s.penColor)
	}
	s.x = x
	s.y = y
}

func (s *turtle) GetPos() (x, y float64) {
	return s.x, s.y
}

func (s *turtle) Left(angle float64) {
	angle = s.angleToRad(angle)
	if s.compassEn {
		s.angle -= angle
	} else {
		s.angle += angle
	}
}

func (s *turtle) Right(angle float64) {
	s.Left(-angle)
}

func (s *turtle) SetAngle(angle float64) {
	angle = s.angleToRad(angle)
	s.angle = angle
}

func (s *turtle) PointToward(x, y float64) {
	deltaX := x - s.x
	deltaY := y - s.y
	s.angle = math.Atan2(deltaY, deltaX)
}

func (s *turtle) GetAngle() float64 {
	return s.radToAngle(s.angle)
}

func (s *turtle) SetDegreesMode() {
	s.degreesEn = true
}

func (s *turtle) SetRadianMode() {
	s.degreesEn = false
}

func (s *turtle) EnableCompassAngleMode(in bool) {
	s.compassEn = in
}

func (s *turtle) SetSpeed(PixelsPerSecond float64) {
	if PixelsPerSecond < 1 {
		return
	}
	s.speed = PixelsPerSecond
}

func (s *turtle) PenUp() {
	s.penDown = false
}
func (s *turtle) PenDown() {
	s.penDown = true
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

// //////////////////////
func (s *turtle) angleToRad(angle float64) float64 {
	if s.degreesEn {
		angle = angle * (math.Pi / 180.0)
	}
	if s.compassEn {
		angle = -angle + math.Pi/2
	}

	return angle
}

func (s *turtle) radToAngle(angle float64) float64 {
	if s.compassEn {
		angle = -angle + math.Pi/2
	}

	if s.degreesEn {
		angle = angle * (180.0 / math.Pi)
	}

	return angle
}

func (s *turtle) drawLine(x1, y1, x2, y2 int, c color.RGBA) {
	xDelta := x2 - x1
	yDelta := y2 - y1

	intAbs := func(in int) int {
		if in < 0 {
			return -in
		}
		return in
	}

	largerDelta := intAbs(xDelta)
	if intAbs(yDelta) > largerDelta {
		largerDelta = intAbs(yDelta)
	}

	xStep := float64(xDelta) / float64(largerDelta)
	yStep := float64(yDelta) / float64(largerDelta)

	x := float64(x1)
	y := float64(y1)

	distance := math.Sqrt(float64(xDelta*xDelta + yDelta*yDelta))
	timeToDraw := float64(distance) / s.speed
	timePerPixel := timeToDraw / float64(largerDelta)

	sleepTime := time.Duration(timePerPixel * 1.0e9)
	tNow := time.Now()

	for i := 0; i < largerDelta; i++ {
		s.can.SetCartesianPixel(int(x), int(y), c)
		if s.penSize > 0 {
			s.drawFilledCircle(x, y, s.penSize, c)
		}

		x += xStep
		y += yStep
		tNow = tNow.Add(sleepTime)
		time.Sleep(time.Until(tNow))
	}
}

func (s *turtle) drawFilledCircle(x, y, size float64, c color.RGBA) {
	halfSize := s.penSize / 2
	halfSizeSqrd := halfSize * halfSize
	xMax := int(x + halfSize)
	xMin := int(x - halfSize)
	yMax := int(y + halfSize)
	yMin := int(y - halfSize)

	for widthY := yMin; widthY <= yMax; widthY++ {
		for widthX := xMin; widthX <= xMax; widthX++ {
			deltaX := (float64(widthX) - x)
			deltaY := (float64(widthY) - y)
			distanceSqrd := deltaX*deltaX + deltaY*deltaY
			if distanceSqrd <= halfSizeSqrd {
				s.can.SetCartesianPixel(widthX, widthY, c)
			}

		}
	}
}
