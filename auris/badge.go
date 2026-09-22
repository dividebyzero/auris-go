package auris

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type BadgeVariant int

const (
	BadgeAmber BadgeVariant = iota
	BadgeGold
	BadgeSlate
	BadgeDanger
	BadgeSuccess
)

func badgeColor(v BadgeVariant, s Scheme) color.Color {
	switch v {
	case BadgeGold: return s.PrimaryActive
	case BadgeSlate: return s.Secondary
	case BadgeDanger: return s.Danger
	case BadgeSuccess: return s.Success
	default: return s.PrimaryDim
	}
}

func NewBadge(label string, variant BadgeVariant) fyne.CanvasObject {
	s := DarkScheme()
	c := badgeColor(variant, s)
	t := DataText(strings.ToUpper(label), 11, c)
padded := container.NewPadded(t)
	return NewContainer(padded, fyne.NewSize(100, 28), s.Bevel.XS, withColorAlpha(c, 0x1f), withColorAlpha(c, 0x8c))
}

func withColorAlpha(c color.Color, a uint8) color.Color {
	r, g, b, _ := c.RGBA()
	return color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), a}
}
