package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

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

	rows := container.NewVBox(
		auris.NewDataRow("Core temp", "612 K", false),
		auris.NewDataRow("Output", "99.2 %", true),
		auris.NewDataRow("Status", "NOMINAL", false),
	)
	panel := auris.NewPanel("Reactor Core", "RC-09", rows, fyne.NewSize(620, 190), true)

	progress := auris.NewProgress(.68, 20, "Shield integrity", "68 / 100", auris.ProgressPrimary)
	badges := container.NewHBox(
		auris.NewBadge("online", auris.BadgeSuccess),
		auris.NewBadge("armed", auris.BadgeGold),
		auris.NewBadge("warning", auris.BadgeDanger),
	)

	content := container.NewVBox(title, subtitle, panel, progress, badges)
	w.SetContent(container.NewPadded(content))
	w.ShowAndRun()
}
