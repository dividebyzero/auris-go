package auris

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// polygon draws an arbitrary polygon using a solid fill raster plus canvas lines.
// Fyne's canvas.Polygon is a regular-polygon primitive and cannot represent
// chamfers or slanted quadrilaterals defined by arbitrary points.
func polygon(points []fyne.Position, size fyne.Size, fill, stroke color.Color, strokeWidth float32) fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, len(points)+1)
	if fill != nil {
		raster := canvas.NewRasterWithPixels(func(x, y, w, h int) color.Color {
			if pointInPolygon(float32(x)+0.5, float32(y)+0.5, points) {
				return fill
			}
			return color.Transparent
		})
		raster.Resize(size)
		objects = append(objects, raster)
	}
	if stroke != nil && strokeWidth > 0 && len(points) > 1 {
		for i := range points {
			line := canvas.NewLine(stroke)
			line.StrokeWidth = strokeWidth
			line.Position1 = points[i]
			line.Position2 = points[(i+1)%len(points)]
			objects = append(objects, line)
		}
	}
	group := container.NewWithoutLayout(objects...)
	group.Resize(size)
	return group
}

func pointInPolygon(x, y float32, points []fyne.Position) bool {
	inside := false
	j := len(points) - 1
	for i := range points {
		xi, yi := points[i].X, points[i].Y
		xj, yj := points[j].X, points[j].Y
		if (yi > y) != (yj > y) {
			crossX := (xj-xi)*(y-yi)/(yj-yi) + xi
			if x < crossX {
				inside = !inside
			}
		}
		j = i
	}
	return inside
}
