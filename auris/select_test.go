package auris

import (
	"testing"
	"fyne.io/fyne/v2"
)

func TestSelectKeyboardNavigationAndCommit(t *testing.T) {
	got:=""
	s:=NewSelect([]string{"ALPHA","BETA","GAMMA"},"ALPHA",func(v string){ got=v })
	s.TypedKey(&fyne.KeyEvent{Name:fyne.KeyDown})
	if !s.Open || s.Highlighted!=1 { t.Fatalf("down should open and highlight beta: open=%v index=%d",s.Open,s.Highlighted) }
	s.TypedKey(&fyne.KeyEvent{Name:fyne.KeyEnter})
	if s.Selected!="BETA" || got!="BETA" || s.Open { t.Fatal("enter must commit highlighted option and close") }
}

func TestSelectNavigationWraps(t *testing.T) {
	s:=NewSelect([]string{"A","B","C"},"A",nil)
	s.TypedKey(&fyne.KeyEvent{Name:fyne.KeyUp})
	if s.Highlighted!=2 { t.Fatalf("up from first must wrap to last, got %d",s.Highlighted) }
}

func TestSelectDisabledIgnoresInput(t *testing.T) {
	s:=NewSelect([]string{"A","B"},"A",nil)
	s.SetDisabled(true)
	s.Tapped(nil)
	s.TypedKey(&fyne.KeyEvent{Name:fyne.KeyDown})
	if s.Open || s.Selected!="A" { t.Fatal("disabled select must ignore input") }
}
