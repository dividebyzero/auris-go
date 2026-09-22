package auris

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type StepState int
const (
	StepInactive StepState=iota
	StepActive
	StepComplete
	StepError
)

func NewStepIndicator(step int,state StepState,size float32) fyne.CanvasObject {
	s:=CurrentScheme(); if size<=0 { size=28 }
	border,fill,fg:=s.Border,s.SurfaceInset,s.TextMid
	text:=strconv.Itoa(step); glow:=false
	switch state {
	case StepActive:
		border,fill,fg,glow=s.PrimaryActive,withColorAlpha(Gold,0x29),s.PrimaryActive,true
	case StepComplete:
		border,fill,fg,text=s.PrimaryActive,s.PrimaryActive,Void,"✓"
	case StepError:
		border,fill,fg,text,glow=s.Danger,withColorAlpha(DangerBright,0x38),s.Danger,"!",true
	}
	bg:=NewChamfer(fill,border,s.Bevel.XS).Object(fyne.NewSize(size,size))
	var label fyne.CanvasObject
	if glow { label=NewGlowText(text,size*.42,fg,DepthSubtle) } else {
		t:=DataText(text,size*.42,fg); t.Alignment=fyne.TextAlignCenter; label=t
	}
	label.Resize(fyne.NewSize(size,size)); label.Move(fyne.NewPos(0,size*.24))
	stack:=container.NewWithoutLayout(bg,label); stack.Resize(fyne.NewSize(size,size))
	return stack
}
