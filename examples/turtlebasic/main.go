package main

import (
	"math"
	"time"

	"github.com/gary23b/turtle/ebitencanvas"
	"github.com/gary23b/turtle/models"
	"github.com/gary23b/turtle/turtle"
	"github.com/gary23b/turtle/turtleutil"
)

// Draws some basic shapes with 5 different turtles.
// Part of what this shows is how each turtle can run at the same time.
// They also can be programmed completely independently.
func main() {
	params := ebitencanvas.CanvasParams{Width: 1000, Height: 1000, ShowFPS: true}
	ebitencanvas.StartEbitenTurtleCanvas(params, drawFunc)
}

func drawFunc(can models.Canvas) {
	can.ClearScreen(turtleutil.White)
	time.Sleep(time.Second * 1)

	go func() {
		var t models.Turtle = turtle.NewTurtle(can)
		t.SetShapeAsArrow()
		t.SetVisible(true)
		t.SetDegreesMode()
		t.PenDown()
		t.Forward(100)
		t.Left(90)
		t.PenColor(turtleutil.Red)
		t.Backward(50)
		t.PenUp()
		t.Forward(55)
		t.PenDown()
		t.PenColor(turtleutil.Black)
		t.Forward(45)

		t.Right(45)
		t.Forward(10)
		t.Backward(10)
		t.PenColor(turtleutil.Blue)

		t.PanLeftward(100)
		t.Forward(50)
		t.PanRightward(100)
	}()

	go func() {
		var t2 models.Turtle = turtle.NewTurtle(can)
		t2.SetRadianMode()
		t2.GoTo(-200.0, 0.0)
		t2.PenDown()
		t2.SetAngle(math.Pi / 4)
		t2.Forward(100)
		t2.SetAngle(-math.Pi / 2)
		t2.SetSpeed(500)
		t2.Forward(500)
		t2.PaintDot(40)
	}()

	go func() {
		var t3 models.Turtle = turtle.NewTurtle(can)
		t3.SetDegreesMode()
		t3.EnableCompassAngleMode(true)
		t3.PenDown()
		t3.PenColor(turtleutil.Green)
		t3.PenSize(20)
		t3.SetAngle(-45)
		t3.Forward(100)
		t3.GoTo(-100/math.Sqrt2, -100/math.Sqrt2)
		t3.SetAngle(45)
		t3.Forward(95)
	}()

	go func() {
		var t models.Turtle = turtle.NewTurtle(can)
		t.SetVisible(true)
		t.SetShapeScale(.5)
		t.GoTo(300, -300)
		t.PenDown()
		t.PenColor(turtleutil.Red)
		penSize := 0.1
		panDistance := 0.5
		speed := 50.0
		t.SetSpeed(speed)
		for i := 0; i < 110; i++ {
			t.PanRightward(panDistance)
			t.PointToward(300, -300)
			penSize *= 1.05
			t.PenSize(penSize)
			panDistance *= 1.05
			speed *= 1.02
			t.SetSpeed(speed)
		}
	}()

	go func() {
		var t models.Turtle = turtle.NewTurtle(can)
		t.SetVisible(true)
		t.GoTo(-400, 300)
		t.EnableCompassAngleMode(true)
		t.PenDown()
		t.PenColor(turtleutil.Purple)
		t.PenSize(5)
		t.SetSpeed(400)

		for i := 0; i < 15; i++ {
			t.Forward(250)
			t.Right(168)
		}
	}()
}
