package auris

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type NotificationVariant int
const (
	NotificationInfo NotificationVariant=iota
	NotificationSuccess
	NotificationWarning
	NotificationError
)

func notificationColor(v NotificationVariant,s Scheme) color.Color {
	switch v {
	case NotificationSuccess:return s.Success
	case NotificationError:return s.Danger
	case NotificationWarning:return s.PrimaryDim
	default:return s.PrimaryActive
	}
}
func notificationIcon(v NotificationVariant) string {
	switch v {
	case NotificationSuccess:return "✓"
	case NotificationWarning:return "!"
	case NotificationError:return "×"
	default:return "i"
	}
}

func NewNotification(title,message,code string,variant NotificationVariant,size fyne.Size) fyne.CanvasObject {
	return NewDismissibleNotification(title,message,code,variant,size,nil)
}

func NewDismissibleNotification(title,message,code string,variant NotificationVariant,size fyne.Size,onDismiss func()) fyne.CanvasObject {
	s:=CurrentScheme(); accent:=notificationColor(variant,s)
	bar:=canvas.NewRectangle(accent); bar.SetMinSize(fyne.NewSize(4,1))
	icon:=DataText(notificationIcon(variant),15,accent)
	t:=DisplayText(strings.ToUpper(title),14,accent)
	c:=DataText(code,11,s.TextMid)
	headItems:=[]fyne.CanvasObject{icon,t,layout.NewSpacer(),c}
	if onDismiss!=nil {
		close:=NewActionText("×",s.TextMid,onDismiss)
		headItems=append(headItems,close)
	}
	head:=container.NewHBox(headItems...)
	m:=BodyText(message,13,s.TextMid)
	body:=container.NewPadded(container.NewVBox(head,m))
	if size.Width<=0 { size.Width=body.MinSize().Width+20 }
	if size.Height<=0 { size.Height=body.MinSize().Height+12 }
	bar.SetMinSize(fyne.NewSize(4,size.Height))
	row:=container.NewBorder(nil,nil,bar,nil,body)
	return NewContainer(row,size,s.Bevel.MD,s.SurfacePanel,withColorAlpha(accent,0x73))
}
