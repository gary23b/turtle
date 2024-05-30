# Turtle

[![Go Reference](https://pkg.go.dev/badge/github.com/gary23b/turtle.svg)](https://pkg.go.dev/github.com/gary23b/turtle)
[![Go CI](https://github.com/gary23b/turtle/actions/workflows/go.yml/badge.svg)](https://github.com/gary23b/turtle/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gary23b/turtle)](https://goreportcard.com/report/github.com/gary23b/turtle)
[![Coverage Badge](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/gary23b/fc54fb0b86a835ca3f907efc85a1d61c/raw/gary23b_turtle_main.json)](https://github.com/gary23b/turtle/actions)

## A Turtle Graphics System for Golang

Based on the python turtle, this Go package provides an environment to learn Go programming while getting instant feedback. The screen is updated in real time using the [Ebitengine](https://ebitengine.org/).

Using Turtle is super easy and beginner-friendly. This is a good tool to use for an introduction to programming.

There is very little boilerplate needed to do simple things. Below is an example to draw a triangle.

```go
func drawTriangle(window turtle.Window) {
	t := window.NewTurtle()
	t.ShowTurtle()
	t.PenDown()

	t.Forward(100)
	t.Left(120)
	t.Forward(100)
	t.Left(120)
	t.Forward(100)
}
```

### Angle Modes

Three angle modes are supported:

- `DegreesMode`: (default). 0 is in the x direction, positive counterclockwise
- `RadiansMode`: 0 is in the x direction, positive counterclockwise
- `CompassMode`: 0 is North, or up, and positive clockwise
  - Can be used for solving orienteering problems

### Turtle Sprite

The turtle sprite can be hidden(default) or show. Its shape can be changed to an arrow or to a user provided `image.Image`. The scale of the sprite can also be adjusted.

## Install

Go 1.22 or later is required.<br>
Ebitengine is the main dependency. [Check here for the system specific instructions](https://ebitengine.org/en/documents/install.html).

## Examples

### 5 Turtles At Once

```bash
go run github.com/gary23b/turtle/examples/fiveturtles@latest
```

![Golang Turtle drawing with 5 individually controlled turtles](https://github.com/gary23b/turtle/blob/main/examples/fiveturtles/turtlebasic.gif)

### Go Gopher

Converted from the python script seen in [this youtube video](https://www.youtube.com/watch?v=d8A1jqOGzNE).

```bash
go run github.com/gary23b/turtle/examples/gogopher@latest
```

![Golang Turtle drawing the go gopher](https://github.com/gary23b/turtle/blob/main/examples/gogopher/GoGopher.gif)

### Recursive Tree

![Golang Turtle drawing a recursive tree](https://github.com/gary23b/turtle/blob/main/examples/fractaltree/fractaltree.gif)

## Basic Example Program

```bash
go run github.com/gary23b/turtle/examples/spiralweb@latest
```

```go
package main

import (
	"image/color"

	"github.com/gary23b/turtle"
)

func main() {
	params := turtle.Params{Width: 500, Height: 500}
	turtle.Start(params, drawFunc)
}

// drawFunc is started as a goroutine.
func drawFunc(window turtle.Window) {
	window.GetCanvas().ClearScreen(turtle.Black)
	t := window.NewTurtle()
	t.ShowTurtle()
	t.Speed(1000)
	t.PenDown()

	colors := []color.Color{turtle.Red, turtle.Yellow, turtle.Green, turtle.Purple, turtle.Blue, turtle.Orange}
	for x := 1; x < 200; x++ {
		t.Color(colors[x%6])        // setting color
		t.Size(float64(x)/75.0 + 1) // setting width
		t.Forward(float64(x))       // moving forward
		t.Left(59)                  // Turn left
	}
}
```

![Golang Turtle drawing spiral web](https://github.com/gary23b/turtle/blob/main/examples/spiralweb/spiralweb.gif)

## Turtle Controls

When a new turtle is created, it is given a Canvas interface. The turtle itself fulfills the Turtle interface seen below.
Which means each turtle created can perform this exact list of actions.

```go
type Turtle interface {
	Forward(distance float64)
	F(distance float64) // Forward alias
	Backward(distance float64)
	B(distance float64) // Backward alias
	PanRightward(distance float64)
	PanR(distance float64) // PanRightward alias
	PanLeftward(distance float64)
	PanL(distance float64) // PanLeftward alias

	GoTo(x, y float64)      // Cartesian (x,y). Center in the middle of the window
	GetPos() (x, y float64) // Cartesian (x,y). Center in the middle of the window

	Left(angle float64)
	L(angle float64) // Turn Left alias
	Right(angle float64)
	R(angle float64) // Turn Right alias
	Angle(angle float64)
	GetAngle() float64
	PointToward(x, y float64)

	DegreesMode() // Default is degrees mode.
	RadiansMode()
	CompassMode() // Make it so North is 0 degrees, East is 90...
	GetAngleMode() AngleMode

	Speed(PixelsPerSecond float64)
	GetSpeed() float64

	PenUp()
	PU()  // Pen Up alias
	Off() // Pen Up alias
	PenDown()
	PD() // Pen Down alias
	On() // Pen Down alias
	IsPenDown() bool
	Color(c color.Color)
	GetColor() color.Color
	Size(size float64)
	GetSize() float64
	Dot(size float64)
	Fill(c color.Color)

	// Draw a circle with given radius. The center is radius units left of the turtle; angleAmountToDraw determines
	// which part of the circle is drawn. If angleAmountToDraw is not a full circle, one endpoint of the arc is
	// the current pen position. Draw the arc in counterclockwise direction if radius is positive,
	// otherwise in clockwise direction. Finally the direction of the turtle is changed by the amount of angleAmountToDraw.
	//
	// As the circle is approximated by an inscribed regular polygon, steps determines the number of steps to use.
	// May be used to draw regular polygons.
	Circle(radius, angleAmountToDraw float64, steps int)

	ShowTurtle()
	HideTurtle()    // Default
	ShapeAsTurtle() // Default
	ShapeAsArrow()
	ShapeAsImage(in image.Image)
	ShapeScale(scale float64) // Default = 0.35

	Clone() Turtle // Create a clone of the turtle
}
```

## Canvas

The `pen` that implements the `Turtle` interface is able to implement all of required methods using the just the `Canvas` interface.

```go
type Canvas interface {
	CreateNewSprite() Sprite
	SetCartesianPixel(x, y int, c color.Color) // Cartesian (x,y). Center in the middle of the window
	SetPixel(x, y int, c color.Color)          // Computer graphics (x,y). So x=0, y=0 is the top-left of the window, positive down-right.
	Fill(x, y int, c color.Color)              // Cartesian (x,y). Center in the middle of the window
	ClearScreen(c color.Color)
	GetScreenshot() image.Image

	GetWidth() int
	GetHeight() int

	PressedUserInput() *UserInput
	SubscribeToJustPressedUserInput() chan *UserInput
	UnSubscribeToJustPressedUserInput(in chan *UserInput)

	Exit()
}
```

A user can also create their own `Pen` if desired. Or they can interact with the canvas directly. This can be seen in the `randomblue` example.
