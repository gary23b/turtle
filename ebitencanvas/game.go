package ebitencanvas

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/gary23b/turtle/models"
	"github.com/gary23b/turtle/tools"

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
	spritesChan  chan *spriteToDraw
	img          *image.RGBA
	sprites      []*spriteToDraw
	exitFlag     bool

	justPressedBroker   *tools.Broker[*models.UserInput]
	controlState        SavedControlState
	controlsPressed     *models.UserInput
	controlsJustPressed *models.UserInput
}

func newGame(width, height int, showFPS bool, commands chan drawCmd) *game {
	g := &game{
		screenWidth:       width,
		screenHeight:      height,
		showFPS:           showFPS,
		commands:          commands,
		spritesChan:       make(chan *spriteToDraw, 100),
		img:               image.NewRGBA(image.Rect(0, 0, width, height)),
		sprites:           []*spriteToDraw{},
		justPressedBroker: tools.NewBroker[*models.UserInput](),
	}

	white := color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
	clearImage(g.img, white)

	// ebiten.SetTPS(120)
	// ebiten.SetVsyncEnabled(false) // For some reason, on Windows, there is quite a bit of lag.
	// setting this to false clears it up, but also makes it run at 1000Hz...
	ebiten.SetWindowSize(g.screenWidth, g.screenHeight)
	ebiten.SetWindowTitle("Go Turtle Graphics")
	return g
}

func (g *game) addSprite(in *spriteToDraw) {
	g.spritesChan <- in
}

// This function will not return. It must be run on the main thread.
func (g *game) runGame() {
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func clearImage(i *image.RGBA, c color.Color) {
	xMin := i.Rect.Min.X
	yMin := i.Rect.Min.Y
	xMax := i.Rect.Max.X
	yMax := i.Rect.Max.Y
	c1, _ := color.RGBAModel.Convert(c).(color.RGBA)
	for y := yMin; y < yMax; y++ {
		for x := xMin; x < xMax; x++ {
			i.SetRGBA(x, y, c1)
		}
	}
}

func bucketFill(i *image.RGBA, x, y int, c color.Color) {
	colorMatches := func(a, b color.RGBA) bool {
		return a.R == b.R && a.G == b.G && a.B == b.B && a.A == b.A
	}

	type upNextStruct struct {
		x, y int
	}
	upNextStack := []upNextStruct{}

	xMin := i.Rect.Min.X
	yMin := i.Rect.Min.Y
	xMax := i.Rect.Max.X - 1
	yMax := i.Rect.Max.Y - 1

	c1, _ := color.RGBAModel.Convert(c).(color.RGBA)
	srcColor := i.RGBAAt(x, y)
	if colorMatches(srcColor, c1) {
		// The selected pixes is already the correct color
		return
	}
	upNextStack = append(upNextStack, upNextStruct{x: x, y: y})

	for len(upNextStack) > 0 {
		xy := upNextStack[len(upNextStack)-1]
		x := xy.x
		y := xy.y
		upNextStack = upNextStack[:len(upNextStack)-1]

		i.SetRGBA(x, y, c1)
		if x > xMin && colorMatches(i.RGBAAt(x-1, y), srcColor) {
			upNextStack = append(upNextStack, upNextStruct{x: x - 1, y: y})
		}
		if x < xMax && colorMatches(i.RGBAAt(x+1, y), srcColor) {
			upNextStack = append(upNextStack, upNextStruct{x: x + 1, y: y})
		}
		if y > yMin && colorMatches(i.RGBAAt(x, y-1), srcColor) {
			upNextStack = append(upNextStack, upNextStruct{x: x, y: y - 1})
		}
		if y < yMax && colorMatches(i.RGBAAt(x, y+1), srcColor) {
			upNextStack = append(upNextStack, upNextStruct{x: x, y: y + 1})
		}
	}
}

func (g *game) Update() error {
	if g.exitFlag {
		return ebiten.Termination
	}

	g.controlsPressed, g.controlsJustPressed = g.controlState.GetUserInput(g.screenWidth, g.screenHeight)
	if g.controlsJustPressed.AnyPressed {
		g.justPressedBroker.Publish(g.controlsJustPressed)
	}

	// Pull all the items out of the command channel until it is empty.
EatDrawCommandsLoop:
	for {
		select {
		case cmd := <-g.commands:
			switch {
			case cmd.fill:
				bucketFill(g.img, cmd.x, cmd.y, cmd.c)
			case cmd.clearScreen:
				clearImage(g.img, cmd.c)
			default:
				g.img.Set(cmd.x, cmd.y, cmd.c)
			}
		default:
			// receiving from g.commands would block
			break EatDrawCommandsLoop
		}
	}

EatNewSpritesLoop:
	for {
		select {
		case sprite := <-g.spritesChan:
			g.sprites = append(g.sprites, sprite)
		default:
			break EatNewSpritesLoop
		}
	}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.img.Pix)

	for _, sprite := range g.sprites {
		if !sprite.visible {
			continue
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(sprite.width)/2, -float64(sprite.height)/2)
		op.GeoM.Rotate(-sprite.angle) // This command rotates clockwise for some reason.
		op.GeoM.Scale(sprite.scale, sprite.scale)
		op.GeoM.Translate(float64(g.screenWidth/2), float64(g.screenHeight/2))
		op.GeoM.Translate(sprite.x, -sprite.y)

		screen.DrawImage(sprite.spriteImage, op)
	}

	if g.showFPS {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *game) PressedUserInput() *models.UserInput {
	if g == nil || g.controlsPressed == nil {
		return &models.UserInput{}
	}

	return g.controlsPressed
}

func (g *game) TellGameToExit() {
	g.exitFlag = true
}
