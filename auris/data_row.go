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
	l := BodyText(strings.ToUpper(label), 12, s.TextMid)
valueColor := s.TextBright
	if highlight {
		valueColor = s.Highlight
	}
	v := DataText(value, 13, valueColor)
row := container.NewHBox(l, layout.NewSpacer(), v)
	return container.NewVBox(row, NewRule(1))
}
