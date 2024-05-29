package ebitencanvas

// Fails to start these tests inside GitHub Actions....

import (
	"testing"

	"github.com/gary23b/turtle/turtleutil"
	"github.com/stretchr/testify/require"
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

	canvas.SetPixel(0, 0, turtleutil.Black)
	require.NotEmpty(t, cmdChan)
	d := <-cmdChan
	require.Equal(t, drawCmd{x: 0, y: 0, c: turtleutil.Black}, d)

	canvas.SetCartesianPixel(1, 1, turtleutil.Water)
	require.NotEmpty(t, cmdChan)
	d = <-cmdChan
	require.Equal(t, drawCmd{x: 5, y: 1, c: turtleutil.Water}, d)

	canvas.FillScreen(turtleutil.Red)
	require.NotEmpty(t, cmdChan)
	d = <-cmdChan
	require.Equal(t, drawCmd{clearScreen: &turtleutil.Red}, d)
}
