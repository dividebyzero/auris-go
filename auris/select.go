package auris

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Select keeps Fyne's popup mechanics while Auris owns the trigger surface.
// A fully custom HUD popup remains a later fidelity step.
type Select struct {
	widget.BaseWidget
	Options []string
	Selected, Placeholder string
	Disabled bool
	OnChanged func(string)
}

func NewSelect(options []string, selected string, changed func(string)) *Select {
	s:=&Select{Options:options,Selected:selected,Placeholder:"SELECT",OnChanged:changed}
	s.ExtendBaseWidget(s)
	return s
}
func (s *Select) SetSelected(value string) { s.Selected=value; s.Refresh() }
func (s *Select) SetDisabled(disabled bool) { s.Disabled=disabled; s.Refresh() }
func (s *Select) CreateRenderer() fyne.WidgetRenderer { return &selectRenderer{owner:s} }

type selectRenderer struct { owner *Select; objects []fyne.CanvasObject }
func (r *selectRenderer) MinSize() fyne.Size { return fyne.NewSize(220,42) }
func (r *selectRenderer) Layout(size fyne.Size) { r.rebuild(size) }
func (r *selectRenderer) Refresh() { r.rebuild(r.owner.Size()); canvas.Refresh(r.owner) }
func (r *selectRenderer) Objects() []fyne.CanvasObject { return r.objects }
func (r *selectRenderer) Destroy() {}
func (r *selectRenderer) rebuild(size fyne.Size) {
	s:=DarkScheme()
	selectWidget:=widget.NewSelect(r.owner.Options,func(v string){
		r.owner.Selected=v
		if r.owner.OnChanged != nil { r.owner.OnChanged(v) }
	})
	selectWidget.PlaceHolder=r.owner.Placeholder
	selectWidget.Selected=r.owner.Selected
	if r.owner.Disabled { selectWidget.Disable() }
	selectWidget.Resize(size)
	border:=s.BorderBright
	if r.owner.Selected!="" { border=s.PrimaryDim }
	if r.owner.Disabled { border=withColorAlpha(border,0x80) }
	bg:=NewChamfer(s.SurfaceInset,border,s.Bevel.MD).Object(size)
	stack:=container.NewWithoutLayout(bg,selectWidget); stack.Resize(size)
	r.objects=[]fyne.CanvasObject{stack}
}
