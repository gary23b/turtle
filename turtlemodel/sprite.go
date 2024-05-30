package turtlemodel

import "image"

type SpriteInfo struct {
	X       float64
	Y       float64
	Angle   float64 // radians
	Visible bool
	Img     image.Image
	Scale   float64
}

// A sprite is a image that can be place on the screen. The x/y location, rotation, size, and visibility are all controllable.
type Sprite interface {
	SetSpriteImage(image.Image)
	SetSpriteImageTurtle()
	SetSpriteImageArrow()
	SetRotation(radianAngle float64)
	SetPosition(cartX, cartY float64) // Cartesian (x,y). Center in the middle of the window
	SetVisible(visible bool)
	SetScale(scale float64)
	Set(visible bool, cartX, cartY, radianAngle float64) // Cartesian (x,y). Center in the middle of the window
	Get() SpriteInfo
}
