package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/dividebyzero/auris-go/auris"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(auris.NewTheme())

	w := a.NewWindow("Auris Go")
	w.Resize(fyne.NewSize(920, 640))

	title := canvas.NewText("AURIS // FYNE", auris.TextBright)
	title.TextSize = 26
	subtitle := canvas.NewText("augmentation-era interface system", auris.TextMid)
	subtitle.TextSize = 14

	progress := widget.NewProgressBar()
	progress.SetValue(.68)

	content := container.NewVBox(
		title,
		subtitle,
		widget.NewSeparator(),
		widget.NewLabel("REACTOR CORE"),
		progress,
		widget.NewCheck("AUTO-STABILIZER", nil),
		widget.NewButton("INITIALIZE", func() {}),
		widget.NewEntry(),
	)

	w.SetContent(container.NewPadded(content))
	w.ShowAndRun()
}
