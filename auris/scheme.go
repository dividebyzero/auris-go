package auris

import "image/color"

// Scheme is the resolved semantic Auris design layer. Widgets consume these
// roles instead of depending directly on primitive palette values.
type Scheme struct {
	SurfacePage   color.Color
	SurfacePanel  color.Color
	SurfaceInset  color.Color
	TextBright    color.Color
	TextMid       color.Color
	PrimaryDim    color.Color
	PrimaryActive color.Color
	Highlight     color.Color
	Secondary     color.Color
	Border        color.Color
	BorderBright  color.Color
	Danger        color.Color
	Success       color.Color
	Bevel         BevelScale
	GlowScale     float32
}

type BevelScale struct {
	XS, SM, MD, LG, XL float32
}

func DarkScheme() Scheme {
	return Scheme{
		SurfacePage: Void, SurfacePanel: Panel, SurfaceInset: PanelAlt,
		TextBright: BrightWhite, TextMid: TextMid,
		PrimaryDim: Amber, PrimaryActive: Gold, Highlight: Bright,
		Secondary: Slate, Border: Border, BorderBright: BorderBright,
		Danger: DangerBright, Success: SuccessBright,
		Bevel: BevelScale{BevelXS, BevelSM, BevelMD, BevelLG, BevelXL},
		GlowScale: 1,
	}
}
