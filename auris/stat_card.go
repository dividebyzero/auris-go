package auris

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func NewStatCard(label, value, unit, delta string, positiveIsGood bool, size fyne.Size) fyne.CanvasObject {
	s := DarkScheme()
	l := canvas.NewText(strings.ToUpper(label), s.TextMid); l.TextSize=12
	v := canvas.NewText(value, s.PrimaryActive); v.TextSize=34
	u := canvas.NewText(unit, s.TextMid); u.TextSize=14
	valueRow := container.NewHBox(v,u)
	items := []fyne.CanvasObject{l,valueRow}
	if delta != "" {
		negative := strings.HasPrefix(strings.TrimSpace(delta), "-")
		good := positiveIsGood != negative
		c := s.Danger
		arrow := "↑ "
		if negative { arrow="↓ " }
		if good { c=s.Success }
		d := canvas.NewText(arrow+delta,c); d.TextSize=13
		items=append(items,d)
	}
	body := container.NewPadded(container.NewVBox(items...))
	return NewContainer(body,size,s.Bevel.MD,s.SurfacePanel,s.BorderBright)
}
