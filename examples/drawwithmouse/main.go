package main

import (
	"time"

	"github.com/gary23b/turtle"
	"github.com/gary23b/turtle/models"
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

	ratio := 0.0
	prevUserIn := &models.UserInput{}
	for {
		userIn := can.GetUserInput()
		if userIn.KeysDown.C && !prevUserIn.KeysDown.C {
			can.ClearScreen(turtle.White)
		}

		if userIn.MouseDown.Left {
			desiredColor := turtle.LerpColor(turtle.Blue, turtle.Red, ratio)
			t.Color(desiredColor)
			t.GoTo(float64(prevUserIn.MouseX), float64(prevUserIn.MouseY))
			t.PenDown()
			t.GoTo(float64(userIn.MouseX), float64(userIn.MouseY))
			t.PenUp()
		}
		prevUserIn = &userIn
		ratio += .001
		if ratio > 1 {
			ratio = 0
		}
		time.Sleep(1 * time.Millisecond)
	}
}
