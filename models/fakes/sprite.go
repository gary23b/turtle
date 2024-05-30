package fakes

import (
	"image"

	"github.com/gary23b/turtle/models"
)

type spriteToDraw struct {
	Img     image.Image
	Width   int
	Height  int
	X       float64
	Y       float64
	Angle   float64
	Visible bool
	Scale   float64
}

var _ models.Sprite = &spriteToDraw{}

func NewSprite() *spriteToDraw {
	ret := &spriteToDraw{
		X:       0,
		Y:       0,
		Angle:   0,
		Visible: false,
		Scale:   1,
	}

	return ret
}

func (s *spriteToDraw) SetSpriteImage(in image.Image) {
	s.Img = in
	bounds := in.Bounds()
	s.Width = bounds.Max.X
	s.Height = bounds.Max.Y
	s.Scale = 1
}

func (s *spriteToDraw) SetSpriteImageTurtle() {
}

func (s *spriteToDraw) SetSpriteImageArrow() {
}

func (s *spriteToDraw) SetRotation(radianAngle float64) {
	s.Angle = radianAngle
}

func (s *spriteToDraw) SetPosition(cartX, cartY float64) {
	s.X = cartX
	s.Y = cartY
}

func (s *spriteToDraw) SetVisible(visible bool) {
	s.Visible = visible
}

func (s *spriteToDraw) SetScale(scale float64) {
	s.Scale = scale
}

func (s *spriteToDraw) Set(visible bool, cartX, cartY, radianAngle float64) {
	s.Visible = visible
	s.X = cartX
	s.Y = cartY
	s.Angle = radianAngle
}
