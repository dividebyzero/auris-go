package auris

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Theme struct{}

func NewTheme() fyne.Theme { return &Theme{} }

func schemeForVariant(v fyne.ThemeVariant) Scheme {
	if v==theme.VariantLight { return LightScheme() }
	return DarkScheme()
}

func (t *Theme) Color(name fyne.ThemeColorName,variant fyne.ThemeVariant) color.Color {
	s:=schemeForVariant(variant)
	SetScheme(s)
	switch name {
	case theme.ColorNameBackground:return s.SurfacePage
	case theme.ColorNameButton,theme.ColorNameInputBackground:return s.SurfaceInset
	case theme.ColorNameForeground:return s.TextBright
	case theme.ColorNamePrimary:return s.PrimaryActive
	case theme.ColorNameHover:return withColorAlpha(s.PrimaryDim,0x24)
	case theme.ColorNameFocus:return s.Highlight
	case theme.ColorNameDisabled:return s.TextMid
	case theme.ColorNameError:return s.Danger
	case theme.ColorNameSuccess:return s.Success
	default:return theme.DefaultTheme().Color(name,variant)
	}
}
func (t *Theme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Monospace{return FontData}; if style.Bold{return FontDisplayStrong}; return FontBody
}
func (t *Theme) Icon(name fyne.ThemeIconName) fyne.Resource { return theme.DefaultTheme().Icon(name) }
func (t *Theme) Size(name fyne.ThemeSizeName) float32 { return theme.DefaultTheme().Size(name) }
