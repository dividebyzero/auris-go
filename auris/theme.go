package auris

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// Theme maps standard Fyne widgets and Auris typography roles onto the
// canonical embedded assets. Auris custom text uses:
//   default   -> Exo 2
//   bold      -> Rajdhani SemiBold
//   monospace -> Share Tech Mono
type Theme struct{}

func NewTheme() fyne.Theme { return &Theme{} }

func (t *Theme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground: return Void
	case theme.ColorNameButton, theme.ColorNameInputBackground: return PanelAlt
	case theme.ColorNameForeground: return BrightWhite
	case theme.ColorNamePrimary: return Gold
	case theme.ColorNameHover: return withAlpha(Amber, 0x24)
	case theme.ColorNameFocus: return Bright
	case theme.ColorNameDisabled: return TextMid
	case theme.ColorNameError: return DangerBright
	case theme.ColorNameSuccess: return SuccessBright
	default: return theme.DefaultTheme().Color(name, theme.VariantDark)
	}
}

func (t *Theme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Monospace { return FontData }
	if style.Bold { return FontDisplayStrong }
	return FontBody
}

func (t *Theme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *Theme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func withAlpha(c color.NRGBA, alpha uint8) color.Color {
	c.A = alpha
	return c
}
