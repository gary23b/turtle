package ebitencanvas

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"os"
	"sync"

	"github.com/gary23b/turtle/models"
	"github.com/hajimehoshi/ebiten/v2"
)

func LoadSpriteFile(path string) (image.Image, error) {
	spriteFileData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read sprite file: %s, %w", path, err)
	}
	img, _, err := image.Decode(bytes.NewReader(spriteFileData))
	if err != nil {
		return nil, fmt.Errorf("Failed to decode image data: %s, %w", path, err)
	}
	return img, nil
}

type spriteToDraw struct {
	m           *sync.Mutex
	spriteImage *ebiten.Image
	width       int
	height      int
	x           float64
	y           float64
	angle       float64
	visible     bool
	scale       float64
}

var _ models.Sprite = &spriteToDraw{}

func NewSprite() *spriteToDraw {
	ret := &spriteToDraw{
		m:       &sync.Mutex{},
		x:       0,
		y:       0,
		angle:   0,
		visible: false,
		scale:   1,
	}

	ret.SetSpriteImageTurtle()
	return ret
}

func (s *spriteToDraw) SetSpriteImage(in image.Image) {
	s.m.Lock()
	defer s.m.Unlock()
	s.spriteImage = ebiten.NewImageFromImage(in)
	bounds := in.Bounds()
	s.width = bounds.Max.X
	s.height = bounds.Max.Y
	s.scale = 1
}

func (s *spriteToDraw) SetSpriteImageTurtle() {
	fileData, err := base64.StdEncoding.DecodeString(turtleImage)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(bytes.NewReader(fileData))
	if err != nil {
		panic(err)
	}
	s.SetSpriteImage(img)
	s.SetScale(.35)
}

func (s *spriteToDraw) SetSpriteImageArrow() {
	fileData, err := base64.StdEncoding.DecodeString(arrowImage)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(bytes.NewReader(fileData))
	if err != nil {
		panic(err)
	}
	s.SetSpriteImage(img)
	s.SetScale(.5)
}

func (s *spriteToDraw) SetRotation(radianAngle float64) {
	s.m.Lock()
	defer s.m.Unlock()
	s.angle = radianAngle
}

func (s *spriteToDraw) SetPosition(cartX, cartY float64) {
	s.m.Lock()
	defer s.m.Unlock()
	s.x = cartX
	s.y = cartY
}

func (s *spriteToDraw) SetVisible(visible bool) {
	s.m.Lock()
	defer s.m.Unlock()
	s.visible = visible
}

func (s *spriteToDraw) SetScale(scale float64) {
	s.m.Lock()
	defer s.m.Unlock()
	s.scale = scale
}

func (s *spriteToDraw) Set(visible bool, cartX, cartY, radianAngle float64) {
	s.m.Lock()
	defer s.m.Unlock()
	s.visible = visible
	s.x = cartX
	s.y = cartY
	s.angle = radianAngle
}

// func (s *spriteToDraw) get() (visible bool, cartX, cartY, radianAngle float64) {
// 	s.m.Lock()
// 	defer s.m.Unlock()
// 	return s.visible, s.x, s.y, s.angle
// }
