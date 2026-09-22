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
	chat := container.NewVBox(
		auris.DisplayText("CHAT // COMMS", 18, auris.Gold),
		auris.NewChatBubble(auris.ChatBubble{Author: "Rin", Role: auris.ChatAssistant, Message: "Ready for the next operation.", Timestamp: "13:14:03", Width: 390}),
		container.NewHBox(
			container.NewGridWrap(fyne.NewSize(170, 1)),
			auris.NewChatBubble(auris.ChatBubble{Author: "Zero", Role: auris.ChatUser, Message: "Show me the Auris version.", Timestamp: "13:14:28", Width: 390}),
		),
		auris.NewChatBubble(auris.ChatBubble{Author: "Rin", Role: auris.ChatAssistant, State: auris.ChatSending, Timestamp: "13:15:22", Width: 300}),
	)

	content := container.NewVBox(title, subtitle, panel, progress, badges, stats, notice, steps, sw, radio, selectControl, scan, hex, terminal, chat)
	w.SetContent(container.NewVScroll(container.NewPadded(content)))
	w.ShowAndRun()
}
