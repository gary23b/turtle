package main

import (
	"math/rand"
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
	time.Sleep(time.Second * 3)
	can.FillScreen(turtleutil.Green)

	width, height := can.GetWidth(), can.GetHeight()

	// st := time.Now()

	for t := 0; t < 60*60; t++ {
		for i := 0; i < 3000; i++ {
			w := rand.Intn(width)
			h := rand.Intn(height)
			whichColor := rand.Intn(2)
			c := turtleutil.Water
			if whichColor == 1 {
				c = turtleutil.Black
			}

			can.SetPixel(w, h, c)
		}
		time.Sleep(time.Second / 75)
		// userIn := can.GetUserInput()
		// et := time.Now()
		// deltaTime := et.Sub(st)
		// refreshRate := 1.0 / deltaTime.Seconds()
		// st = et
		// fmt.Printf("%5.2f mouse x,y=%d,%d\n", refreshRate, userIn.MouseX, userIn.MouseY)
	}

}
