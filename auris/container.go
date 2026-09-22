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
	stack:=container.NewWithoutLayout(surface,content)
	stack.Resize(size)
	return stack
}

// NewDepthContainer adds the tight Auris glow convention without changing the
// caller's requested content size.
func NewDepthContainer(content fyne.CanvasObject,size fyne.Size,cut float32,fill,border color.Color,depth Depth) fyne.CanvasObject {
	spread:=depth.Spread
	glow:=NewGlowSurface(size,cut,fill,border,depth)
	content.Move(fyne.NewPos(spread,spread))
	content.Resize(size)
	stack:=container.NewWithoutLayout(glow,content)
	stack.Resize(fyne.NewSize(size.Width+spread*2,size.Height+spread*2))
	return stack
}

func NewInsetContainer(content fyne.CanvasObject, size fyne.Size) fyne.CanvasObject {
	s := DarkScheme()
	return NewContainer(content, size, s.Bevel.MD, s.SurfaceInset, s.BorderBright)
}

func NewPanelSurface(content fyne.CanvasObject, size fyne.Size, accent bool) fyne.CanvasObject {
	s := DarkScheme()
	if accent {
		return NewDepthContainer(content,size,s.Bevel.LG,s.SurfacePanel,s.PrimaryActive,DepthSubtle)
	}
	return NewContainer(content,size,s.Bevel.LG,s.SurfacePanel,s.BorderBright)
}

func NewRule(width float32) fyne.CanvasObject {
	r := canvas.NewRectangle(Border)
	r.SetMinSize(fyne.NewSize(width, 1))
	return r
}
