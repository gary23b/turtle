package models

import "image"

type Sprite interface {
	SetSpriteImage(image.Image)
	SetSpriteImageTurtle()
	SetSpriteImageArrow()
	SetRotation(radianAngle float64)
	SetPosition(cartX, cartY float64)
	SetVisible(visible bool)
	SetScale(scale float64)
	Set(visible bool, cartX, cartY, radianAngle float64)
}
