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
	Selected, Disabled bool
	Focused bool
	OnSelected func()
}

func NewRadio(label string, selected bool, changed func()) *Radio {
	r:=&Radio{Label:label,Selected:selected,OnSelected:changed}
	r.ExtendBaseWidget(r)
	return r
}
func (r *Radio) SetSelected(selected bool) { r.Selected=selected; r.Refresh() }
func (r *Radio) SetDisabled(disabled bool) { r.Disabled=disabled; r.Refresh() }
func (r *Radio) FocusGained() { r.Focused=true; r.Refresh() }
func (r *Radio) FocusLost() { r.Focused=false; r.Refresh() }
func (r *Radio) TypedRune(ch rune) { if ch == ' ' { r.Tapped(nil) } }
func (r *Radio) TypedKey(e *fyne.KeyEvent) { if e.Name == fyne.KeyReturn || e.Name == fyne.KeyEnter { r.Tapped(nil) } }
func (r *Radio) Tapped(*fyne.PointEvent) {
	if r.Disabled { return }
	r.Selected=true
	r.Refresh()
	if r.OnSelected != nil { r.OnSelected() }
}
func (r *Radio) CreateRenderer() fyne.WidgetRenderer { return &radioRenderer{owner:r} }

type radioRenderer struct { owner *Radio; objects []fyne.CanvasObject }
func (r *radioRenderer) MinSize() fyne.Size { return fyne.NewSize(140,28) }
func (r *radioRenderer) Layout(size fyne.Size) { r.rebuild(size) }
func (r *radioRenderer) Refresh() { r.rebuild(r.owner.Size()); canvas.Refresh(r.owner) }
func (r *radioRenderer) Objects() []fyne.CanvasObject { return r.objects }
func (r *radioRenderer) Destroy() {}
func (r *radioRenderer) rebuild(size fyne.Size) {
	s:=CurrentScheme()
	alpha:=uint8(0xff); if r.owner.Disabled { alpha=0x80 }
	border:=s.BorderBright; if r.owner.Selected || r.owner.Focused { border=s.PrimaryActive }
	box:=NewChamfer(withColorAlpha(s.SurfaceInset,alpha),withColorAlpha(border,alpha),s.Bevel.XS).Object(fyne.NewSize(18,18))
	indicator:=container.NewWithoutLayout(box); indicator.Resize(fyne.NewSize(24,24))
	if r.owner.Selected {
		pip:=NewChamfer(withColorAlpha(s.PrimaryActive,alpha),withColorAlpha(s.PrimaryActive,alpha),2).Object(fyne.NewSize(8,8))
		pip.Move(fyne.NewPos(5,5)); indicator.Add(pip)
	}
	label:=BodyText(r.owner.Label,13,withColorAlpha(s.TextBright,alpha))
	row:=container.NewHBox(indicator,label); row.Resize(size)
	r.objects=[]fyne.CanvasObject{row}
}
