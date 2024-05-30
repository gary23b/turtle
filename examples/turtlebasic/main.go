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
		t.ShapeAsArrow()
		t.ShowTurtle()
		t.DegreesMode()
		t.PenDown()
		t.Forward(100)
		t.Left(90)
		t.Color(turtleutil.Red)
		t.Backward(50)
		t.PenUp()
		t.Forward(55)
		t.PenDown()
		t.Color(turtleutil.Black)
		t.Forward(45)

		t.Right(45)
		t.Forward(10)
		t.Backward(10)
		t.Color(turtleutil.Blue)

		t.PanLeftward(100)
		t.Forward(50)
		t.PanRightward(100)
	}()

	go func() {
		var t2 models.Turtle = turtle.NewTurtle(can)
		t2.RadiansMode()
		t2.GoTo(-200.0, 0.0)
		t2.PenDown()
		t2.Angle(math.Pi / 4)
		t2.Forward(100)
		t2.Angle(-math.Pi / 2)
		t2.Speed(500)
		t2.Forward(500)
		t2.Dot(40)
	}()

	go func() {
		var t3 models.Turtle = turtle.NewTurtle(can)
		t3.CompassMode()
		t3.PenDown()
		t3.Color(turtleutil.Green)
		t3.Size(20)
		t3.Angle(-45)
		t3.Forward(100)
		t3.GoTo(-100/math.Sqrt2, -100/math.Sqrt2)
		t3.Angle(45)
		t3.Forward(95)
	}()

	go func() {
		var t models.Turtle = turtle.NewTurtle(can)
		t.ShowTurtle()
		t.ShapeScale(.5)
		t.GoTo(300, -300)
		t.PenDown()
		t.Color(turtleutil.Red)
		penSize := 0.1
		panDistance := 0.5
		speed := 50.0
		t.Speed(speed)
		for i := 0; i < 110; i++ {
			t.PanRightward(panDistance)
			t.PointToward(300, -300)
			penSize *= 1.05
			t.Size(penSize)
			panDistance *= 1.05
			speed *= 1.02
			t.Speed(speed)
		}
	}()

	go func() {
		var t models.Turtle = turtle.NewTurtle(can)
		t.ShowTurtle()
		t.GoTo(-400, 300)
		t.CompassMode()
		t.PenDown()
		t.Color(turtleutil.Purple)
		t.Size(5)
		t.Speed(400)

		for i := 0; i < 15; i++ {
			t.Forward(250)
			t.Right(168)
		}
	}()
}
