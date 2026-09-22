package auris

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func newSlant(size fyne.Size, slant float32, fill, stroke color.Color) *canvas.Polygon {
	if slant < 0 { slant = 0 }
	if slant > size.Width/2 { slant = size.Width/2 }
	p := canvas.NewPolygon([]fyne.Position{
		{X: slant, Y: 0}, {X: size.Width, Y: 0},
		{X: size.Width-slant, Y: size.Height}, {X: 0, Y: size.Height},
	})
	p.FillColor = fill
	p.StrokeColor = stroke
	p.StrokeWidth = 1
	p.Resize(size)
	return p
}
