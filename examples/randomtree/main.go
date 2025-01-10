package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/gary23b/turtle"
	"github.com/gary23b/turtle/turtlemodel"
)

func main() {
	params := turtle.Params{Width: 600, Height: 600}
	turtle.Start(params, turtleMain)
}

func turtleMain(window turtle.Window) {
	canvas := window.GetCanvas()
	canvas.ClearScreen(color.RGBA{R: 113, G: 170, B: 227, A: 255})
	time.Sleep(time.Second * 1)

	// go turtle.CreateGif(window, time.Millisecond*20, time.Millisecond*20, "./randomtree.gif", 150)

	drawNewTree(window)

	justPressedChan := canvas.SubscribeToJustPressedUserInput()
	for {
		justPressed := turtle.GetNewestJustPressedFromChan(justPressedChan)
		if justPressed != nil && justPressed.Mouse.Left {
			drawNewTree(window)
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func drawNewTree(window turtle.Window) {
	canvas := window.GetCanvas()
	canvas.ClearScreen(color.RGBA{R: 113, G: 170, B: 227, A: 255})
	t := window.NewTurtle()
	t.DegreesMode()
	t.Speed(turtlemodel.MaxSpeed)

	RandomTree(t, 0, -300, 90, 30, 90, 10)
}

func ToRadian(in float64) float64 {
	return in * math.Pi / 180.0
}

func RandomTree(t turtlemodel.Turtle, x, y, l, width, angle float64, dp int) {
	t.PenUp()
	t.Angle(angle)
	t.Teleport(x, y)
	t.Size(width)

	if dp <= 2 {
		g := uint8((.2 + rand.Float64()/4.0) * 256)
		t.Color(color.RGBA{R: 0, G: g, B: 0, A: 0xFF})
	} else {
		r := uint8((.15 + rand.Float64()/8.0) * 256)
		t.Color(color.RGBA{R: r, G: 25, B: 0, A: 0xFF})
	}

	t.PenDown()

	x += l * math.Cos(ToRadian(angle))
	y += l * math.Sin(ToRadian(angle))
	t.GoTo(x, y)

	if dp == 1 {
		t.HideTurtle()
		return

	}

	nb := 2
	if dp < 9 {
		nb++
	}

	for i := 1; i <= nb; i++ {
		an := angle + (float64(i)-float64(nb)/2.0-0.5)*20.0
		an += (rand.Float64() - .5) * 60.0
		ln := l * (0.7 + rand.Float64()*0.3)

		RandomTree(t, x, y, ln, width*.7, an, dp-1)
	}

}
