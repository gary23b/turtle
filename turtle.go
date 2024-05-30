package turtle

import (
	"github.com/gary23b/turtle/ebitencanvas"
	"github.com/gary23b/turtle/models"
	"github.com/gary23b/turtle/pen"
)

type Window interface {
	GetCanvas() models.Canvas
	NewTurtle() models.Turtle
}

type window struct {
	can models.Canvas
}

var _ Window = &window{}

func (s *window) GetCanvas() models.Canvas {
	return s.can
}

func (s *window) NewTurtle() models.Turtle {
	return pen.NewPen(s.can)
}

///////////////////////////////////////////////////
///////////////////////////////////////////////////

type Params struct {
	Width   int
	Height  int
	ShowFPS bool
}

// Wrap the starting of ebitencanvas inside something that implements the Window interface
// so that most of the time a user will only need one import statement from this repo to make a turtle graphic.
// But the actual game, drawing, and sprite implementations can still be separated nicely into packages.
func Start(params Params, drawFunc func(Window)) {
	canvasParams := ebitencanvas.CanvasParams{
		Width:   params.Width,
		Height:  params.Height,
		ShowFPS: params.ShowFPS,
	}

	// Create a callback that translates the models.Canvas into a Window
	initCallback := func(can models.Canvas) {
		drawFunc(&window{can: can})
	}
	ebitencanvas.StartEbitenTurtleCanvas(canvasParams, initCallback)
}
