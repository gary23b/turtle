package main

import (
	"math"
	"time"

	"github.com/gary23b/turtle"
	"github.com/gary23b/turtle/turtlemodel"
)

// translated from the python tutorial found here: https://www.geeksforgeeks.org/y-fractal-tree-in-python-using-turtle/#
func main() {
	params := turtle.Params{Width: 600, Height: 400}
	turtle.Start(params, drawFunc)
}

const (
	LevelCount = 16
	StartSize  = 80
	TrunkWidth = 15
)

func drawFunc(window turtle.Window) {
	time.Sleep(time.Second * 1)

	// go func() {
	// 	t := window.NewTurtle()

	// 	t.ShapeAsArrow()
	// 	t.ShowTurtle()
	// 	t.DegreesMode()
	// 	t.PenDown()
	// 	t.Speed(turtlemodel.MaxSpeed)

	// 	// turning the turtle to face upwards
	// 	t.Left(90)
	// 	t.Size(TrunkWidth)

	// 	recursiveTree(t, StartSize, LevelCount)
	// }()

	go RecursiveTree2(nil, window, 0, -180, 90, StartSize, LevelCount)

	// go turtle.CreateGif(window, time.Millisecond*200, time.Millisecond*200, "./examples/fractaltree/fractaltree.gif", 75)
}

func RecursiveTree(t turtlemodel.Turtle, length float64, level int) {
	if level <= 0 {
		return
	}
	// the acute angle between
	// the base and branch of the Y
	angle := 30.0

	// splitting the rgb range for green
	// into equal intervals for each level
	// setting the color according
	// to the current level
	ratio := dualRate(LevelCount, 1, 0, 1, float64(level))
	c := turtle.LerpColor(turtle.SaddleBrown, turtle.Lime, ratio)
	t.Color(c)

	s := t.GetSize() * 3 / 4
	t.Size(s)

	// drawing the base
	t.Forward(length)

	t.Right(angle)

	// recursive call for
	// the right subtree
	RecursiveTree(t, 0.8*length, level-1)
	t.Size(s)
	t.Color(c)

	t.Left(2 * angle)

	// recursive call for
	// the left subtree
	RecursiveTree(t, 0.8*length, level-1)
	t.Size(s)
	t.Color(c)

	t.Right(angle)
	t.Forward(-length)
}

func dualRate(x1, x2, y1, y2, in float64) float64 {
	slope := (y2 - y1) / (x2 - x1)
	deltaX := in - x1

	ret := deltaX*slope + y1

	return ret
}

func RecursiveTree2(t turtlemodel.Turtle, window turtle.Window, x, y, angle, length float64, level int) {
	if level <= 0 {
		return
	}

	if t == nil {
		t = window.NewTurtle()
		t.ShapeAsArrow()
		t.ShapeScale(.25)
		t.DegreesMode()
		// t.Speed(turtlemodel.MaxSpeed)
		t.ShowTurtle()
	}

	t.PenUp()
	t.Angle(angle)
	t.GoTo(x, y)
	t.PenDown()

	width := TrunkWidth * math.Pow(.75, float64(LevelCount-level))
	t.Size(width)

	// splitting the rgb range for green
	// into equal intervals for each level
	// setting the color according
	// to the current level
	ratio := dualRate(LevelCount, 1, 0, 1, float64(level))
	c := turtle.LerpColor(turtle.SaddleBrown, turtle.Lime, ratio)
	t.Color(c)

	t.Forward(length)

	curX, curY := t.GetPos()
	curAngle := t.GetAngle()

	if level > LevelCount-9 {
		t.HideTurtle()
		go RecursiveTree2(nil, window, curX, curY, curAngle+30, 0.8*length, level-1)
		go RecursiveTree2(nil, window, curX, curY, curAngle-30, 0.8*length, level-1)
	} else {
		RecursiveTree2(t, window, curX, curY, curAngle+30, 0.8*length, level-1)
		RecursiveTree2(t, window, curX, curY, curAngle-30, 0.8*length, level-1)
	}

	if level == LevelCount-9 {
		t.HideTurtle()
	}
}
