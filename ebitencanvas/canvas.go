package ebitencanvas

import (
	"image"
	"image/color"

	"github.com/gary23b/turtle/models"
)

type drawCmd struct {
	x           int
	y           int
	c           color.Color
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

// Turtle Window creation parameters
type CanvasParams struct {
	Width   int  // pixels
	Height  int  // pixels
	ShowFPS bool // Show frame per second and other info in the top left corner.
}

// Start the Ebiten game.
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

// Add a new sprite to the game.
func (s *ebitenTurtleCanvas) CreateNewSprite() models.Sprite {
	ret := NewSprite()
	s.g.addSprite(ret)
	return ret
}

// Set a single pixel to the given color
// 0,0 is the center of the screen. positive X is right, positive y is up.
func (s *ebitenTurtleCanvas) SetCartesianPixel(x, y int, c color.Color) {
	s.SetPixel(x+s.width/2, -y+s.height/2, c)
}

// Set a single pixel to the given color
// 0,0 is the top/left of the screen. positive X is right, positive y is down.
func (s *ebitenTurtleCanvas) SetPixel(x, y int, c color.Color) {
	newCmd := drawCmd{
		x: x,
		y: y,
		c: c,
	}
	s.cmdChan <- newCmd
}

// Bucket Fill, starting from the given cartesian pixel
func (s *ebitenTurtleCanvas) Fill(x, y int, c color.Color) {
	s.cmdChan <- drawCmd{
		x:    x + s.width/2,
		y:    -y + s.height/2,
		c:    c,
		fill: true,
	}
}

// Set the entire screen to the given color
func (s *ebitenTurtleCanvas) ClearScreen(c color.Color) {
	s.cmdChan <- drawCmd{
		c:           c,
		clearScreen: true,
	}
}

// Get a screenshot of the window. It will be created on the next frame.
// This will block until the screenshot is ready.
func (s *ebitenTurtleCanvas) GetScreenshot() image.Image {
	screenshotChan := make(chan image.Image)
	s.g.getScreenshot(screenshotChan)
	screenshot := <-screenshotChan
	return screenshot
}

// Get the screen width in pixels.
func (s *ebitenTurtleCanvas) GetWidth() int {
	return s.width
}

// Get the screen height in pixels.
func (s *ebitenTurtleCanvas) GetHeight() int {
	return s.height
}

// Get the currently pressed user input. The returned point is read only. Do not edit it.
func (s *ebitenTurtleCanvas) PressedUserInput() *models.UserInput {
	ret := s.g.PressedUserInput()
	return ret
}

// Get a channel that receives updates every time a key-press event is detected.
func (s *ebitenTurtleCanvas) SubscribeToJustPressedUserInput() chan *models.UserInput {
	return s.g.justPressedBroker.Subscribe()
}

// After calling SubscribeToJustPressedUserInput, if the events are no longer wanted, call this function to stop the updates and close the channel.
func (s *ebitenTurtleCanvas) UnSubscribeToJustPressedUserInput(in chan *models.UserInput) {
	s.g.justPressedBroker.Unsubscribe(in)
}

// Request that Ebiten exits on the next frame.
func (s *ebitenTurtleCanvas) Exit() {
	s.g.TellGameToExit()
}
