package auris

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func NewPanel(title, code string, body fyne.CanvasObject, size fyne.Size, accent bool) fyne.CanvasObject {
	s := DarkScheme()
	titleColor := s.TextBright
	if accent { titleColor = s.PrimaryActive }

	left := canvas.NewText("⌐", s.PrimaryDim)
	left.TextSize = 15
	heading := canvas.NewText(strings.ToUpper(title), titleColor)
	heading.TextSize = 15
	right := canvas.NewText("¬", s.PrimaryDim)
	right.TextSize = 15
	status := canvas.NewText(code, s.TextMid)
	status.TextSize = 11

	header := container.NewHBox(left, heading, right, layout.NewSpacer(), status)
	inside := container.NewPadded(container.NewVBox(header, NewRule(size.Width), body))
	return NewPanelSurface(inside, size, accent)
}
