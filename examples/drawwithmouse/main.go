package main

import (
	"image/color"
	"time"
	"worldsim/ebitencanvas"
	"worldsim/models"
	"worldsim/turtleutil"
)

func main() {
	params := ebitencanvas.CanvasParams{Width: 1000, Height: 1000, ShowFPS: true}
	ebitencanvas.StartEbitenTurtleCanvas(params, drawFunc)
}

func drawFunc(can models.Canvas) {
	can.FillScreen(turtleutil.White)

	// width, height := can.GetWidth(), can.GetHeight()
	ratio := 0.0
	prevUserIn := &models.UserInput{}
	for {
		userIn := can.GetUserInput()
		if userIn.MouseDown.Left {
			desiredColor := turtleutil.LerpColor(turtleutil.Blue, turtleutil.Red, ratio)
			DrawLine(can, prevUserIn.MouseX, prevUserIn.MouseY, userIn.MouseX, userIn.MouseY, desiredColor)
		}
		prevUserIn = &userIn
		ratio += .001
		if ratio > 1 {
			ratio = 0
		}
		time.Sleep(1 * time.Millisecond)
	}

}

func DrawLine(can models.Canvas, x1, y1, x2, y2 int, c color.RGBA) {
	xDelta := x2 - x1
	yDelta := y2 - y1

	intAbs := func(in int) int {
		if in < 0 {
			return -in
		}
		return in
	}

	largerDelta := intAbs(xDelta)
	if intAbs(yDelta) > largerDelta {
		largerDelta = intAbs(yDelta)
	}

	xStep := float64(xDelta) / float64(largerDelta)
	yStep := float64(yDelta) / float64(largerDelta)

	x := float64(x1)
	y := float64(y1)
	for i := 0; i < largerDelta; i++ {
		can.SetPixel(int(x), int(y), c)
		x += xStep
		y += yStep
	}
}
