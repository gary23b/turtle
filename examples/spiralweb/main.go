package main

import (
	"image/color"

	"github.com/gary23b/turtle"
)

func main() {
	params := turtle.Params{Width: 1000, Height: 1000}
	turtle.Start(params, drawFunc)
}

// drawFunc is started as a goroutine.
// https://www.geeksforgeeks.org/draw-colorful-spiral-web-using-turtle-graphics-in-python/
func drawFunc(window turtle.Window) {
	window.GetCanvas().ClearScreen(turtle.Black)
	t := window.NewTurtle()
	t.ShowTurtle()
	t.Speed(1000)
	t.PenDown()

	colors := []color.RGBA{turtle.Red, turtle.Yellow, turtle.Green, turtle.Purple, turtle.Blue, turtle.Orange}
	for x := 1; x < 200; x++ {
		t.Color(colors[x%6])        // setting color
		t.Size(float64(x)/75.0 + 1) // setting width
		t.Forward(float64(x))       // moving forward
		t.Left(59)                  // Turt left
	}
}
