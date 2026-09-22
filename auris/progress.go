package auris

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type ProgressVariant int

const (
	ProgressPrimary ProgressVariant = iota
	ProgressSecondary
	ProgressDanger
	ProgressSuccess
)

func progressColor(v ProgressVariant, s Scheme) color.Color {
	switch v {
	case ProgressSecondary:
		return s.Secondary
	case ProgressDanger:
		return s.Danger
	case ProgressSuccess:
		return s.Success
	default:
		return s.PrimaryActive
	}
}

func NewProgress(value float64, segments int, label, valueLabel string, variant ProgressVariant) fyne.CanvasObject {
	if value < 0 {
		value = 0
	}
	if value > 1 {
		value = 1
	}
	if segments < 1 {
		segments = 1
	}

	s := DarkScheme()
	active := progressColor(variant, s)
	filled := int(value*float64(segments) + .5)
	cells := make([]fyne.CanvasObject, 0, segments)
	for i := 0; i < segments; i++ {
		c := withColorAlpha(s.BorderBright, 0x70)
		if i < filled {
			c = withColorAlpha(active, 0x72)
		}
		if i == filled-1 {
			c = active
		}
		cells = append(cells, newSlant(fyne.NewSize(18, 10), 3, c, c))
	}

	bar := container.New(layout.NewGridLayoutWithColumns(segments), cells...)
	if label == "" && valueLabel == "" {
		return bar
	}

	l := DataText(strings.ToUpper(label), 11, s.PrimaryDim)
	v := DataText(valueLabel, 11, s.TextMid)
	return container.NewVBox(container.NewHBox(l, layout.NewSpacer(), v), bar)
}
