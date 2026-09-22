package auris

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type StepState int
const (
	StepInactive StepState = iota
	StepActive
	StepComplete
	StepError
)

func NewStepIndicator(step int,state StepState,size float32) fyne.CanvasObject {
	s:=DarkScheme()
	border,fill,fg:=s.Border,s.SurfaceInset,s.TextMid
	text:=strconv.Itoa(step)
	switch state {
	case StepActive:
		border,fill,fg=s.PrimaryActive,withColorAlpha(Gold,0x29),s.PrimaryActive
	case StepComplete:
		border,fill,fg,text=s.PrimaryActive,s.PrimaryActive,Void,"✓"
	case StepError:
		border,fill,fg,text=s.Danger,withColorAlpha(DangerBright,0x38),s.Danger,"!"
	}
	bg:=NewChamfer(fill,border,s.Bevel.XS).Object(fyne.NewSize(size,size))
	label:=canvas.NewText(text,fg); label.TextSize=size*.42; label.Alignment=fyne.TextAlignCenter
	label.Resize(fyne.NewSize(size,size)); label.Move(fyne.NewPos(0,size*.24))
	stack:=container.NewWithoutLayout(bg,label); stack.Resize(fyne.NewSize(size,size))
	return stack
}
