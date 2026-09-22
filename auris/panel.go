package auris

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func NewPanel(title, code string, body fyne.CanvasObject, size fyne.Size, accent bool) fyne.CanvasObject {
	s := DarkScheme()
	titleColor := s.TextBright
	if accent { titleColor = s.PrimaryActive }

	left := DisplayText("⌐", 15, s.PrimaryDim)
heading := DisplayText(strings.ToUpper(title), 15, titleColor)
right := DisplayText("¬", 15, s.PrimaryDim)
status := DataText(code, 11, s.TextMid)
header := container.NewHBox(left, heading, right, layout.NewSpacer(), status)
	inside := container.NewPadded(container.NewVBox(header, NewRule(size.Width), body))
	return NewPanelSurface(inside, size, accent)
}
