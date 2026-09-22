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
