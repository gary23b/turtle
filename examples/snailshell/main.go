package main

import (
	"github.com/gary23b/turtle"
)

func main() {
	params := turtle.Params{Width: 1000, Height: 1000}
	turtle.Start(params, drawFunc)
}

// drawFunc is started as a goroutine.
// https://codewithcurious.com/projects/turtle-patterns/
func drawFunc(window turtle.Window) {
	window.GetCanvas().ClearScreen(turtle.Black)
	t := window.NewTurtle()
	t.ShowTurtle()
	t.Speed(100000)
	t.PenDown()
	for i := 1; i < 160; i++ {
		t.Color(turtle.Red)
		t.Circle(float64(i), 360, 100)
		t.Color(turtle.Orange)
		t.Circle(float64(i)*.8, 360, 100)
		t.Right(3)
		t.Forward(3)
	}
}
