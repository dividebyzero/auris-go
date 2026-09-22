package auris

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// Chamfer describes the signature Auris surface: top-left and bottom-right
// corners are cut at 45 degrees while the other corners remain square.
type Chamfer struct {
	Fill        color.Color
	Stroke      color.Color
	Cut         float32
	StrokeWidth float32
}

func NewChamfer(fill, stroke color.Color, cut float32) *Chamfer {
	return &Chamfer{Fill: fill, Stroke: stroke, Cut: cut, StrokeWidth: 1}
}

func (c *Chamfer) Object(size fyne.Size) fyne.CanvasObject {
	cut := c.Cut
	if cut < 0 {
		cut = 0
	}
	if cut > size.Width/2 {
		cut = size.Width / 2
	}
	if cut > size.Height/2 {
		cut = size.Height / 2
	}

	p := canvas.NewPolygon([]fyne.Position{
		{X: cut, Y: 0},
		{X: size.Width, Y: 0},
		{X: size.Width, Y: size.Height - cut},
		{X: size.Width - cut, Y: size.Height},
		{X: 0, Y: size.Height},
		{X: 0, Y: cut},
	})
	p.FillColor = c.Fill
	p.StrokeColor = c.Stroke
	p.StrokeWidth = c.StrokeWidth
	p.Resize(size)
	return p
}
