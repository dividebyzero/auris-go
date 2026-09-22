package auris

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type ActionText struct {
	widget.BaseWidget
	Text string
	Color color.Color
	OnTapped func()
}

func NewActionText(text string,c color.Color,tapped func()) *ActionText {
	a:=&ActionText{Text:text,Color:c,OnTapped:tapped}; a.ExtendBaseWidget(a); return a
}
func (a *ActionText) Tapped(*fyne.PointEvent) { if a.OnTapped!=nil { a.OnTapped() } }
func (a *ActionText) CreateRenderer() fyne.WidgetRenderer {
	t:=DataText(a.Text,14,a.Color)
	return widget.NewSimpleRenderer(t)
}
func (a *ActionText) MinSize() fyne.Size { return canvas.NewText(a.Text,a.Color).MinSize() }
