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

// This returns nil if there is no new data.
// This will throw away all but the newest set of data available. So this should be called faster that the game update rate (60Hz),
// otherwise sim.PressedUserInput() should be used instead.
func GetNewestJustPressedFromChan(justPressedChan chan *models.UserInput) *models.UserInput {
	var ret *models.UserInput

ChanExtractionLoop:
	for {
		select {
		case i := <-justPressedChan:
			ret = i
		default:
			// receiving from chan would block
			break ChanExtractionLoop
		}
	}
	return ret
}
