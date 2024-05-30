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

	ret.g = newGame(params.Width, params.Height, params.ShowFPS, cmdChan)
	go drawFunc(ret)
	ret.g.runGame()
}

func (s *ebitenTurtleCanvas) CreateNewSprite() models.Sprite {
	ret := NewSprite()
	s.g.addSprite(ret)
	return ret
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

func (s *ebitenTurtleCanvas) PressedUserInput() *models.UserInput {
	ret := s.g.PressedUserInput()
	return ret
}

func (s *ebitenTurtleCanvas) SubscribeToJustPressedUserInput() chan *models.UserInput {
	return s.g.justPressedBroker.Subscribe()
}

func (s *ebitenTurtleCanvas) UnSubscribeToJustPressedUserInput(in chan *models.UserInput) {
	s.g.justPressedBroker.Unsubscribe(in)
}

func (s *ebitenTurtleCanvas) Exit() {
	s.g.TellGameToExit()
}
