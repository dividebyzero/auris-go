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
	Value       bool
	Label       string
	OffLabel    string
	OnLabel     string
	OnChanged   func(bool)
}

func NewSwitch(label string, value bool, changed func(bool)) *Switch {
	s := &Switch{Label: label, Value: value, OffLabel: "OFF", OnLabel: "ON", OnChanged: changed}
	s.ExtendBaseWidget(s)
	return s
}

func (s *Switch) Tapped(*fyne.PointEvent) {
	if s.OnChanged == nil { return }
	s.Value = !s.Value
	s.OnChanged(s.Value)
	s.Refresh()
}

func (s *Switch) CreateRenderer() fyne.WidgetRenderer {
	return &switchRenderer{owner: s, objects: []fyne.CanvasObject{}}
}

type switchRenderer struct {
	owner *Switch
	objects []fyne.CanvasObject
}

func (r *switchRenderer) Layout(size fyne.Size) {
	r.rebuild(size)
}
func (r *switchRenderer) MinSize() fyne.Size { return fyne.NewSize(210, 32) }
func (r *switchRenderer) Refresh() { r.rebuild(r.owner.Size()); canvas.Refresh(r.owner) }
func (r *switchRenderer) Objects() []fyne.CanvasObject { return r.objects }
func (r *switchRenderer) Destroy() {}

func (r *switchRenderer) rebuild(size fyne.Size) {
	scheme := DarkScheme()
	label := BodyText(r.owner.Label, 14, scheme.TextBright)
trackSize := fyne.NewSize(48, 24)
	trackFill, trackBorder, thumb := scheme.SurfaceInset, scheme.Border, scheme.PrimaryDim
	thumbX := float32(4)
	status := r.owner.OffLabel
	if r.owner.Value {
		trackFill = withColorAlpha(Gold, 0x38)
		trackBorder, thumb, thumbX, status = scheme.PrimaryActive, scheme.PrimaryActive, 28, r.owner.OnLabel
	}
	track := newSlant(trackSize, 4, trackFill, trackBorder)
	knob := newSlant(fyne.NewSize(16,16), 3, thumb, thumb)
	knob.Move(fyne.NewPos(thumbX,4))
	trackStack := container.NewWithoutLayout(track, knob)
	trackStack.Resize(trackSize)

	state := DataText(strings.ToUpper(status), 12, func() color.Color {
		if r.owner.Value { return scheme.PrimaryActive }
		return scheme.TextMid
	}())
row := container.NewHBox(label, trackStack, state)
	row.Resize(size)
	r.objects = []fyne.CanvasObject{row}
}
