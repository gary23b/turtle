package models

type Keys struct {
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
	Delete     bool
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

type MouseButtons struct {
	Left   bool
	Right  bool
	Center bool
	// Back    bool
	// Forward bool
}

type UserInput struct {
	KeysDown        Keys
	KeysJustPressed Keys

	MouseDown        MouseButtons
	MouseJustPressed MouseButtons
	MouseX           int
	MouseY           int
	MouseScroll      float64
}
