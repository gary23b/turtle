package main

import (
	"time"

	"github.com/gary23b/turtle"
)

// Drag the mouse with the left button pressed to draw on the canvas.
// Press c to clear the screen.
func main() {
	params := turtle.Params{Width: 1000, Height: 1000, ShowFPS: true}
	turtle.Start(params, drawFunc)
}

func drawFunc(window turtle.Window) {
	can := window.GetCanvas()
	can.ClearScreen(turtle.White)
	t := window.NewTurtle()
	t.Speed(1e9)

	justPressedChan := can.SubscribeToJustPressedUserInput()
	prevUserIn := can.PressedUserInput()
	ratio := 0.0

	for {
		userIn := can.PressedUserInput()
		justPressed := turtle.GetNewestJustPressedFromChan(justPressedChan)
		if justPressed != nil {
			if justPressed.Keys.C {
				can.ClearScreen(turtle.White)
			}
			if justPressed.IsPressedByName("Q") {
				can.Exit()
			}
		}

		if userIn.Mouse.Left {
			desiredColor := turtle.LerpColor(turtle.Blue, turtle.Red, ratio)
			t.Color(desiredColor)
			t.GoTo(float64(prevUserIn.Mouse.MouseX), float64(prevUserIn.Mouse.MouseY))
			t.PenDown()
			t.GoTo(float64(userIn.Mouse.MouseX), float64(userIn.Mouse.MouseY))
			t.PenUp()
		}
		prevUserIn = userIn
		ratio += .001
		if ratio > 1 {
			ratio = 0
		}
		time.Sleep(1 * time.Millisecond)
	}
}
