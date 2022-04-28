package factoryColor

import "image/color"

func NewOrangeredHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xff, G: 0x45, B: 0x00, A: 0x80} // rgb(255, 69, 0)
}