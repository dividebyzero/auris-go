package auris

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func line(pos fyne.Position, size fyne.Size, c color.Color) *canvas.Rectangle {
	r:=canvas.NewRectangle(c); r.Move(pos); r.Resize(size); return r
}

func scanParts(child fyne.CanvasObject,size fyne.Size,length,stroke float32,c color.Color) []fyne.CanvasObject {
	w,h:=size.Width,size.Height
	return []fyne.CanvasObject{
		child,
		line(fyne.NewPos(0,0),fyne.NewSize(length,stroke),c),
		line(fyne.NewPos(0,0),fyne.NewSize(stroke,length),c),
		line(fyne.NewPos(w-length,0),fyne.NewSize(length,stroke),c),
		line(fyne.NewPos(w-stroke,0),fyne.NewSize(stroke,length),c),
		line(fyne.NewPos(0,h-stroke),fyne.NewSize(length,stroke),c),
		line(fyne.NewPos(0,h-length),fyne.NewSize(stroke,length),c),
		line(fyne.NewPos(w-length,h-stroke),fyne.NewSize(length,stroke),c),
		line(fyne.NewPos(w-stroke,h-length),fyne.NewSize(stroke,length),c),
	}
}

func NewScanBracket(child fyne.CanvasObject,size fyne.Size,length,stroke float32) fyne.CanvasObject {
	return NewScanBracketState(child,size,length,stroke,1)
}

// NewScanBracketState renders a pulse frame. progress is 0..1; callers may
// drive it with Motion.Progress. This keeps timing outside presentation.
func NewScanBracketState(child fyne.CanvasObject,size fyne.Size,length,stroke,progress float32) fyne.CanvasObject {
	if progress<0 { progress=0 }; if progress>1 { progress=1 }
	alpha:=uint8(Lerp(0x59,0xff,progress))
	g:=container.NewWithoutLayout(scanParts(child,size,length,stroke,withColorAlpha(Gold,alpha))...)
	g.Resize(size)
	return g
}
