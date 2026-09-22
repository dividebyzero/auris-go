package auris

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewStatCard(label, value, unit, delta string, positiveIsGood bool, size fyne.Size) fyne.CanvasObject {
	s := DarkScheme()
	l := BodyText(strings.ToUpper(label), 12, s.TextMid)
	v := DisplayText(value, 34, s.PrimaryActive)
	u := DataText(unit, 14, s.TextMid)
	valueRow := container.NewHBox(v, u)
	items := []fyne.CanvasObject{l, valueRow}

	if delta != "" {
		negative := strings.HasPrefix(strings.TrimSpace(delta), "-")
		good := positiveIsGood != negative
		c := s.Danger
		arrow := "↑ "
		if negative {
			arrow = "↓ "
		}
		if good {
			c = s.Success
		}
		items = append(items, DataText(arrow+delta, 13, c))
	}

	body := container.NewPadded(container.NewVBox(items...))
	return NewContainer(body, size, s.Bevel.MD, s.SurfacePanel, s.BorderBright)
}
