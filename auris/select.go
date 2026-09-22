package auris

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Select is the Fyne-native Auris single-value selector.
// Fyne owns popup/focus behavior; Auris owns palette and surrounding geometry.
type Select struct {
	widget.BaseWidget
	Options []string
	Selected string
	Placeholder string
	OnChanged func(string)
}

func NewSelect(options []string, selected string, changed func(string)) *Select {
	s := &Select{Options:options, Selected:selected, Placeholder:"SELECT", OnChanged:changed}
	s.ExtendBaseWidget(s)
	return s
}

func (s *Select) CreateRenderer() fyne.WidgetRenderer {
	return &selectRenderer{owner:s}
}

type selectRenderer struct { owner *Select; objects []fyne.CanvasObject }
func (r *selectRenderer) MinSize() fyne.Size { return fyne.NewSize(220,42) }
func (r *selectRenderer) Layout(size fyne.Size) { r.rebuild(size) }
func (r *selectRenderer) Refresh() { r.rebuild(r.owner.Size()); canvas.Refresh(r.owner) }
func (r *selectRenderer) Objects() []fyne.CanvasObject { return r.objects }
func (r *selectRenderer) Destroy() {}
func (r *selectRenderer) rebuild(size fyne.Size) {
	scheme := DarkScheme()
	selectWidget := widget.NewSelect(r.owner.Options, func(v string) {
		r.owner.Selected=v
		if r.owner.OnChanged != nil { r.owner.OnChanged(v) }
	})
	selectWidget.PlaceHolder=r.owner.Placeholder
	selectWidget.Selected=r.owner.Selected
	selectWidget.Resize(size)
	bg := NewChamfer(scheme.SurfaceInset,scheme.BorderBright,scheme.Bevel.MD).Object(size)
	stack := container.NewStack(bg,selectWidget)
	stack.Resize(size)
	r.objects=[]fyne.CanvasObject{stack}
}
