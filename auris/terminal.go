package auris

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type TerminalLineType int
const (
	TerminalNormal TerminalLineType = iota
	TerminalOK
	TerminalError
	TerminalAugment
	TerminalWarning
)

type TerminalLine struct { Text string; Type TerminalLineType }

func terminalColor(t TerminalLineType, s Scheme) color.Color {
	switch t {
	case TerminalOK: return s.Success
	case TerminalError: return s.Danger
	case TerminalAugment: return s.PrimaryActive
	case TerminalWarning: return s.PrimaryDim
	default: return s.TextMid
	}
}

func NewTerminal(title, code string, lines []TerminalLine, size fyne.Size) fyne.CanvasObject {
	s := DarkScheme()
	rows := make([]fyne.CanvasObject,0,len(lines))
	for i,line := range lines {
		text := line.Text
		if i == len(lines)-1 { text += "  █" }
		t := DataText(text, 12.5, terminalColor(line.Type,s))
.5
		rows=append(rows,t)
	}
	log := container.NewVBox(rows...)
	scroll := container.NewVScroll(log)
	scroll.SetMinSize(fyne.NewSize(size.Width-32,size.Height-54))
	return NewPanel(strings.ToUpper(title),code,scroll,size,false)
}
