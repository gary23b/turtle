package pen

import (
	"image"
	"image/color"
	"math"
	"testing"

	"github.com/gary23b/turtle/models"
	"github.com/gary23b/turtle/models/fakes"
	"github.com/gary23b/turtle/models/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	Water  color.RGBA = color.RGBA{0x23, 0x89, 0xDA, 0xFF} // 2389DA
	Black  color.RGBA = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	White  color.RGBA = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	Red    color.RGBA = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	Green  color.RGBA = color.RGBA{0x00, 0xFF, 0x00, 0xFF}
	Blue   color.RGBA = color.RGBA{0x00, 0x00, 0xFF, 0xFF}
	Purple color.RGBA = color.RGBA{0xFF, 0x00, 0xFF, 0xFF}
)

func Test_turtle_absoluteAngleToRad(t *testing.T) {
	type args struct {
		angle float64
	}
	tests := []struct {
		name string
		s    *pen
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "North",
			s:    &pen{degreesEn: true, compassEn: true},
			args: args{angle: 0},
			want: math.Pi / 2,
		},
		{
			name: "East",
			s:    &pen{degreesEn: true, compassEn: true},
			args: args{angle: 90},
			want: 0,
		},
		{
			name: "South",
			s:    &pen{degreesEn: true, compassEn: true},
			args: args{angle: 180},
			want: -math.Pi / 2,
		},
		{
			name: "West",
			s:    &pen{degreesEn: true, compassEn: true},
			args: args{angle: 270},
			want: -math.Pi,
		},
		{
			name: "North-West",
			s:    &pen{degreesEn: true, compassEn: true},
			args: args{angle: -45},
			want: 3.0 / 4.0 * math.Pi,
		},
		//
		{
			name: "0 degrees",
			s:    &pen{degreesEn: true},
			args: args{angle: 0},
			want: 0,
		},
		{
			name: "45 degrees",
			s:    &pen{degreesEn: true},
			args: args{angle: 45},
			want: math.Pi / 4,
		},
		{
			name: "-100 degrees",
			s:    &pen{degreesEn: true},
			args: args{angle: -100},
			want: -100.0 * math.Pi / 180.0,
		},
		{
			name: "360+45 degrees",
			s:    &pen{degreesEn: true},
			args: args{angle: 360 + 45},
			want: (360 + 45) * math.Pi / 180.0,
		},
		//
		{
			name: "360+45 degrees",
			s:    &pen{},
			args: args{angle: 1234.12345},
			want: 1234.12345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// require.InDelta
			got := tt.s.absoluteAngleToRad(tt.args.angle)
			require.Equal(t, tt.want, got)
			got2 := tt.s.radToAbsoluteAngle(got)
			require.Equal(t, tt.args.angle, got2)
		})
	}
}

func TestNewTurtleBasicTests(t *testing.T) {
	canvas := mocks.NewCanvas(t)
	sprite := fakes.NewSprite()
	canvas.On("CreateNewSprite").Return(sprite)
	bob := NewPen(canvas)

	require.Equal(t, &pen{
		can:       canvas,
		sprite:    sprite,
		x:         0,
		y:         0,
		angle:     0,
		degreesEn: true,
		compassEn: false,
		speed:     75,
		penDown:   false,
		penColor:  Black,
		penSize:   0,
	}, bob)

	// Test speed setting
	bob.Speed(100)
	require.Equal(t, 100.0, bob.GetSpeed())
	// Setting below 1 is not allowed
	bob.Speed(0)
	require.Equal(t, 100.0, bob.GetSpeed())
	bob.Speed(1)
	require.Equal(t, 1.0, bob.speed)

	// Test the Pen on/off control
	bob.PenDown()
	require.Equal(t, true, bob.penDown)
	bob.PenUp()
	require.Equal(t, false, bob.penDown)
	bob.PD()
	require.Equal(t, true, bob.penDown)
	bob.PU()
	require.Equal(t, false, bob.penDown)
	bob.On()
	require.Equal(t, true, bob.penDown)
	bob.Off()
	require.Equal(t, false, bob.penDown)

	// Test enabling and disabling angle modes
	bob.RadiansMode()
	require.Equal(t, false, bob.degreesEn)
	require.Equal(t, false, bob.compassEn)
	require.Equal(t, 0.0, bob.GetAngle())
	require.Equal(t, models.RadiansMode, bob.GetAngleMode())
	bob.CompassMode()
	require.Equal(t, true, bob.degreesEn)
	require.Equal(t, true, bob.compassEn)
	require.Equal(t, 90.0, bob.GetAngle())
	require.Equal(t, models.CompassMode, bob.GetAngleMode())
	bob.Angle(0)
	require.Equal(t, 0.0, bob.GetAngle())

	bob.DegreesMode()
	require.Equal(t, true, bob.degreesEn)
	require.Equal(t, false, bob.compassEn)
	require.Equal(t, 90.0, bob.GetAngle())
	require.Equal(t, models.DegreesMode, bob.GetAngleMode())
	bob.RadiansMode()
	require.Equal(t, false, bob.degreesEn)
	require.Equal(t, false, bob.compassEn)
	require.Equal(t, math.Pi/2, bob.GetAngle())
	require.Equal(t, models.RadiansMode, bob.GetAngleMode())

	// Test setting color and size
	bob.Color(Green)
	require.Equal(t, Green, bob.penColor)
	require.Equal(t, Green, bob.GetColor())
	bob.Size(1000)
	require.Equal(t, 1000.0, bob.penSize)
	require.Equal(t, 1000.0, bob.GetSize())
	// Setting size below 0 is not allowed
	bob.Size(-1)
	require.Equal(t, 1000.0, bob.penSize)
	bob.Size(0)
	require.Equal(t, 0.0, bob.penSize)

	x, y := bob.GetPos()
	require.Equal(t, 0.0, x)
	require.Equal(t, 0.0, y)
	bob.x = 1
	bob.y = 2
	x, y = bob.GetPos()
	require.Equal(t, 1.0, x)
	require.Equal(t, 2.0, y)
}

func TestNewTurtleTurning(t *testing.T) {
	canvas := mocks.NewCanvas(t)
	canvas.On("CreateNewSprite").Return(fakes.NewSprite())
	bob := NewPen(canvas)

	require.Equal(t, 0.0, bob.GetAngle())

	bob.Left(10)
	require.Equal(t, 10.0, bob.GetAngle())
	bob.L(10)
	require.Equal(t, 20.0, bob.GetAngle())
	bob.Right(10)
	require.Equal(t, 10.0, bob.GetAngle())
	bob.R(10)
	require.Equal(t, 0.0, bob.GetAngle())

	// Case rollovers
	bob.L(370)
	require.InDelta(t, 10, bob.GetAngle(), 1e-6)
	bob.Angle(0)
	bob.R(370)
	require.InDelta(t, -10.0, bob.GetAngle(), 1e-6)

	// Test PointToward
	bob.PointToward(1, 0)
	require.Equal(t, 0.0, bob.GetAngle())
	bob.PointToward(-1, 1)
	require.Equal(t, 135.0, bob.GetAngle())
	bob.PointToward(0, -1)
	require.Equal(t, -90.0, bob.GetAngle())

	bob.GoTo(1, 1)
	require.Equal(t, -90.0, bob.GetAngle())
	bob.PointToward(1, 0)
	require.Equal(t, -90.0, bob.GetAngle())
	bob.PointToward(0, 0)
	require.Equal(t, -135.0, bob.GetAngle())
	bob.PointToward(0, 2)
	require.Equal(t, 135.0, bob.GetAngle())
}

func TestNewTurtleMoveWoPen(t *testing.T) {
	canvas := mocks.NewCanvas(t)
	canvas.On("CreateNewSprite").Return(fakes.NewSprite())
	b := NewPen(canvas) // bob the turtle

	b.F(10)
	require.InDeltaSlice(t, []float64{10, 0}, []float64{b.x, b.y}, 1e-6)
	b.PanL(10)
	require.InDeltaSlice(t, []float64{10, 10}, []float64{b.x, b.y}, 1e-6)
	b.B(20)
	require.InDeltaSlice(t, []float64{-10, 10}, []float64{b.x, b.y}, 1e-6)
	b.PanR(20)
	require.InDeltaSlice(t, []float64{-10, -10}, []float64{b.x, b.y}, 1e-6)

	b.GoTo(0, 0)
	b.F(1)
	require.InDeltaSlice(t, []float64{1, 0}, []float64{b.x, b.y}, 1e-6)
	b.L(45)
	b.F(math.Sqrt2)
	require.InDeltaSlice(t, []float64{2, 1}, []float64{b.x, b.y}, 1e-6)
	b.R(90)
	b.F(math.Sqrt2)
	require.InDeltaSlice(t, []float64{3, 0}, []float64{b.x, b.y}, 1e-6)
}

type drawCmd struct {
	x           int
	y           int
	c           color.Color
	fill        bool // Bucket will starting from the x,y position
	clearScreen bool // When you want to set the entire screen to a color. Only fill in this when clearing screen.
}

type canvasFake struct {
	*mocks.Canvas
	calls []drawCmd
}

func newCanvasFake(t *testing.T) *canvasFake {
	t.Helper()
	ret := &canvasFake{
		Canvas: mocks.NewCanvas(t),
	}
	ret.Canvas.On("CreateNewSprite").Return(fakes.NewSprite())
	return ret
}

func (s *canvasFake) SetCartesianPixel(x int, y int, c color.Color) {
	s.calls = append(s.calls, drawCmd{x: x, y: y, c: c})
}

func (s *canvasFake) Fill(x int, y int, c color.Color) {
	s.calls = append(s.calls, drawCmd{x: x, y: y, c: c, fill: true})
}

func TestNewTurtleBasicDraw(t *testing.T) {
	canFake := newCanvasFake(t)
	b := NewPen(canFake) // bob the turtle

	b.GoTo(.1, .1) // move away from the pixel boundary
	b.Color(Black)
	b.On()
	b.F(1)
	require.Len(t, canFake.calls, 2)
	require.Equal(t, drawCmd{x: 0, y: 0, c: Black}, canFake.calls[0])
	require.Equal(t, drawCmd{x: 1, y: 0, c: Black}, canFake.calls[1])
	canFake.calls = nil

	b.Angle(90)
	b.F(3)
	require.Len(t, canFake.calls, 4)
	require.Equal(t, drawCmd{x: 1, y: 0, c: Black}, canFake.calls[0])
	require.Equal(t, drawCmd{x: 1, y: 1, c: Black}, canFake.calls[1])
	require.Equal(t, drawCmd{x: 1, y: 2, c: Black}, canFake.calls[2])
	require.Equal(t, drawCmd{x: 1, y: 3, c: Black}, canFake.calls[3])
	canFake.calls = nil

	b.Angle(180)
	b.Size(.001)
	b.F(2)
	require.Len(t, canFake.calls, 3)
	require.Equal(t, drawCmd{x: 1, y: 3, c: Black}, canFake.calls[0])
	require.Equal(t, drawCmd{x: 0, y: 3, c: Black}, canFake.calls[1])
	require.Equal(t, drawCmd{x: -1, y: 3, c: Black}, canFake.calls[2])
	canFake.calls = nil
}

func TestNewTurtleFilledCircleDraw(t *testing.T) {
	canvas := mocks.NewCanvas(t)
	canvas.On("CreateNewSprite").Return(fakes.NewSprite())
	canvas.On("SetCartesianPixel", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("color.RGBA")).Return()
	b := NewPen(canvas) // bob the turtle

	b.Color(Black)
	b.drawFilledCircle(0, 0, 2.1, Black)
	// fmt.Print(canvas.Calls)
	canvas.AssertCalled(t, "SetCartesianPixel", 0, 0, Black)
	canvas.AssertCalled(t, "SetCartesianPixel", 1, 0, Black)
	canvas.AssertCalled(t, "SetCartesianPixel", -1, 0, Black)
	canvas.AssertCalled(t, "SetCartesianPixel", 0, 1, Black)
	canvas.AssertCalled(t, "SetCartesianPixel", 0, -1, Black)
	canvas.AssertNumberOfCalls(t, "SetCartesianPixel", 5)
}

func Test_floatPosToPixel(t *testing.T) {
	// Test pixel (0,0)
	intX, intY := floatPosToPixel(0, 0)
	require.Equal(t, []int{0, 0}, []int{intX, intY})
	intX, intY = floatPosToPixel(0.49999999999, 0)
	require.Equal(t, []int{0, 0}, []int{intX, intY})
	intX, intY = floatPosToPixel(0.49999999999, 0.499999999999)
	require.Equal(t, []int{0, 0}, []int{intX, intY})
	intX, intY = floatPosToPixel(-.5, 0.499999999999)
	require.Equal(t, []int{0, 0}, []int{intX, intY})
	intX, intY = floatPosToPixel(-.5, -.5)
	require.Equal(t, []int{0, 0}, []int{intX, intY})
	intX, intY = floatPosToPixel(0.499999999999, -.5)
	require.Equal(t, []int{0, 0}, []int{intX, intY})
	intX, intY = floatPosToPixel(0.499999999999, 0.499999999999)
	require.Equal(t, []int{0, 0}, []int{intX, intY})

	// Test pixel (-1,-1)
	intX, intY = floatPosToPixel(-1, -1)
	require.Equal(t, []int{-1, -1}, []int{intX, intY})
	intX, intY = floatPosToPixel(-0.5000000001, -0.5000000001)
	require.Equal(t, []int{-1, -1}, []int{intX, intY})
	intX, intY = floatPosToPixel(-1.4999999999, -0.5000000001)
	require.Equal(t, []int{-1, -1}, []int{intX, intY})
	intX, intY = floatPosToPixel(-1.4999999999, -1.4999999999)
	require.Equal(t, []int{-1, -1}, []int{intX, intY})
	intX, intY = floatPosToPixel(-0.5000000001, -1.4999999999)
	require.Equal(t, []int{-1, -1}, []int{intX, intY})

	// Test pixel (2,-1)
	intX, intY = floatPosToPixel(2, -1)
	require.Equal(t, []int{2, -1}, []int{intX, intY})
	intX, intY = floatPosToPixel(2.4999999999, -0.5000000001)
	require.Equal(t, []int{2, -1}, []int{intX, intY})
	intX, intY = floatPosToPixel(1.5, -0.5000000001)
	require.Equal(t, []int{2, -1}, []int{intX, intY})
	intX, intY = floatPosToPixel(1.5, -1.4999999999)
	require.Equal(t, []int{2, -1}, []int{intX, intY})
	intX, intY = floatPosToPixel(2.4999999999, -1.4999999999)
	require.Equal(t, []int{2, -1}, []int{intX, intY})
}

func Test_Fill(t *testing.T) {
	canFake := newCanvasFake(t)
	b := NewPen(canFake) // bob the turtle
	b.Fill(Black)

	require.Len(t, canFake.calls, 1)
	require.Equal(t, drawCmd{x: 0, y: 0, c: Black, fill: true}, canFake.calls[0])
}

func Test_PaintDot(t *testing.T) {
	canFake := newCanvasFake(t)
	b := NewPen(canFake) // bob the turtle
	b.GoTo(5.5, 5.5)
	b.Dot(0)

	require.Len(t, canFake.calls, 1)
	require.Equal(t, drawCmd{x: 6, y: 6, c: Black}, canFake.calls[0])
	canFake.calls = nil

	b.Dot(2)

	require.Len(t, canFake.calls, 4)
	require.Equal(t, drawCmd{x: 5, y: 5, c: Black}, canFake.calls[0])
	require.Equal(t, drawCmd{x: 6, y: 5, c: Black}, canFake.calls[1])
	require.Equal(t, drawCmd{x: 5, y: 6, c: Black}, canFake.calls[2])
	require.Equal(t, drawCmd{x: 6, y: 6, c: Black}, canFake.calls[3])
	canFake.calls = nil
}

func Test_Circle(t *testing.T) {
	canFake := newCanvasFake(t)
	b := NewPen(canFake) // bob the turtle
	b.GoTo(10, 20)

	b.Circle(10, 360, 3)
	require.InDelta(t, 10, b.x, 1e-6)
	require.InDelta(t, 20, b.y, 1e-6)
	require.InDelta(t, 360, b.GetAngle(), 1e-6)

	b.Circle(10, 180, 3)
	require.InDelta(t, 10, b.x, 1e-6)
	require.InDelta(t, 40, b.y, 1e-6)
	require.InDelta(t, 360+180, b.GetAngle(), 1e-6)

	b.Angle(90)
	b.Circle(-10, 400, 3)
	require.InDelta(t, 10, b.x, 1e-6)
	require.InDelta(t, 40, b.y, 1e-6)
	require.InDelta(t, -360+90, b.GetAngle(), 1e-6)

	b.Angle(90)
	b.Circle(-10, -400, 3)
	require.InDelta(t, 10, b.x, 1e-6)
	require.InDelta(t, 40, b.y, 1e-6)
	require.InDelta(t, 360+90, b.GetAngle(), 1e-6)

	b.Angle(90)
	b.Circle(-10, 90, 3)
	require.InDelta(t, 20, b.x, 1e-6)
	require.InDelta(t, 50, b.y, 1e-6)
	require.InDelta(t, 0.0, b.GetAngle(), 1e-6)
}

func Test_SpriteControl(t *testing.T) {
	canFake := newCanvasFake(t)
	b := NewPen(canFake) // bob the turtle
	fakeSprite, ok := b.sprite.(*fakes.SpriteToDraw)
	require.True(t, ok)

	b.HideTurtle()
	require.Equal(t, false, fakeSprite.Visible)
	b.ShowTurtle()
	require.Equal(t, true, fakeSprite.Visible)
	b.ShapeAsArrow()
	require.Equal(t, "arrow", fakeSprite.CurrentImage)
	b.ShapeAsTurtle()
	require.Equal(t, "turtle", fakeSprite.CurrentImage)
	b.ShapeAsImage(image.NewRGBA(image.Rect(0, 0, 32, 32)))
	require.Equal(t, "custom", fakeSprite.CurrentImage)

	b.ShapeScale(100.9)
	require.Equal(t, 100.9, fakeSprite.Scale)
	b.GoTo(100, 150)
	b.Angle(2)
	require.Equal(t, 100.0, fakeSprite.X)
	require.Equal(t, 150.0, fakeSprite.Y)
	require.Equal(t, 2*math.Pi/180.0, fakeSprite.Angle)
}
