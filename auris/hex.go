package auris

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func hexagon(center fyne.Position, radius float32, stroke color.Color) fyne.CanvasObject {
	points := make([]fyne.Position, 6)
	for i := range points {
		a := math.Pi/3*float64(i) - math.Pi/6
		points[i] = fyne.NewPos(
			center.X+radius*float32(math.Cos(a)),
			center.Y+radius*float32(math.Sin(a)),
		)
	}
	size := fyne.NewSize(center.X+radius, center.Y+radius)
	return polygon(points, size, nil, stroke, 1)
}

// NewHexOrnament returns a decorative, non-interactive hex cluster.
func NewHexOrnament(size fyne.Size, radius float32) fyne.CanvasObject {
	c := withColorAlpha(BorderBright, 0x70)
	cx, cy := size.Width/2, size.Height/2
	dx := radius * 1.5
	dy := radius * 0.866
	objects := []fyne.CanvasObject{
		hexagon(fyne.NewPos(cx, cy), radius, c),
		hexagon(fyne.NewPos(cx+dx, cy+dy), radius, c),
		hexagon(fyne.NewPos(cx-dx, cy+dy), radius, c),
		hexagon(fyne.NewPos(cx+dx, cy-dy), radius, c),
		hexagon(fyne.NewPos(cx-dx, cy-dy), radius, c),
	}
	group := container.NewWithoutLayout(objects...)
	group.Resize(size)
	return group
}
