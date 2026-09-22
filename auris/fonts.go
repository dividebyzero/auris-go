package auris

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

// The canonical Auris font assets are shared with the Flutter implementation.
//
//go:embed fonts/Rajdhani-Medium.ttf
var rajdhaniMediumBytes []byte

//go:embed fonts/Rajdhani-SemiBold.ttf
var rajdhaniSemiBoldBytes []byte

//go:embed fonts/Rajdhani-Bold.ttf
var rajdhaniBoldBytes []byte

//go:embed fonts/Exo2-Regular.ttf
var exoTwoBytes []byte

//go:embed fonts/ShareTechMono-Regular.ttf
var shareTechMonoBytes []byte

var (
	FontDisplay       = fyne.NewStaticResource("Rajdhani-Medium.ttf", rajdhaniMediumBytes)
	FontDisplayStrong = fyne.NewStaticResource("Rajdhani-SemiBold.ttf", rajdhaniSemiBoldBytes)
	FontDisplayBold   = fyne.NewStaticResource("Rajdhani-Bold.ttf", rajdhaniBoldBytes)
	FontBody          = fyne.NewStaticResource("Exo2-Regular.ttf", exoTwoBytes)
	FontData          = fyne.NewStaticResource("ShareTechMono-Regular.ttf", shareTechMonoBytes)
)

func FontFor(role TypeRole) fyne.Resource {
	switch role {
	case TypeDisplay:
		return FontDisplay
	case TypeData:
		return FontData
	default:
		return FontBody
	}
}
