package main

import (
	"math"
	"time"

	"github.com/gary23b/turtle"
	"github.com/gary23b/turtle/tools"
)

// Try to touch the worm. It runs away. The closer you get the faster it goes. But if you catch it, it explodes.
// Press c to clear the screen.
func main() {
	params := turtle.Params{Width: 1000, Height: 1000, ShowFPS: true}
	turtle.Start(params, drawFunc)
}

func drawFunc(window turtle.Window) {
	can := window.GetCanvas()
	can.ClearScreen(turtle.White)
	t := window.NewTurtle()
	t.ShowTurtle()
	t.PenDown()
	t.Size(5)
	speed := 300.0
	justPressedChan := can.SubscribeToJustPressedUserInput()
	time.Sleep(time.Second)

	for {
		justPressed := turtle.GetNewestJustPressedFromChan(justPressedChan)
		if justPressed != nil {
			if justPressed.Keys.C {
				can.ClearScreen(turtle.White)
			}
			if justPressed.Keys.Q {
				can.Exit()
			}
		}
		userIn := can.PressedUserInput()
		tX, tY := t.GetPos()
		mX, mY := float64(userIn.Mouse.MouseX), float64(userIn.Mouse.MouseY)

		t.PointToward(mX, mY)
		t.Right(180)

		deltaX := tX - mX
		deltaY := tY - mY
		dist := math.Sqrt(deltaX*deltaX + deltaY*deltaY)

		if dist < 300 {
			ratio := 2.0 * (1.0 - dist/100.0)
			switch {
			case ratio <= 1:
				t.Color(tools.LerpColor(turtle.Blue, turtle.Green, ratio))
			case ratio > 1:
				t.Color(tools.LerpColor(turtle.Green, turtle.Red, ratio-1.0))
			}

			if dist < 5 {
				t.Dot(100)
			} else {
				speed := 9.0 * 300.0 / dist
				t.Speed(speed)
			}
			t.Forward(speed * 0.010)
		}

		time.Sleep(10 * time.Millisecond)
	}
}
