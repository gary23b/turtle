package turtlemodel

import "strings"

// All the currently supported keys.
type KeysStruct struct {
	LeftArrow  bool
	RightArrow bool
	UpArrow    bool
	DownArrow  bool

	A bool
	B bool
	C bool
	D bool
	E bool
	F bool
	G bool
	H bool
	I bool
	J bool
	K bool
	L bool
	M bool
	N bool
	O bool
	P bool
	Q bool
	R bool
	S bool
	T bool
	U bool
	V bool
	W bool
	X bool
	Y bool
	Z bool

	Number0 bool
	Number1 bool
	Number2 bool
	Number3 bool
	Number4 bool
	Number5 bool
	Number6 bool
	Number7 bool
	Number8 bool
	Number9 bool

	F1  bool
	F2  bool
	F3  bool
	F4  bool
	F5  bool
	F6  bool
	F7  bool
	F8  bool
	F9  bool
	F10 bool
	F11 bool
	F12 bool

	Space      bool
	Backspace  bool
	Tab        bool
	LeftShift  bool
	RightShift bool
	LeftCtrl   bool
	RightCtrl  bool
	LeftAlt    bool
	RightAlt   bool
	Enter      bool
	Insert     bool
	Delete     bool
	Home       bool
	End        bool
	PageUp     bool
	PageDown   bool
	Escape     bool

	Backquote          bool
	Minus              bool
	Equal              bool
	Comma              bool
	Period             bool
	SemiColon          bool
	Apostrophe         bool
	ForwardSlash       bool
	BackSlash          bool
	OpenSquareBracket  bool
	CloseSquareBracket bool
}

// All the currently supported mouse inputs.
type MouseStruct struct {
	Left   bool
	Right  bool
	Center bool
	// Back    bool
	// Forward bool

	MouseX      int
	MouseY      int
	MouseScroll float64
}

// Used for currently pressed and for the just pressed functionality.
type UserInput struct {
	AnyPressed bool

	Keys  KeysStruct
	Mouse MouseStruct
}

// Check if a key is pressed by its name.
// Some of the keys have multiple names, such as:: case "space", " ":
func (s *UserInput) IsPressedByName(name string) bool {
	if s == nil {
		return false
	}

	k := s.Keys
	m := s.Mouse

	switch strings.ToLower(name) {
	case "0":
		return k.Number0
	case "1":
		return k.Number1
	case "2":
		return k.Number2
	case "3":
		return k.Number3
	case "4":
		return k.Number4
	case "5":
		return k.Number5
	case "6":
		return k.Number6
	case "7":
		return k.Number6
	case "8":
		return k.Number8
	case "9":
		return k.Number9

	case "a":
		return k.A
	case "b":
		return k.B
	case "c":
		return k.C
	case "d":
		return k.D
	case "e":
		return k.E
	case "f":
		return k.F
	case "g":
		return k.G
	case "h":
		return k.H
	case "i":
		return k.I
	case "j":
		return k.J
	case "k":
		return k.K
	case "l":
		return k.L
	case "m":
		return k.M
	case "n":
		return k.N
	case "o":
		return k.O
	case "p":
		return k.P
	case "q":
		return k.Q
	case "r":
		return k.R
	case "s":
		return k.S
	case "t":
		return k.T
	case "u":
		return k.U
	case "v":
		return k.V
	case "w":
		return k.W
	case "x":
		return k.X
	case "y":
		return k.Y
	case "z":
		return k.Z
	case "space", " ":
		return k.Space

	case "f1":
		return k.F1
	case "f2":
		return k.F2
	case "f3":
		return k.F3
	case "f4":
		return k.F4
	case "f5":
		return k.F5
	case "f6":
		return k.F6
	case "f7":
		return k.F7
	case "f8":
		return k.F8
	case "f9":
		return k.F9
	case "f10":
		return k.F10
	case "f11":
		return k.F11
	case "f12":
		return k.F12

	case "arrowdown", "down":
		return k.DownArrow
	case "arrowleft", "left":
		return k.LeftArrow
	case "arrowright", "right":
		return k.RightArrow
	case "arrowup", "up":
		return k.UpArrow

	case "backspace":
		return k.Backspace
	case "enter", "\n", "\r", "\r\n":
		return k.Enter
	case "alt":
		return k.LeftAlt || k.RightAlt
	case "altleft":
		return k.LeftAlt
	case "altright":
		return k.RightAlt
	case "control", "ctrl":
		return k.LeftCtrl || k.RightCtrl
	case "controlleft", "ctrlleft":
		return k.LeftCtrl
	case "controlright", "ctrlright":
		return k.RightCtrl
	case "shift":
		return k.LeftShift || k.RightShift
	case "shiftleft":
		return k.LeftShift
	case "shiftright":
		return k.RightShift
	case "tab", "\t":
		return k.Tab
	case "escape":
		return k.Escape
	case "insert":
		return k.Insert
	case "delete":
		return k.Delete
	case "home":
		return k.Home
	case "end":
		return k.End
	case "pageup":
		return k.PageUp
	case "pagedown":
		return k.PageDown

	case "backquote", "`":
		return k.Backquote
	case "minus", "-":
		return k.Minus
	case "equal", "=":
		return k.Equal
	case "leftbracket", "[":
		return k.OpenSquareBracket
	case "rightbracket", "]":
		return k.CloseSquareBracket
	case "backslash", "\\":
		return k.BackSlash
	case "semicolon", ";":
		return k.SemiColon
	case "apostrophe", "singlequote", "quote", "'":
		return k.Apostrophe
	case "comma", ",":
		return k.Comma
	case "period", ".":
		return k.Period
	case "forwardslash", "slash", "/":
		return s.Keys.ForwardSlash

	// Mouse
	case "leftmouse":
		return m.Left
	case "rightmouse":
		return m.Right
	case "centermouse":
		return m.Center
	}

	return false
}
