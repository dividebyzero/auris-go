package auris

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type NotificationVariant int
const (
	NotificationInfo NotificationVariant = iota
	NotificationSuccess
	NotificationWarning
	NotificationError
)

func notificationColor(v NotificationVariant,s Scheme) fyne.Color {
	switch v {
	case NotificationSuccess: return s.Success
	case NotificationError: return s.Danger
	case NotificationWarning: return s.PrimaryDim
	default: return s.PrimaryActive
	}
}

func NewNotification(title,message,code string,variant NotificationVariant,size fyne.Size) fyne.CanvasObject {
	s:=DarkScheme()
	accent:=notificationColor(variant,s)
	bar:=canvas.NewRectangle(accent); bar.SetMinSize(fyne.NewSize(4,size.Height))
	t:=DisplayText(strings.ToUpper(title), 14, accent)
	c:=DataText(code, 11, s.TextMid)
	head:=container.NewHBox(t,layout.NewSpacer(),c)
	m:=BodyText(message, 13, s.TextMid)
	body:=container.NewPadded(container.NewVBox(head,m))
	row:=container.NewBorder(nil,nil,bar,nil,body)
	return NewContainer(row,size,s.Bevel.MD,s.SurfacePanel,withColorAlpha(accent,0x73))
}
