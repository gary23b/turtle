package turtle

import (
	"image/color"
	"math"
	"testing"

	"github.com/gary23b/turtle/mocks"
	"github.com/gary23b/turtle/turtleutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_turtle_absoluteAngleToRad(t *testing.T) {
	type args struct {
		angle float64
	}
	tests := []struct {
		name string
		s    *turtle
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "North",
			s:    &turtle{degreesEn: true, compassEn: true},
			args: args{angle: 0},
			want: math.Pi / 2,
		},
		{
			name: "East",
			s:    &turtle{degreesEn: true, compassEn: true},
			args: args{angle: 90},
			want: 0,
		},
		{
			name: "South",
			s:    &turtle{degreesEn: true, compassEn: true},
			args: args{angle: 180},
			want: -math.Pi / 2,
		},
		{
			name: "West",
			s:    &turtle{degreesEn: true, compassEn: true},
			args: args{angle: 270},
			want: -math.Pi,
		},
		{
			name: "North-West",
			s:    &turtle{degreesEn: true, compassEn: true},
			args: args{angle: -45},
			want: 3.0 / 4.0 * math.Pi,
		},
		//
		{
			name: "0 degrees",
			s:    &turtle{degreesEn: true},
			args: args{angle: 0},
			want: 0,
		},
		{
			name: "45 degrees",
			s:    &turtle{degreesEn: true},
			args: args{angle: 45},
			want: math.Pi / 4,
		},
		{
			name: "-100 degrees",
			s:    &turtle{degreesEn: true},
			args: args{angle: -100},
			want: -100.0 * math.Pi / 180.0,
		},
		{
			name: "360+45 degrees",
			s:    &turtle{degreesEn: true},
			args: args{angle: 360 + 45},
			want: (360 + 45) * math.Pi / 180.0,
		},
		//
		{
			name: "360+45 degrees",
			s:    &turtle{},
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
	bob := NewTurtle(canvas)

	require.Equal(t, &turtle{
		can:       canvas,
		x:         0,
		y:         0,
		angle:     0,
		degreesEn: true,
		compassEn: false,
		speed:     75,
		penDown:   false,
		penColor:  turtleutil.Black,
		penSize:   0,
	}, bob)

	// Test speed setting
	bob.SetSpeed(100)
	require.Equal(t, 100.0, bob.speed)
	// Setting below 1 is not allowed
	bob.SetSpeed(0)
	require.Equal(t, 100.0, bob.speed)
	bob.SetSpeed(1)
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
	bob.SetRadianMode()
	require.Equal(t, false, bob.degreesEn)
	require.Equal(t, 0.0, bob.GetAngle())
	bob.EnableCompassAngleMode(true)
	require.Equal(t, true, bob.degreesEn)
	require.Equal(t, true, bob.compassEn)
	require.Equal(t, 90.0, bob.GetAngle())
	bob.SetAngle(0)
	require.Equal(t, 0.0, bob.GetAngle())

	// enabling radian mode is not allowed while in compass mode.
	bob.SetRadianMode()
	require.Equal(t, true, bob.degreesEn)
	require.Equal(t, true, bob.compassEn)
	bob.EnableCompassAngleMode(false)
	require.Equal(t, true, bob.degreesEn)
	require.Equal(t, false, bob.compassEn)
	require.Equal(t, 90.0, bob.GetAngle())
	bob.SetRadianMode()
	require.Equal(t, false, bob.degreesEn)
	require.Equal(t, false, bob.compassEn)
	require.Equal(t, math.Pi/2, bob.GetAngle())

	// Test setting color and size
	bob.PenColor(turtleutil.Green)
	require.Equal(t, turtleutil.Green, bob.penColor)
	bob.PenSize(1000)
	require.Equal(t, 1000.0, bob.penSize)
	// Setting size below 0 is not allowed
	bob.PenSize(-1)
	require.Equal(t, 1000.0, bob.penSize)
	bob.PenSize(0)
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
	bob := NewTurtle(canvas)

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
	bob.SetAngle(0)
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
	b := NewTurtle(canvas) // bob the turtle

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
	x int
	y int
	c color.RGBA
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
	return ret
}

func (s *canvasFake) SetCartesianPixel(x int, y int, c color.RGBA) {
	s.calls = append(s.calls, drawCmd{x: x, y: y, c: c})
}

func TestNewTurtleBasicDraw(t *testing.T) {
	canFake := newCanvasFake(t)
	b := NewTurtle(canFake) // bob the turtle

	b.GoTo(.1, .1) // move away from the pixel boundary
	b.PenColor(turtleutil.Black)
	b.On()
	b.F(1)
	require.Len(t, canFake.calls, 2)
	require.Equal(t, drawCmd{x: 0, y: 0, c: turtleutil.Black}, canFake.calls[0])
	require.Equal(t, drawCmd{x: 1, y: 0, c: turtleutil.Black}, canFake.calls[1])
	canFake.calls = nil

	b.SetAngle(90)
	b.F(3)
	require.Len(t, canFake.calls, 4)
	require.Equal(t, drawCmd{x: 1, y: 0, c: turtleutil.Black}, canFake.calls[0])
	require.Equal(t, drawCmd{x: 1, y: 1, c: turtleutil.Black}, canFake.calls[1])
	require.Equal(t, drawCmd{x: 1, y: 2, c: turtleutil.Black}, canFake.calls[2])
	require.Equal(t, drawCmd{x: 1, y: 3, c: turtleutil.Black}, canFake.calls[3])
	canFake.calls = nil

	b.SetAngle(180)
	b.PenSize(.001)
	b.F(2)
	require.Len(t, canFake.calls, 3)
	require.Equal(t, drawCmd{x: 1, y: 3, c: turtleutil.Black}, canFake.calls[0])
	require.Equal(t, drawCmd{x: 0, y: 3, c: turtleutil.Black}, canFake.calls[1])
	require.Equal(t, drawCmd{x: -1, y: 3, c: turtleutil.Black}, canFake.calls[2])
	canFake.calls = nil
}

func TestNewTurtleFilledCircleDraw(t *testing.T) {
	canvas := mocks.NewCanvas(t)
	canvas.On("SetCartesianPixel", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("color.RGBA")).Return()
	b := NewTurtle(canvas) // bob the turtle

	b.PenColor(turtleutil.Black)
	b.drawFilledCircle(0, 0, 2.1, turtleutil.Black)
	// fmt.Print(canvas.Calls)
	canvas.AssertCalled(t, "SetCartesianPixel", 0, 0, turtleutil.Black)
	canvas.AssertCalled(t, "SetCartesianPixel", 1, 0, turtleutil.Black)
	canvas.AssertCalled(t, "SetCartesianPixel", -1, 0, turtleutil.Black)
	canvas.AssertCalled(t, "SetCartesianPixel", 0, 1, turtleutil.Black)
	canvas.AssertCalled(t, "SetCartesianPixel", 0, -1, turtleutil.Black)
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
