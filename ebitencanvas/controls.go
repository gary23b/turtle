package ebitencanvas

import (
	"github.com/gary23b/turtle/turtlemodel"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func fillKeyStruct(input []ebiten.Key, out *turtlemodel.UserInput) {
	out.AnyPressed = len(input) > 0
	keys := &out.Keys
	for _, k := range input {
		switch k { //nolint:exhaustive
		case ebiten.KeyA:
			keys.A = true
		case ebiten.KeyB:
			keys.B = true
		case ebiten.KeyC:
			keys.C = true
		case ebiten.KeyD:
			keys.D = true
		case ebiten.KeyE:
			keys.E = true
		case ebiten.KeyF:
			keys.F = true
		case ebiten.KeyG:
			keys.G = true
		case ebiten.KeyH:
			keys.H = true
		case ebiten.KeyI:
			keys.I = true
		case ebiten.KeyJ:
			keys.J = true
		case ebiten.KeyK:
			keys.K = true
		case ebiten.KeyL:
			keys.L = true
		case ebiten.KeyM:
			keys.M = true
		case ebiten.KeyN:
			keys.N = true
		case ebiten.KeyO:
			keys.O = true
		case ebiten.KeyP:
			keys.P = true
		case ebiten.KeyQ:
			keys.Q = true
		case ebiten.KeyR:
			keys.R = true
		case ebiten.KeyS:
			keys.S = true
		case ebiten.KeyT:
			keys.T = true
		case ebiten.KeyU:
			keys.U = true
		case ebiten.KeyV:
			keys.V = true
		case ebiten.KeyW:
			keys.W = true
		case ebiten.KeyX:
			keys.X = true
		case ebiten.KeyY:
			keys.Y = true
		case ebiten.KeyZ:
			keys.Z = true

		case ebiten.KeyArrowLeft:
			keys.LeftArrow = true
		case ebiten.KeyArrowRight:
			keys.RightArrow = true
		case ebiten.KeyArrowUp:
			keys.UpArrow = true
		case ebiten.KeyArrowDown:
			keys.DownArrow = true

		case ebiten.KeyDigit0:
			keys.Number0 = true
		case ebiten.KeyDigit1:
			keys.Number1 = true
		case ebiten.KeyDigit2:
			keys.Number2 = true
		case ebiten.KeyDigit3:
			keys.Number3 = true
		case ebiten.KeyDigit4:
			keys.Number4 = true
		case ebiten.KeyDigit5:
			keys.Number5 = true
		case ebiten.KeyDigit6:
			keys.Number6 = true
		case ebiten.KeyDigit7:
			keys.Number7 = true
		case ebiten.KeyDigit8:
			keys.Number8 = true
		case ebiten.KeyDigit9:
			keys.Number9 = true

		case ebiten.KeyF1:
			keys.F1 = true
		case ebiten.KeyF2:
			keys.F2 = true
		case ebiten.KeyF3:
			keys.F3 = true
		case ebiten.KeyF4:
			keys.F4 = true
		case ebiten.KeyF5:
			keys.F5 = true
		case ebiten.KeyF6:
			keys.F6 = true
		case ebiten.KeyF7:
			keys.F7 = true
		case ebiten.KeyF8:
			keys.F8 = true
		case ebiten.KeyF9:
			keys.F9 = true
		case ebiten.KeyF10:
			keys.F10 = true
		case ebiten.KeyF11:
			keys.F11 = true
		case ebiten.KeyF12:
			keys.F12 = true

		case ebiten.KeySpace:
			keys.Space = true
		case ebiten.KeyBackspace:
			keys.Backspace = true
		case ebiten.KeyTab:
			keys.Tab = true
		case ebiten.KeyShiftRight:
			keys.RightShift = true
		case ebiten.KeyShiftLeft:
			keys.LeftShift = true
		case ebiten.KeyControlLeft:
			keys.LeftCtrl = true
		case ebiten.KeyControlRight:
			keys.RightCtrl = true
		case ebiten.KeyAltLeft:
			keys.LeftAlt = true
		case ebiten.KeyAltRight:
			keys.RightAlt = true
		case ebiten.KeyEscape:
			keys.Escape = true
		case ebiten.KeyEnter:
			keys.Enter = true
		case ebiten.KeyInsert:
			keys.Insert = true
		case ebiten.KeyDelete:
			keys.Delete = true
		case ebiten.KeyHome:
			keys.Home = true
		case ebiten.KeyEnd:
			keys.End = true
		case ebiten.KeyPageUp:
			keys.PageUp = true
		case ebiten.KeyPageDown:
			keys.PageDown = true

		case ebiten.KeyBackquote:
			keys.Backquote = true
		case ebiten.KeyMinus:
			keys.Minus = true
		case ebiten.KeyEqual:
			keys.Equal = true
		case ebiten.KeyComma:
			keys.Comma = true
		case ebiten.KeyPeriod:
			keys.Period = true
		case ebiten.KeySemicolon:
			keys.SemiColon = true
		case ebiten.KeyQuote:
			keys.Apostrophe = true
		case ebiten.KeySlash:
			keys.ForwardSlash = true
		case ebiten.KeyBackslash:
			keys.BackSlash = true
		case ebiten.KeyBracketLeft:
			keys.OpenSquareBracket = true
		case ebiten.KeyBracketRight:
			keys.CloseSquareBracket = true
		}
	}
}

// Struct to hold ebiten key state information. It is used every frame during the update.
type SavedControlState struct {
	keysDown        []ebiten.Key
	keysJustPressed []ebiten.Key
}

// Generate a new struct for pressed and just pressed. It should be read only to everyone else.
func (s *SavedControlState) GetUserInput(screenWidth, screenHeight int) (pressed, justPressed *turtlemodel.UserInput) {
	s.keysDown = inpututil.AppendPressedKeys(s.keysDown[:0])
	s.keysJustPressed = inpututil.AppendJustPressedKeys(s.keysJustPressed[:0])
	cursorX, cursorY := ebiten.CursorPosition()
	_, yScroll := ebiten.Wheel()
	Left := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	Center := ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle)
	Right := ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	LeftJp := inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	CenterJp := inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonMiddle)
	RightJp := inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight)

	// Now fill out the down struct
	pressed = &turtlemodel.UserInput{}
	fillKeyStruct(s.keysDown, pressed)
	pressed.Mouse.MouseX = cursorX - screenWidth/2
	pressed.Mouse.MouseY = -cursorY + screenHeight/2
	pressed.Mouse.MouseScroll = yScroll
	pressed.Mouse.Left = Left
	pressed.Mouse.Center = Center
	pressed.Mouse.Right = Right
	pressed.AnyPressed = pressed.AnyPressed || Left || Center || Right || yScroll != 0

	// Now fill out the justPressed struct
	justPressed = &turtlemodel.UserInput{}
	fillKeyStruct(s.keysJustPressed, justPressed)
	justPressed.Mouse.MouseX = cursorX - screenWidth/2
	justPressed.Mouse.MouseY = -cursorY + screenHeight/2
	justPressed.Mouse.MouseScroll = yScroll
	justPressed.Mouse.Left = LeftJp
	justPressed.Mouse.Center = CenterJp
	justPressed.Mouse.Right = RightJp
	justPressed.AnyPressed = justPressed.AnyPressed || LeftJp || CenterJp || RightJp || yScroll != 0

	return pressed, justPressed
}
