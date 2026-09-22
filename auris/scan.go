package auris

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func line(pos fyne.Position, size fyne.Size, c fyne.Color, width float32) *canvas.Rectangle {
	r := canvas.NewRectangle(c)
	r.Move(pos)
	r.Resize(size)
	return r
}

// NewScanBracket frames content with four HUD targeting corners.
func NewScanBracket(child fyne.CanvasObject, size fyne.Size, length, stroke float32) fyne.CanvasObject {
	c := Gold
	w, h := size.Width, size.Height
	parts := []fyne.CanvasObject{
		child,
		line(fyne.NewPos(0,0),fyne.NewSize(length,stroke),c,stroke),
		line(fyne.NewPos(0,0),fyne.NewSize(stroke,length),c,stroke),
		line(fyne.NewPos(w-length,0),fyne.NewSize(length,stroke),c,stroke),
		line(fyne.NewPos(w-stroke,0),fyne.NewSize(stroke,length),c,stroke),
		line(fyne.NewPos(0,h-stroke),fyne.NewSize(length,stroke),c,stroke),
		line(fyne.NewPos(0,h-length),fyne.NewSize(stroke,length),c,stroke),
		line(fyne.NewPos(w-length,h-stroke),fyne.NewSize(length,stroke),c,stroke),
		line(fyne.NewPos(w-stroke,h-length),fyne.NewSize(stroke,length),c,stroke),
	}
	g := container.NewWithoutLayout(parts...)
	g.Resize(size)
	return g
}
