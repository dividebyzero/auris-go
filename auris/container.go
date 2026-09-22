package auris

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// NewContainer places content over an Auris chamfered surface.
func NewContainer(content fyne.CanvasObject, size fyne.Size, cut float32, fill, border color.Color) fyne.CanvasObject {
	surface := NewChamfer(fill, border, cut).Object(size)
	content.Resize(size)
	return container.NewStack(surface, content)
}

// NewInsetContainer creates the common inset HUD surface.
func NewInsetContainer(content fyne.CanvasObject, size fyne.Size) fyne.CanvasObject {
	s := DarkScheme()
	return NewContainer(content, size, s.Bevel.MD, s.SurfaceInset, s.BorderBright)
}

// NewPanelSurface creates the common panel HUD surface.
func NewPanelSurface(content fyne.CanvasObject, size fyne.Size, accent bool) fyne.CanvasObject {
	s := DarkScheme()
	border := s.BorderBright
	if accent {
		border = s.PrimaryActive
	}
	return NewContainer(content, size, s.Bevel.LG, s.SurfacePanel, border)
}

func NewRule(width float32) fyne.CanvasObject {
	r := canvas.NewRectangle(Border)
	r.SetMinSize(fyne.NewSize(width, 1))
	return r
}
