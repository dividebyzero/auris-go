package auris

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Select is an Auris-owned single-value selector. Popup presentation is built
// from Auris primitives instead of wrapping Fyne's stock Select.
type Select struct {
	widget.BaseWidget
	Options []string
	Selected, Placeholder string
	Disabled, Focused, Open bool
	Highlighted int
	OnChanged func(string)
}

func NewSelect(options []string, selected string, changed func(string)) *Select {
	s:=&Select{Options:options,Selected:selected,Placeholder:"SELECT",Highlighted:-1,OnChanged:changed}
	for i,v:=range options { if v==selected { s.Highlighted=i; break } }
	s.ExtendBaseWidget(s)
	return s
}

func (s *Select) SetSelected(value string) {
	s.Selected=value
	for i,v:=range s.Options { if v==value { s.Highlighted=i; break } }
	s.Refresh()
}
func (s *Select) SetDisabled(v bool) { s.Disabled=v; if v { s.Open=false }; s.Refresh() }
func (s *Select) FocusGained() { s.Focused=true; s.Refresh() }
func (s *Select) FocusLost() { s.Focused=false; s.Open=false; s.Refresh() }
func (s *Select) Tapped(*fyne.PointEvent) { if !s.Disabled { s.Open=!s.Open; s.Refresh() } }
func (s *Select) TypedRune(r rune) { if r==' ' { s.Tapped(nil) } }
func (s *Select) TypedKey(e *fyne.KeyEvent) {
	if s.Disabled { return }
	switch e.Name {
	case fyne.KeyReturn,fyne.KeyEnter:
		if s.Open { s.commitHighlighted() } else { s.Open=true; s.ensureHighlight() }
	case fyne.KeyEscape:
		s.Open=false
	case fyne.KeyDown:
		s.Open=true; s.moveHighlight(1)
	case fyne.KeyUp:
		s.Open=true; s.moveHighlight(-1)
	}
	s.Refresh()
}
func (s *Select) ensureHighlight() {
	if len(s.Options)==0 { s.Highlighted=-1; return }
	if s.Highlighted<0 { s.Highlighted=0 }
}
func (s *Select) moveHighlight(delta int) {
	if len(s.Options)==0 { s.Highlighted=-1; return }
	s.ensureHighlight()
	s.Highlighted=(s.Highlighted+delta+len(s.Options))%len(s.Options)
}
func (s *Select) commitHighlighted() {
	if s.Highlighted<0 || s.Highlighted>=len(s.Options) { return }
	v:=s.Options[s.Highlighted]; s.Selected=v; s.Open=false
	if s.OnChanged!=nil { s.OnChanged(v) }
}
func (s *Select) CreateRenderer() fyne.WidgetRenderer { return &selectRenderer{owner:s} }

type selectRenderer struct { owner *Select; objects []fyne.CanvasObject }
func (r *selectRenderer) MinSize() fyne.Size { return fyne.NewSize(220,42) }
func (r *selectRenderer) Layout(size fyne.Size) { r.rebuild(size) }
func (r *selectRenderer) Refresh() { r.rebuild(r.owner.Size()); canvas.Refresh(r.owner) }
func (r *selectRenderer) Objects() []fyne.CanvasObject { return r.objects }
func (r *selectRenderer) Destroy() {}

func (r *selectRenderer) rebuild(size fyne.Size) {
	s:=DarkScheme(); alpha:=uint8(0xff); if r.owner.Disabled { alpha=0x80 }
	border:=s.BorderBright; if r.owner.Focused || r.owner.Open { border=s.PrimaryActive }
	bg:=NewChamfer(withColorAlpha(s.SurfaceInset,alpha),withColorAlpha(border,alpha),s.Bevel.MD).Object(fyne.NewSize(size.Width,42))
	value:=r.owner.Selected; if value=="" { value=r.owner.Placeholder }
	text:=DataText(strings.ToUpper(value),13,withColorAlpha(s.TextBright,alpha))
	caret:="▼"; if r.owner.Open { caret="▲" }
	caretText:=DataText(caret,11,withColorAlpha(s.PrimaryActive,alpha))
	trigger:=container.NewPadded(container.NewHBox(text,layout.NewSpacer(),caretText))
	trigger.Resize(fyne.NewSize(size.Width,42))
	objects:=[]fyne.CanvasObject{container.NewWithoutLayout(bg,trigger)}
	if r.owner.Open && !r.owner.Disabled {
		rows:=make([]fyne.CanvasObject,0,len(r.owner.Options))
		for i,opt:=range r.owner.Options {
			fg:=s.TextMid
			prefix:="  "
			if opt==r.owner.Selected { prefix="✓ "; fg=s.PrimaryActive }
			if i==r.owner.Highlighted { fg=s.TextBright }
			row:=DataText(prefix+strings.ToUpper(opt),12,fg)
			rows=append(rows,container.NewPadded(row))
		}
		list:=container.NewVBox(rows...)
		popupH:=list.MinSize().Height+8
		popupBg:=NewChamfer(s.SurfacePanel,s.BorderBright,s.Bevel.SM).Object(fyne.NewSize(size.Width,popupH))
		popup:=container.NewWithoutLayout(popupBg,container.NewPadded(list))
		popup.Move(fyne.NewPos(0,46)); popup.Resize(fyne.NewSize(size.Width,popupH))
		objects=append(objects,popup)
	}
	root:=container.NewWithoutLayout(objects...)
	root.Resize(size)
	r.objects=[]fyne.CanvasObject{root}
}
