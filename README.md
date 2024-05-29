# Turtle

[![Go Reference](https://pkg.go.dev/badge/github.com/gary23b/turtle.svg)](https://pkg.go.dev/github.com/gary23b/turtle)

## A Turtle Graphics System for Golang

Based on the python turtle, this Go package provides an environment to learn Go programming while getting instant feedback. The screen is updated in real time using the [Ebitengine](https://ebitengine.org/).

## Install

Ebitengine is the main dependency. [Check here the system specific instructions](https://ebitengine.org/en/documents/install.html).

## Example

```bash
go run github.com/gary23b/turtle/examples/turtlebasic@latest
```

![Example Picture](https://github.com/gary23b/turtle/blob/main/examples/turtlebasic/turtlebasic.png)

## Basic Example Program

```go
package main

import (
	"github.com/gary23b/turtle/ebitencanvas"
	"github.com/gary23b/turtle/models"
	"github.com/gary23b/turtle/turtle"
	"github.com/gary23b/turtle/turtleutil"
)

func main() {
	// Create the Ebitengine canvas.
	params := ebitencanvas.CanvasParams{Width: 1000, Height: 1000}
	ebitencanvas.StartEbitenTurtleCanvas(params, drawFunc)
}

// drawFunc is started as a goroutine.
func drawFunc(can models.Canvas) {
	var t models.Turtle = turtle.NewTurtle(can)
	t.SetDegreesMode()
	t.PenColor(turtleutil.White)
	t.SetSpeed(1000)
	t.PenDown()

	t.Left(45)
	t.Forward(100)
	t.Left(135)
	t.Forward(200 * 0.707)
	t.GoTo(0, 0)
	t.SetAngle(-15)
	for i := 0; i < 12; i++ {
		t.Forward(100)
		t.Right(30)
	}
}
```
