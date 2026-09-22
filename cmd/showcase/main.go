package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"github.com/dividebyzero/auris-go/auris"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(auris.NewTheme())
	w := a.NewWindow("Auris Go")
	w.Resize(fyne.NewSize(920, 760))

	title := auris.DisplayText("AURIS // FYNE", 26, auris.TextBright)
	subtitle := auris.BodyText("augmentation-era interface system", 14, auris.TextMid)

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

	sw := auris.NewSwitch("Primary reactor", true, func(bool) {})
	sw.OffLabel, sw.OnLabel = "OFFLINE", "ONLINE"
	radio := auris.NewRadio("Channel Alpha", true, func() {})
	selectControl := auris.NewSelect([]string{"ALPHA", "BETA", "GAMMA"}, "ALPHA", func(string) {})
	scanLabel := auris.DataText("SUBJECT LOCK", 14, auris.Gold)
	scanLabel.Alignment = fyne.TextAlignCenter
	scan := auris.NewScanBracket(scanLabel, fyne.NewSize(220, 64), 14, 2)
	hex := auris.NewHexOrnament(fyne.NewSize(180, 90), 18)
	stats := container.NewHBox(
		auris.NewStatCard("Throughput", "847", "req/s", "+12.4%", true, fyne.NewSize(200, 120)),
		auris.NewStatCard("Latency", "18", "ms", "-3.1%", false, fyne.NewSize(200, 120)),
	)
	notice := auris.NewNotification("Link established", "Desktop body channel is responding.", "DL-01", auris.NotificationSuccess, fyne.NewSize(620, 78))
	steps := container.NewHBox(
		auris.NewStepIndicator(1, auris.StepComplete, 28),
		auris.NewStepIndicator(2, auris.StepActive, 28),
		auris.NewStepIndicator(3, auris.StepInactive, 28),
		auris.NewStepIndicator(4, auris.StepError, 28),
	)
	terminal := auris.NewTerminal("System Log", "LIVE", []auris.TerminalLine{
		{Text: "[boot] neural interface online", Type: auris.TerminalNormal},
		{Text: "[ok] telemetry link established", Type: auris.TerminalOK},
		{Text: "[augment] body channel attached", Type: auris.TerminalAugment},
		{Text: "[warn] synthetic demo data", Type: auris.TerminalWarning},
	}, fyne.NewSize(620, 190))

	content := container.NewVBox(title, subtitle, panel, progress, badges, stats, notice, steps, sw, radio, selectControl, scan, hex, terminal)
	w.SetContent(container.NewVScroll(container.NewPadded(content)))
	w.ShowAndRun()
}
