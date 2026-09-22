package auris

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type Depth struct {
	Color color.Color
	Spread float32
	Alpha uint8
}

var (
	DepthSubtle = Depth{Color: Amber, Spread: 3, Alpha: 0x32}
	DepthActive = Depth{Color: Gold, Spread: 4, Alpha: 0x52}
	DepthDanger = Depth{Color: DangerBright, Spread: 4, Alpha: 0x48}
)

// NewGlowSurface approximates Auris' tight depth glow using nested geometry.
// It deliberately keeps spread small: glow is depth, not a large drop shadow.
func NewGlowSurface(size fyne.Size, cut float32, fill, border color.Color, depth Depth) fyne.CanvasObject {
	spread:=depth.Spread
	outerSize:=fyne.NewSize(size.Width+spread*2,size.Height+spread*2)
	glow:=NewChamfer(withColorAlpha(depth.Color,depth.Alpha),withColorAlpha(depth.Color,0),cut+spread).Object(outerSize)
	surface:=NewChamfer(fill,border,cut).Object(size)
	surface.Move(fyne.NewPos(spread,spread))
	stack:=container.NewWithoutLayout(glow,surface)
	stack.Resize(outerSize)
	return stack
}

func NewGlowText(text string, c color.Color, size float32) fyne.CanvasObject {
	// Fyne canvas.Text has no glyph-shadow primitive. Offset duplicates keep
	// the halo close to the glyphs instead of creating a rectangular box glow.
	shadow:=withColorAlpha(c,0x35)
	objs:=[]fyne.CanvasObject{}
	for _,p:=range []fyne.Position{{X:-1},{X:1},{Y:-1},{Y:1}} {
		t:=canvas.NewText(text,shadow); t.TextSize=size; t.Move(p); objs=append(objs,t)
	}
	main:=canvas.NewText(text,c); main.TextSize=size; objs=append(objs,main)
	return container.NewWithoutLayout(objs...)
}
