package auris

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func NewDataRow(label, value string, highlight bool) fyne.CanvasObject {
	s := DarkScheme()
	l := canvas.NewText(strings.ToUpper(label), s.TextMid)
	l.TextSize = 12
	valueColor := s.TextBright
	if highlight {
		valueColor = s.Highlight
	}
	v := canvas.NewText(value, valueColor)
	v.TextSize = 13
	row := container.NewHBox(l, layout.NewSpacer(), v)
	return container.NewVBox(row, NewRule(1))
}
