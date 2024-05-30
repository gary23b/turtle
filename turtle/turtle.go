package turtle

import (
	"image"
	"image/color"
	"math"
	"time"

	"github.com/gary23b/turtle/models"
	"github.com/gary23b/turtle/turtleutil"
)

type turtle struct {
	can       models.Canvas
	sprite    models.Sprite
	visible   bool
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
		sprite:    can.CreateNewSprite(),
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

	s.sprite.Set(s.visible, s.x, s.y, s.angle)
}

func (s *turtle) GetPos() (x, y float64) {
	return s.x, s.y
}

func (s *turtle) Left(angle float64) {
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

func (s *turtle) L(angle float64) {
	s.Left(angle)
}

func (s *turtle) Right(angle float64) {
	s.Left(-angle)
}

func (s *turtle) R(angle float64) {
	s.Right(angle)
}

func (s *turtle) Angle(angle float64) {
	angle = s.absoluteAngleToRad(angle)
	s.angle = angle
	s.sprite.SetRotation(s.angle)
}

func (s *turtle) GetAngle() float64 {
	return s.radToAbsoluteAngle(s.angle)
}

func (s *turtle) PointToward(x, y float64) {
	deltaX := x - s.x
	deltaY := y - s.y
	s.angle = math.Atan2(deltaY, deltaX)
	s.sprite.SetRotation(s.angle)
}

func (s *turtle) DegreesMode() {
	s.degreesEn = true
	s.compassEn = false
}

func (s *turtle) RadiansMode() {
	s.degreesEn = false
	s.compassEn = false
}

// In compass mode, North is 0 degrees, and East is 90, West is -90. Also forces Degrees mode on.
func (s *turtle) CompassMode() {
	s.degreesEn = true
	s.compassEn = true
}

func (s *turtle) GetAngleMode() models.AngleMode {
	switch {
	case !s.degreesEn:
		return models.RadiansMode
	case s.compassEn:
		return models.CompassMode
	default:
		return models.DegreesMode
	}
}

func (s *turtle) Speed(pixelsPerSecond float64) {
	if pixelsPerSecond < 1 {
		return
	}
	s.speed = pixelsPerSecond
}

func (s *turtle) GetSpeed() float64 {
	return s.speed
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

func (s *turtle) Color(c color.RGBA) {
	s.penColor = c
}

func (s *turtle) GetColor() color.RGBA {
	return s.penColor
}

func (s *turtle) Size(size float64) {
	if size < 0 {
		return
	}
	s.penSize = size
}

func (s *turtle) GetSize() float64 {
	return s.penSize
}

func (s *turtle) Dot(size float64) {
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

func (s *turtle) ShowTurtle() {
	s.visible = true
	s.sprite.SetVisible(true)
}

func (s *turtle) HideTurtle() {
	s.visible = false
	s.sprite.SetVisible(false)
}

func (s *turtle) ShapeAsTurtle() {
	s.sprite.SetSpriteImageTurtle()
}

func (s *turtle) ShapeAsArrow() {
	s.sprite.SetSpriteImageArrow()
}

func (s *turtle) ShapeAsImage(in image.Image) {
	s.sprite.SetSpriteImage(in)
}

func (s *turtle) ShapeScale(scale float64) {
	s.sprite.SetScale(scale)
}

// //////////////////////
func (s *turtle) absoluteAngleToRad(angle float64) float64 {
	if s.degreesEn {
		angle *= (math.Pi / 180.0)

		if s.compassEn {
			angle = -angle + math.Pi/2
		}
	}

	return angle
}

func (s *turtle) radToAbsoluteAngle(angle float64) float64 {
	if s.degreesEn {
		if s.compassEn {
			angle = -angle + math.Pi/2
		}
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
		s.sprite.SetPosition(x, y)
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
