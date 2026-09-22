package auris

import (
	"testing"
	"fyne.io/fyne/v2"
)

func TestSwitchKeyboardAndDisabledBehavior(t *testing.T) {
	calls:=0
	s:=NewSwitch("test",false,func(bool){ calls++ })
	s.TypedRune(' ')
	if !s.Value || calls!=1 { t.Fatal("space should toggle switch") }
	s.TypedKey(&fyne.KeyEvent{Name:fyne.KeyReturn})
	if s.Value || calls!=2 { t.Fatal("return should toggle switch") }
	s.SetDisabled(true)
	s.Tapped(nil)
	if s.Value || calls!=2 { t.Fatal("disabled switch must ignore activation") }
}

func TestSwitchThumbProgressClamps(t *testing.T) {
	s:=NewSwitch("test",false,nil)
	s.SetThumbProgress(-1)
	if s.ThumbProgress!=0 { t.Fatal("thumb progress must clamp low") }
	s.SetThumbProgress(2)
	if s.ThumbProgress!=1 { t.Fatal("thumb progress must clamp high") }
	s.SetValue(false)
	if s.ThumbProgress!=0 { t.Fatal("settled off switch must place thumb at start") }
	s.SetValue(true)
	if s.ThumbProgress!=1 { t.Fatal("settled on switch must place thumb at end") }
}
