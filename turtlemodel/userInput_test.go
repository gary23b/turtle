package turtlemodel

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserInput_IsPressedByName(t *testing.T) {
	var s *UserInput

	require.False(t, s.IsPressedByName("a"))

	s = &UserInput{
		Keys:  KeysStruct{A: true},
		Mouse: MouseStruct{Center: true},
	}

	// the default case
	require.False(t, s.IsPressedByName("aa"))

	// check that case doesn't matter
	require.True(t, s.IsPressedByName("a"))
	require.True(t, s.IsPressedByName("A"))

	// Check the mouse
	require.False(t, s.IsPressedByName("leftmouse"))
	require.False(t, s.IsPressedByName("rightmouse"))
	require.True(t, s.IsPressedByName("centermouse"))

	require.True(t, s.IsPressedByName("a"))
	require.False(t, s.IsPressedByName("b"))
	require.False(t, s.IsPressedByName("c"))
	require.False(t, s.IsPressedByName("d"))
	require.False(t, s.IsPressedByName("e"))
	require.False(t, s.IsPressedByName("f"))
	require.False(t, s.IsPressedByName("g"))
	require.False(t, s.IsPressedByName("h"))
	require.False(t, s.IsPressedByName("i"))
	require.False(t, s.IsPressedByName("j"))
	require.False(t, s.IsPressedByName("k"))
	require.False(t, s.IsPressedByName("l"))
	require.False(t, s.IsPressedByName("m"))
	require.False(t, s.IsPressedByName("n"))
	require.False(t, s.IsPressedByName("o"))
	require.False(t, s.IsPressedByName("p"))
	require.False(t, s.IsPressedByName("q"))
	require.False(t, s.IsPressedByName("r"))
	require.False(t, s.IsPressedByName("s"))
	require.False(t, s.IsPressedByName("t"))
	require.False(t, s.IsPressedByName("u"))
	require.False(t, s.IsPressedByName("v"))
	require.False(t, s.IsPressedByName("w"))
	require.False(t, s.IsPressedByName("x"))
	require.False(t, s.IsPressedByName("y"))
	require.False(t, s.IsPressedByName("z"))
	s.Keys.Z = true
	require.True(t, s.IsPressedByName("z"))

	require.False(t, s.IsPressedByName("0"))
	require.False(t, s.IsPressedByName("1"))
	require.False(t, s.IsPressedByName("2"))
	require.False(t, s.IsPressedByName("3"))
	require.False(t, s.IsPressedByName("4"))
	require.False(t, s.IsPressedByName("5"))
	require.False(t, s.IsPressedByName("6"))
	require.False(t, s.IsPressedByName("7"))
	require.False(t, s.IsPressedByName("8"))
	require.False(t, s.IsPressedByName("9"))
	s.Keys.Number9 = true
	require.True(t, s.IsPressedByName("9"))

	require.False(t, s.IsPressedByName(" "))
	require.False(t, s.IsPressedByName("f1"))
	require.False(t, s.IsPressedByName("f2"))
	require.False(t, s.IsPressedByName("f3"))
	require.False(t, s.IsPressedByName("f4"))
	require.False(t, s.IsPressedByName("f5"))
	require.False(t, s.IsPressedByName("f6"))
	require.False(t, s.IsPressedByName("f7"))
	require.False(t, s.IsPressedByName("f8"))
	require.False(t, s.IsPressedByName("f9"))
	require.False(t, s.IsPressedByName("f10"))
	require.False(t, s.IsPressedByName("f11"))
	require.False(t, s.IsPressedByName("f12"))

	require.False(t, s.IsPressedByName("arrowdown"))
	require.False(t, s.IsPressedByName("arrowleft"))
	require.False(t, s.IsPressedByName("arrowright"))
	require.False(t, s.IsPressedByName("arrowup"))
	require.False(t, s.IsPressedByName("backspace"))
	require.False(t, s.IsPressedByName("enter"))

	require.False(t, s.IsPressedByName("alt"))
	require.False(t, s.IsPressedByName("altleft"))
	require.False(t, s.IsPressedByName("altright"))

	require.False(t, s.IsPressedByName("control"))
	require.False(t, s.IsPressedByName("controlleft"))
	require.False(t, s.IsPressedByName("controlright"))

	require.False(t, s.IsPressedByName("shift"))
	require.False(t, s.IsPressedByName("shiftleft"))
	require.False(t, s.IsPressedByName("shiftright"))

	require.False(t, s.IsPressedByName("tab"))
	require.False(t, s.IsPressedByName("escape"))
	require.False(t, s.IsPressedByName("insert"))
	require.False(t, s.IsPressedByName("delete"))
	require.False(t, s.IsPressedByName("home"))
	require.False(t, s.IsPressedByName("end"))
	require.False(t, s.IsPressedByName("pageup"))
	require.False(t, s.IsPressedByName("pagedown"))

	require.False(t, s.IsPressedByName("backquote"))
	require.False(t, s.IsPressedByName("minus"))
	require.False(t, s.IsPressedByName("equal"))
	require.False(t, s.IsPressedByName("leftbracket"))
	require.False(t, s.IsPressedByName("rightbracket"))
	require.False(t, s.IsPressedByName("backslash"))
	require.False(t, s.IsPressedByName("semicolon"))
	require.False(t, s.IsPressedByName("apostrophe"))
	require.False(t, s.IsPressedByName("comma"))
	require.False(t, s.IsPressedByName("period"))
	require.False(t, s.IsPressedByName("forwardslash"))

	s.Keys.ForwardSlash = true
	require.True(t, s.IsPressedByName("forwardslash"))
}
