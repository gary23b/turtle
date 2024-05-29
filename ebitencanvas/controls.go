package ebitencanvas

import (
	"github.com/gary23b/turtle/models"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func fillKeyStruct(input []ebiten.Key, out *models.Keys) {
	for _, k := range input {
		switch k { //nolint:exhaustive
		case ebiten.KeyA:
			out.A = true
		case ebiten.KeyB:
			out.B = true
		case ebiten.KeyC:
			out.C = true
		case ebiten.KeyD:
			out.D = true
		case ebiten.KeyE:
			out.E = true
		case ebiten.KeyF:
			out.F = true
		case ebiten.KeyG:
			out.G = true
		case ebiten.KeyH:
			out.H = true
		case ebiten.KeyI:
			out.I = true
		case ebiten.KeyJ:
			out.J = true
		case ebiten.KeyK:
			out.K = true
		case ebiten.KeyL:
			out.L = true
		case ebiten.KeyM:
			out.M = true
		case ebiten.KeyN:
			out.N = true
		case ebiten.KeyO:
			out.O = true
		case ebiten.KeyP:
			out.P = true
		case ebiten.KeyQ:
			out.Q = true
		case ebiten.KeyR:
			out.R = true
		case ebiten.KeyS:
			out.S = true
		case ebiten.KeyT:
			out.T = true
		case ebiten.KeyU:
			out.U = true
		case ebiten.KeyV:
			out.V = true
		case ebiten.KeyW:
			out.W = true
		case ebiten.KeyX:
			out.X = true
		case ebiten.KeyY:
			out.Y = true
		case ebiten.KeyZ:
			out.Z = true

		case ebiten.KeyArrowLeft:
			out.LeftArrow = true
		case ebiten.KeyArrowRight:
			out.RightArrow = true
		case ebiten.KeyArrowUp:
			out.UpArrow = true
		case ebiten.KeyArrowDown:
			out.DownArrow = true

		case ebiten.KeyDigit0:
			out.Number0 = true
		case ebiten.KeyDigit1:
			out.Number1 = true
		case ebiten.KeyDigit2:
			out.Number2 = true
		case ebiten.KeyDigit3:
			out.Number3 = true
		case ebiten.KeyDigit4:
			out.Number4 = true
		case ebiten.KeyDigit5:
			out.Number5 = true
		case ebiten.KeyDigit6:
			out.Number6 = true
		case ebiten.KeyDigit7:
			out.Number7 = true
		case ebiten.KeyDigit8:
			out.Number8 = true
		case ebiten.KeyDigit9:
			out.Number9 = true

		case ebiten.KeyF1:
			out.F1 = true
		case ebiten.KeyF2:
			out.F2 = true
		case ebiten.KeyF3:
			out.F3 = true
		case ebiten.KeyF4:
			out.F4 = true
		case ebiten.KeyF5:
			out.F5 = true
		case ebiten.KeyF6:
			out.F6 = true
		case ebiten.KeyF7:
			out.F7 = true
		case ebiten.KeyF8:
			out.F8 = true
		case ebiten.KeyF9:
			out.F9 = true
		case ebiten.KeyF10:
			out.F10 = true
		case ebiten.KeyF11:
			out.F11 = true
		case ebiten.KeyF12:
			out.F12 = true

		case ebiten.KeySpace:
			out.Space = true
		case ebiten.KeyBackspace:
			out.Backspace = true
		case ebiten.KeyTab:
			out.Tab = true
		case ebiten.KeyShiftRight:
			out.RightShift = true
		case ebiten.KeyShiftLeft:
			out.LeftShift = true
		case ebiten.KeyControlLeft:
			out.LeftCtrl = true
		case ebiten.KeyControlRight:
			out.RightCtrl = true
		case ebiten.KeyAltLeft:
			out.LeftAlt = true
		case ebiten.KeyAltRight:
			out.RightAlt = true
		case ebiten.KeyEnter:
			out.Enter = true
		case ebiten.KeyDelete:
			out.Delete = true
		case ebiten.KeyEscape:
			out.Escape = true

		case ebiten.KeyBackquote:
			out.Backquote = true
		case ebiten.KeyMinus:
			out.Minus = true
		case ebiten.KeyEqual:
			out.Equal = true
		case ebiten.KeyComma:
			out.Comma = true
		case ebiten.KeyPeriod:
			out.Period = true
		case ebiten.KeySemicolon:
			out.SemiColon = true
		case ebiten.KeyQuote:
			out.Apostrophe = true
		case ebiten.KeySlash:
			out.ForwardSlash = true
		case ebiten.KeyBackslash:
			out.BackSlash = true
		case ebiten.KeyBracketLeft:
			out.OpenSquareBracket = true
		case ebiten.KeyBracketRight:
			out.CloseSquareBracket = true
		}
	}
}

type SavedControlState struct {
	keysDown        []ebiten.Key
	keysJustPressed []ebiten.Key
}

func (s *SavedControlState) GetUserInput() *models.UserInput {
	controls := models.UserInput{}

	s.keysDown = inpututil.AppendPressedKeys(s.keysDown[:0])
	s.keysJustPressed = inpututil.AppendJustPressedKeys(s.keysJustPressed[:0])
	fillKeyStruct(s.keysDown, &controls.KeysDown)
	// fillKeyStruct(s.keysJustPressed, &controls.KeysJustPressed)

	cursorX, cursorY := ebiten.CursorPosition()
	_, yScroll := ebiten.Wheel()
	controls.MouseX = cursorX
	controls.MouseY = cursorY
	controls.MouseScroll = yScroll

	controls.MouseDown.Left = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	controls.MouseDown.Center = ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle)
	controls.MouseDown.Right = ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	// controls.MouseJustPressed.Left = inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	// controls.MouseJustPressed.Center = inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonMiddle)
	// controls.MouseJustPressed.Right = inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight)

	return &controls
}
