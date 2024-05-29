// Copyright 2019 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	whiteImage = ebiten.NewImage(3, 3)

	// whiteSubImage is an internal sub image of whiteImage.
	// Use whiteSubImage at DrawTriangles instead of whiteImage in order to avoid bleeding edges.
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func init() {
	whiteImage.Fill(color.White)
}

const (
	screenWidth  = 640
	screenHeight = 480
)

func drawEbitenText(screen *ebiten.Image, x, y int) {
	aa := true
	line := true

	var path vector.Path

	// E
	path.MoveTo(20, 20)
	path.LineTo(20, 70)
	path.LineTo(70, 70)
	path.LineTo(70, 60)
	path.LineTo(30, 60)
	path.LineTo(30, 50)
	path.LineTo(70, 50)
	path.LineTo(70, 40)
	path.LineTo(30, 40)
	path.LineTo(30, 30)
	path.LineTo(70, 30)
	path.LineTo(70, 20)
	path.Close()

	var vs []ebiten.Vertex
	var is []uint16
	if line {
		op := &vector.StrokeOptions{}
		op.Width = 5
		op.LineJoin = vector.LineJoinRound
		vs, is = path.AppendVerticesAndIndicesForStroke(nil, nil, op)
	} else {
		vs, is = path.AppendVerticesAndIndicesForFilling(nil, nil)
	}

	for i := range vs {
		vs[i].DstX = (vs[i].DstX + float32(x))
		vs[i].DstY = (vs[i].DstY + float32(y))
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = 0xdb / float32(0xff)
		vs[i].ColorG = 0x56 / float32(0xff)
		vs[i].ColorB = 0x20 / float32(0xff)
		vs[i].ColorA = 1
	}

	op := &ebiten.DrawTrianglesOptions{}
	op.AntiAlias = aa
	if !line {
		op.FillRule = ebiten.EvenOdd
	}
	screen.DrawTriangles(vs, is, whiteSubImage, op)
}

// //////////////////////////////
// //////////////////////////////

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	dst := screen

	dst.Fill(color.RGBA{0xe0, 0xe0, 0xe0, 0xff})
	drawEbitenText(dst, 0, 50)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
