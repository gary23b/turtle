package ebitencanvas

import (
	"image/color"

	"github.com/gary23b/turtle/models"
)

type drawCmd struct {
	x           int
	y           int
	c           color.RGBA
	fill        bool // Bucket will starting from the x,y position
	clearScreen bool // When you want to set the entire screen to a color. Only fill in this when clearing screen.
}

type ebitenTurtleCanvas struct {
	width   int
	height  int
	g       *game
	cmdChan chan drawCmd
}

var _ models.Canvas = &ebitenTurtleCanvas{} // Force the linter to tell us if the interface is implemented

type CanvasParams struct {
	Width   int
	Height  int
	ShowFPS bool
}

// The drawFunc will be started as a go routine.
func StartEbitenTurtleCanvas(params CanvasParams, drawFunc func(models.Canvas)) {
	cmdChan := make(chan drawCmd, 10000)
	ret := &ebitenTurtleCanvas{
		width:   params.Width,
		height:  params.Height,
		cmdChan: cmdChan,
	}

	go drawFunc(ret)
	ret.g = newGame(params.Width, params.Height, params.ShowFPS, cmdChan)
	ret.g.runGame()
}

// 0,0 is the center of the screen. positive X is right, positive y is up.
func (s *ebitenTurtleCanvas) SetCartesianPixel(x, y int, c color.RGBA) {
	s.SetPixel(x+s.width/2, -y+s.height/2, c)
}

func (s *ebitenTurtleCanvas) SetPixel(x, y int, c color.RGBA) {
	newCmd := drawCmd{
		x: x,
		y: y,
		c: c,
	}
	s.cmdChan <- newCmd
}

// Bucket Fill, starting from the given cartesian pixel
func (s *ebitenTurtleCanvas) Fill(x, y int, c color.RGBA) {
	s.cmdChan <- drawCmd{
		x:    x + s.width/2,
		y:    -y + s.height/2,
		c:    c,
		fill: true,
	}
}

func (s *ebitenTurtleCanvas) ClearScreen(c color.RGBA) {
	s.cmdChan <- drawCmd{
		c:           c,
		clearScreen: true,
	}
}

func (s *ebitenTurtleCanvas) GetWidth() int {
	return s.width
}

func (s *ebitenTurtleCanvas) GetHeight() int {
	return s.height
}

func (s *ebitenTurtleCanvas) GetUserInput() models.UserInput {
	ret := s.g.getUserInput()

	// translate game space to turtle space
	ret.MouseX -= s.width / 2
	ret.MouseY = -ret.MouseY + s.height/2

	return ret
}
