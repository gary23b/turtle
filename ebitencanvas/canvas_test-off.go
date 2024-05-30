package ebitencanvas

// Fails to start these tests inside GitHub Actions....

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	Water color.RGBA = color.RGBA{0x23, 0x89, 0xDA, 0xFF} // 2389DA
	Black color.RGBA = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	Red   color.RGBA = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
)

func TestEbitenTurtleCanvas(t *testing.T) {
	cmdChan := make(chan drawCmd, 10000)
	canvas := &ebitenTurtleCanvas{
		width:   8,
		height:  4,
		cmdChan: cmdChan,
	}

	intVar := canvas.GetWidth()
	require.Equal(t, 8, intVar)
	intVar = canvas.GetHeight()
	require.Equal(t, 4, intVar)

	canvas.SetPixel(0, 0, Black)
	require.NotEmpty(t, cmdChan)
	d := <-cmdChan
	require.Equal(t, drawCmd{x: 0, y: 0, c: Black}, d)

	canvas.SetCartesianPixel(1, 1, Water)
	require.NotEmpty(t, cmdChan)
	d = <-cmdChan
	require.Equal(t, drawCmd{x: 5, y: 1, c: Water}, d)

	canvas.ClearScreen(Red)
	require.NotEmpty(t, cmdChan)
	d = <-cmdChan
	require.Equal(t, drawCmd{c: Red, clearScreen: true}, d)
}
