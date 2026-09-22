package auris

import (
	"image/color"

	"fyne.io/fyne/v2"
)

func newSlant(size fyne.Size, slant float32, fill, stroke color.Color) fyne.CanvasObject {
	if slant < 0 {
		slant = 0
	}
	if slant > size.Width/2 {
		slant = size.Width / 2
	}
	points := []fyne.Position{
		{X: slant, Y: 0},
		{X: size.Width, Y: 0},
		{X: size.Width - slant, Y: size.Height},
		{X: 0, Y: size.Height},
	}
	return polygon(points, size, fill, stroke, 1)
}
