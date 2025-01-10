package main

import (
	"image/color"
	"math"
	"time"

	"github.com/gary23b/turtle"
)

// adapted from https://dev.to/giselamd/parallel-mandelbrot-set-in-go-480o

// threaded drawing of a Mandelbrot set

const (
	maxIter = 100
	samples = 10

	numBlocks  = 64
	numThreads = 32
)

type rgb struct {
	r int
	g int
	b int
}

var palette = []rgb{}

func main() {
	params := turtle.Params{Width: 1024, Height: 1024, ShowFPS: false}
	turtle.Start(params, turtleMain)
}

type WorkItem struct {
	x     float64
	y     float64
	scale float64

	maxIter int
	samples int

	initialX int
	finalX   int
	initialY int
	finalY   int
}

func turtleMain(window turtle.Window) {
	canvas := window.GetCanvas()
	canvas.ClearScreen(turtle.White)
	time.Sleep(time.Second * 1)

	for i := range 1000 {
		r, g, b := hslToRGB(float64(i)/1000.0, 1, .5)

		palette = append(palette, rgb{r: r, g: g, b: b})
	}

	// go func() {
	// 	time.Sleep(time.Second * 10)
	// 	turtle.TakeScreenshot(window, "./mandelbrot.png")
	// }()

	workBuffer := make(chan WorkItem, numBlocks)

	for range numThreads {
		go workerThread(window, workBuffer)
	}

	posX := -0.75
	posY := 0.0
	scale := 400.0
	drawMandelbrotInParallel(window, workBuffer, posX, posY, scale)

	justPressedChan := canvas.SubscribeToJustPressedUserInput()
	for {
		justPressed := turtle.GetNewestJustPressedFromChan(justPressedChan)
		if justPressed != nil && justPressed.Mouse.Left {
			posX += float64(justPressed.Mouse.MouseX) / scale
			posY += float64(justPressed.Mouse.MouseY) / scale
			scale *= 1.5
			drawMandelbrotInParallel(window, workBuffer, posX, posY, scale)
		}
		if justPressed != nil && justPressed.Mouse.Right {
			posX = -0.75
			posY = 0.0
			scale = 400.0
			drawMandelbrotInParallel(window, workBuffer, posX, posY, scale)
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func drawMandelbrotInParallel(window turtle.Window, workItemChan chan WorkItem, x, y, scale float64) {
	can := window.GetCanvas()
	imgWidth := can.GetWidth()
	imgHeight := can.GetHeight()

	blocksPerSide := math.Sqrt(numBlocks)
	pixelsPerSide := int(float64(imgWidth) / blocksPerSide)

	for pixX := -imgWidth / 2; pixX < imgWidth/2; pixX += pixelsPerSide {
		for pixY := -imgHeight / 2; pixY < imgHeight/2; pixY += pixelsPerSide {
			workItemChan <- WorkItem{
				x:       x,
				y:       y,
				scale:   scale,
				maxIter: maxIter,
				samples: samples,

				initialX: pixX - 1,
				finalX:   pixX + pixelsPerSide + 1,
				initialY: pixY - 1,
				finalY:   pixY + pixelsPerSide + 1,
			}
		}
	}
}

func workerThread(window turtle.Window, workItemChan chan WorkItem) {
	randSource := NewSavedRandoms()
	for {
		workItem, open := <-workItemChan
		if open {
			worker(window, randSource, &workItem)
		} else {
			return
		}
	}
}

func worker(window turtle.Window, randSource *SavedRandoms, in *WorkItem) {
	canvas := window.GetCanvas()
	for xPixelInt := in.initialX; xPixelInt < in.finalX; xPixelInt++ {
		xDelta := float64(xPixelInt) / in.scale
		for yPixelInt := in.initialY; yPixelInt < in.finalY; yPixelInt++ {
			yDelta := float64(yPixelInt) / in.scale
			var colorRedSum, colorGreenSum, colorBlueSum int
			for k := 0; k < in.samples; k++ {
				a := in.x + xDelta + randSource.RandFloat64()/in.scale
				b := in.y + yDelta + randSource.RandFloat64()/in.scale
				colorRed, colorGreen, colorBlue := pixelColor(mandelbrotInteraction(a, b, in.maxIter))
				colorRedSum += colorRed
				colorGreenSum += colorGreen
				colorBlueSum += colorBlue
			}
			cr := uint8(colorRedSum / in.samples)
			cg := uint8(colorGreenSum / in.samples)
			cb := uint8(colorBlueSum / in.samples)
			canvas.SetCartesianPixel(xPixelInt, yPixelInt, color.RGBA{cr, cg, cb, 255})
		}
	}
}

// https://en.wikipedia.org/wiki/Plotting_algorithms_for_the_Mandelbrot_set
func mandelbrotInteraction(a, b float64, maxIter int) (float64, int) {
	var x, y, xx, yy, xy float64

	for i := 0; i < maxIter; i++ {
		xx, yy, xy = x*x, y*y, x*y
		if xx+yy > 4 {
			return xx + yy, i
		}
		// xn+1 = x^2 - y^2 + a
		x = xx - yy + a
		// yn+1 = 2xy + b
		y = 2*xy + b
	}

	return xx + yy, maxIter
}

var log2 = math.Log(2.0)

func pixelColor(r float64, iter int) (int, int, int) {
	if iter >= maxIter {
		return 0, 0, 0
	}

	// r = x*x + y*y
	log_zn := math.Log(r) / 2.0
	nu := math.Log(log_zn/log2) / log2
	iterFloat := float64(iter) + 1 - nu
	iterRatio := iterFloat / float64(maxIter)

	c := palette[int(iterRatio*950)]
	return c.r, c.g, c.b
}
