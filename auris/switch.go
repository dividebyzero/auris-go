package auris

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Switch struct {
	widget.BaseWidget
	Value bool
	Label, OffLabel, OnLabel string
	Disabled bool
	Focused bool
	Motion Motion
	ThumbProgress float32
	OnChanged func(bool)
}

func NewSwitch(label string, value bool, changed func(bool)) *Switch {
	progress:=float32(0); if value { progress=1 }
	s := &Switch{Label:label, Value:value, ThumbProgress:progress, OffLabel:"OFF", OnLabel:"ON", OnChanged:changed}
	s.ExtendBaseWidget(s)
	return s
}

func (s *Switch) SetValue(value bool) {
	if s.Value == value { return }
	from:=s.ThumbProgress
	s.Value=value
	to:=float32(0); if value { to=1 }
	AnimateValue(s.Motion,DurationNormal,from,to,func(v float32){ s.ThumbProgress=v; s.Refresh() })
	if s.OnChanged != nil { s.OnChanged(value) }
}

func (s *Switch) SetReducedMotion(reduced bool) { s.Motion.Reduced=reduced }
func (s *Switch) SetThumbProgress(progress float32) { if progress<0 { progress=0 }; if progress>1 { progress=1 }; s.ThumbProgress=progress; s.Refresh() }
func (s *Switch) SetDisabled(disabled bool) { s.Disabled=disabled; s.Refresh() }
func (s *Switch) FocusGained() { s.Focused=true; s.Refresh() }
func (s *Switch) FocusLost() { s.Focused=false; s.Refresh() }
func (s *Switch) TypedRune(r rune) { if r == ' ' { s.Tapped(nil) } }
func (s *Switch) TypedKey(e *fyne.KeyEvent) { if e.Name == fyne.KeyReturn || e.Name == fyne.KeyEnter { s.Tapped(nil) } }
func (s *Switch) Tapped(*fyne.PointEvent) {
	if s.Disabled { return }
	s.SetValue(!s.Value)
}
func (s *Switch) CreateRenderer() fyne.WidgetRenderer { return &switchRenderer{owner:s} }

type switchRenderer struct { owner *Switch; objects []fyne.CanvasObject }
func (r *switchRenderer) Layout(size fyne.Size) { r.rebuild(size) }
func (r *switchRenderer) MinSize() fyne.Size { return fyne.NewSize(210,32) }
func (r *switchRenderer) Refresh() { r.rebuild(r.owner.Size()); canvas.Refresh(r.owner) }
func (r *switchRenderer) Objects() []fyne.CanvasObject { return r.objects }
func (r *switchRenderer) Destroy() {}

func (r *switchRenderer) rebuild(size fyne.Size) {
	s:=DarkScheme()
	alpha:=uint8(0xff)
	if r.owner.Disabled { alpha=0x80 }
	label:=BodyText(r.owner.Label,14,withColorAlpha(s.TextBright,alpha))
	trackSize:=fyne.NewSize(48,24)
	trackFill,trackBorder,thumb:=s.SurfaceInset,s.Border,s.PrimaryDim
	if r.owner.Focused { trackBorder=s.PrimaryActive }
	thumbX:=Lerp(4,28,r.owner.ThumbProgress)
	status:=r.owner.OffLabel
	if r.owner.Value {
		trackFill=withColorAlpha(Gold,0x38)
		trackBorder,thumb,status=s.PrimaryActive,s.PrimaryActive,r.owner.OnLabel
	}
	track:=newSlant(trackSize,4,withColorAlpha(trackFill,alpha),withColorAlpha(trackBorder,alpha))
	knob:=newSlant(fyne.NewSize(16,16),3,withColorAlpha(thumb,alpha),withColorAlpha(thumb,alpha))
	knob.Move(fyne.NewPos(thumbX,4))
	trackStack:=container.NewWithoutLayout(track,knob); trackStack.Resize(trackSize)
	stateColor:=color.Color(s.TextMid)
	if r.owner.Value { stateColor=s.PrimaryActive }
	state:=DataText(strings.ToUpper(status),12,withColorAlpha(stateColor,alpha))
	row:=container.NewHBox(label,trackStack,state); row.Resize(size)
	r.objects=[]fyne.CanvasObject{row}
}
