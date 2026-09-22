package auris

import (
	"testing"
	"fyne.io/fyne/v2"
)

func TestRadioKeyboardAndDisabledBehavior(t *testing.T) {
	calls:=0
	r:=NewRadio("test",false,func(){ calls++ })
	r.TypedKey(&fyne.KeyEvent{Name:fyne.KeyEnter})
	if !r.Selected || calls!=1 { t.Fatal("enter should select radio") }
	r.SetSelected(false)
	r.SetDisabled(true)
	r.TypedRune(' ')
	if r.Selected || calls!=1 { t.Fatal("disabled radio must ignore activation") }
}

func TestRadioFocusState(t *testing.T) {
	r:=NewRadio("test",false,nil)
	r.FocusGained()
	if !r.Focused { t.Fatal("focus gained must be recorded") }
	r.FocusLost()
	if r.Focused { t.Fatal("focus lost must be recorded") }
}
