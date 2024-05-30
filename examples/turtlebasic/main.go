package main

import (
	"math"
	"time"

	"github.com/gary23b/turtle"
)

// Draws some basic shapes with 5 different turtles.
// Part of what this shows is how each turtle can run at the same time.
// They also can be programmed completely independently.
func main() {
	params := turtle.Params{Width: 1000, Height: 1000, ShowFPS: true}
	turtle.Start(params, drawFunc)
}

func drawFunc(window turtle.Window) {
	time.Sleep(time.Second * 1)

	go func() {
		t := window.NewTurtle()
		t.ShapeAsArrow()
		t.ShowTurtle()
		t.DegreesMode()
		t.PenDown()
		t.Forward(100)
		t.Left(90)
		t.Color(turtle.Red)
		t.Backward(50)
		t.PenUp()
		t.Forward(55)
		t.PenDown()
		t.Color(turtle.Black)
		t.Forward(45)

		t.Right(45)
		t.Forward(10)
		t.Backward(10)
		t.Color(turtle.Blue)

		t.PanLeftward(100)
		t.Forward(50)
		t.PanRightward(100)
	}()

	go func() {
		t := window.NewTurtle()
		t.RadiansMode()
		t.GoTo(-200.0, 0.0)
		t.PenDown()
		t.Angle(math.Pi / 4)
		t.Forward(100)
		t.Angle(-math.Pi / 2)
		t.Speed(500)
		t.Forward(500)
		t.Dot(40)
	}()

	go func() {
		t := window.NewTurtle()
		t.CompassMode()
		t.PenDown()
		t.Color(turtle.Green)
		t.Size(20)
		t.Angle(-45)
		t.Forward(100)
		t.GoTo(-100/math.Sqrt2, -100/math.Sqrt2)
		t.Angle(45)
		t.Forward(95)
	}()

	go func() {
		t := window.NewTurtle()
		t.ShowTurtle()
		t.ShapeScale(.5)
		t.GoTo(300, -300)
		t.PenDown()
		t.Color(turtle.Red)
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
		t := window.NewTurtle()
		t.ShowTurtle()
		t.GoTo(-400, 300)
		t.CompassMode()
		t.PenDown()
		t.Color(turtle.Purple)
		t.Size(5)
		t.Speed(400)

		for i := 0; i < 15; i++ {
			t.Forward(250)
			t.Right(168)
		}
	}()
}
