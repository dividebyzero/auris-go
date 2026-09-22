package auris

import (
	"image/color"
	"sync"
)

type Scheme struct {
	SurfacePage,SurfacePanel,SurfaceInset color.Color
	TextBright,TextMid color.Color
	PrimaryDim,PrimaryActive,Highlight,Secondary color.Color
	Border,BorderBright,Danger,Success color.Color
	Bevel BevelScale
	GlowScale float32
}

type BevelScale struct { XS,SM,MD,LG,XL float32 }

func DarkScheme() Scheme {
	return Scheme{
		SurfacePage:Void,SurfacePanel:Panel,SurfaceInset:PanelAlt,
		TextBright:BrightWhite,TextMid:TextMid,
		PrimaryDim:Amber,PrimaryActive:Gold,Highlight:Bright,
		Secondary:Slate,Border:Border,BorderBright:BorderBright,
		Danger:DangerBright,Success:SuccessBright,
		Bevel:BevelScale{BevelXS,BevelSM,BevelMD,BevelLG,BevelXL},GlowScale:1,
	}
}

func LightScheme() Scheme {
	return Scheme{
		SurfacePage:color.NRGBA{0xf3,0xef,0xe5,0xff},
		SurfacePanel:color.NRGBA{0xe9,0xe2,0xd3,0xff},
		SurfaceInset:color.NRGBA{0xdd,0xd5,0xc5,0xff},
		TextBright:color.NRGBA{0x20,0x1b,0x14,0xff},
		TextMid:color.NRGBA{0x62,0x55,0x42,0xff},
		PrimaryDim:color.NRGBA{0x9a,0x62,0x00,0xff},
		PrimaryActive:color.NRGBA{0x7a,0x4b,0x00,0xff},
		Highlight:color.NRGBA{0xb8,0x70,0x00,0xff},
		Secondary:color.NRGBA{0x3e,0x67,0x70,0xff},
		Border:color.NRGBA{0xb7,0xa8,0x89,0xff},
		BorderBright:color.NRGBA{0x8e,0x72,0x3d,0xff},
		Danger:color.NRGBA{0xa3,0x2d,0x22,0xff},
		Success:color.NRGBA{0x3c,0x72,0x50,0xff},
		Bevel:BevelScale{BevelXS,BevelSM,BevelMD,BevelLG,BevelXL},GlowScale:.55,
	}
}

var schemeState=struct{ sync.RWMutex; current Scheme }{current:DarkScheme()}

func CurrentScheme() Scheme { schemeState.RLock(); defer schemeState.RUnlock(); return schemeState.current }
func SetScheme(s Scheme) { schemeState.Lock(); schemeState.current=s; schemeState.Unlock() }
func UseDarkScheme() { SetScheme(DarkScheme()) }
func UseLightScheme() { SetScheme(LightScheme()) }
