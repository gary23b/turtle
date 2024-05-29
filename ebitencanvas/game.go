package ebitencanvas

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/gary23b/turtle/models"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

/*
The base Ebiten "Game" is based off the example https://ebitengine.org/en/examples/noise.html
*/

type game struct {
	screenWidth  int
	screenHeight int
	showFPS      bool
	commands     chan drawCmd
	img          *image.RGBA
	controlState SavedControlState
	controls     *models.UserInput
}

func newGame(width, height int, showFPS bool, commands chan drawCmd) *game {
	g := &game{
		screenWidth:  width,
		screenHeight: height,
		showFPS:      showFPS,
		commands:     commands,
		img:          image.NewRGBA(image.Rect(0, 0, width, height)),
	}

	// ebiten.SetTPS(120)
	// ebiten.SetVsyncEnabled(false) // For some reason, on Windows, there is quite a bit of lag.
	//setting this to false clears it up, but also makes it run at 1000Hz...
	ebiten.SetWindowSize(g.screenWidth, g.screenHeight)
	ebiten.SetWindowTitle("Go Turtle Graphics")
	return g
}

// This function will not return. It must be run on the main thread.
func (g *game) runGame() {
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func clearImage(i *image.RGBA, c color.RGBA) {
	xMin := i.Rect.Min.X
	yMin := i.Rect.Min.Y
	xMax := i.Rect.Max.X
	yMax := i.Rect.Max.Y
	for y := yMin; y < yMax; y++ {
		for x := xMin; x < xMax; x++ {
			i.SetRGBA(x, y, c)
		}
	}
}

func (g *game) Update() error {
	g.controls = g.controlState.GetUserInput()

	// Pull all the items out of the command channel until it is empty.
EatDrawCommandsLoop:
	for {
		select {
		case cmd := <-g.commands:
			if cmd.clearScreen != nil {
				clearImage(g.img, *cmd.clearScreen)
			} else {
				g.img.SetRGBA(cmd.x, cmd.y, cmd.c)
			}

		default:
			// receiving from g.commands would block
			break EatDrawCommandsLoop
		}
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {

	// screen.DrawTriangles

	screen.WritePixels(g.img.Pix)
	if g.showFPS {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *game) getUserInput() models.UserInput {
	if g == nil || g.controls == nil {
		return models.UserInput{}
	}

	return *g.controls
}
