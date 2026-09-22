package auris

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func headerTick(left bool, c color.Color) fyne.CanvasObject {
	const w, h, stroke float32 = 9, 9, 1
	horizontal := canvas.NewRectangle(c)
	vertical := canvas.NewRectangle(c)
	horizontal.Resize(fyne.NewSize(w, stroke))
	vertical.Resize(fyne.NewSize(stroke, h))
	if !left {
		vertical.Move(fyne.NewPos(w-stroke, 0))
	}
	group := container.NewWithoutLayout(horizontal, vertical)
	group.Resize(fyne.NewSize(w, h))
	return group
}

func NewPanel(title, code string, body fyne.CanvasObject, size fyne.Size, accent bool) fyne.CanvasObject {
	s := CurrentScheme()
	titleColor := s.TextBright
	if accent {
		titleColor = s.PrimaryActive
	}

	heading := DisplayText(strings.ToUpper(title), 15, titleColor)
	status := DataText(code, 11, s.TextMid)
	header := container.NewHBox(
		headerTick(true, s.PrimaryDim),
		heading,
		headerTick(false, s.PrimaryDim),
		layout.NewSpacer(),
		status,
	)
	inside := container.NewPadded(container.NewVBox(header, NewRule(size.Width), body))
	return NewPanelSurface(inside, size, accent)
}
