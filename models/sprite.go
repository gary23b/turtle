package models

import "image"

type Sprite interface {
	SetSpriteImage(image.Image)
	SetSpriteImageTurtle()
	SetSpriteImageArrow()
	SetRotation(radianAngle float64)
	SetPosition(cartX, cartY float64) // Cartesian (x,y). Center in the middle of the window
	SetVisible(visible bool)
	SetScale(scale float64)
	Set(visible bool, cartX, cartY, radianAngle float64) // Cartesian (x,y). Center in the middle of the window
}
