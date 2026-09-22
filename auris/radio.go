package auris

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Radio struct {
	widget.BaseWidget
	Label string
	Selected bool
	OnSelected func()
}

func NewRadio(label string, selected bool, changed func()) *Radio {
	r := &Radio{Label: label, Selected: selected, OnSelected: changed}
	r.ExtendBaseWidget(r)
	return r
}

func (r *Radio) Tapped(*fyne.PointEvent) {
	if r.OnSelected != nil { r.OnSelected() }
}

func (r *Radio) CreateRenderer() fyne.WidgetRenderer {
	return &radioRenderer{owner:r}
}

type radioRenderer struct { owner *Radio; objects []fyne.CanvasObject }
func (r *radioRenderer) MinSize() fyne.Size { return fyne.NewSize(140,28) }
func (r *radioRenderer) Layout(size fyne.Size) { r.rebuild(size) }
func (r *radioRenderer) Refresh() { r.rebuild(r.owner.Size()); canvas.Refresh(r.owner) }
func (r *radioRenderer) Objects() []fyne.CanvasObject { return r.objects }
func (r *radioRenderer) Destroy() {}
func (r *radioRenderer) rebuild(size fyne.Size) {
	s := DarkScheme()
	border := s.BorderBright
	if r.owner.Selected { border = s.PrimaryActive }
	box := NewChamfer(s.SurfaceInset,border,s.Bevel.XS).Object(fyne.NewSize(18,18))
	indicator := container.NewWithoutLayout(box)
	indicator.Resize(fyne.NewSize(24,24))
	if r.owner.Selected {
		pip := NewChamfer(s.PrimaryActive,s.PrimaryActive,2).Object(fyne.NewSize(8,8))
		pip.Move(fyne.NewPos(5,5))
		indicator.Add(pip)
	}
	label := BodyText(r.owner.Label, 13, s.TextBright);
row := container.NewHBox(indicator,label); row.Resize(size)
	r.objects=[]fyne.CanvasObject{row}
}
