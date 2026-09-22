package auris

import "testing"

func TestRadioGroupIsExclusive(t *testing.T) {
	got:=-1
	g:=NewRadioGroup([]string{"A","B","C"},0,func(i int){ got=i })
	g.Select(2)
	if g.Selected!=2 || got!=2 { t.Fatal("group must report selected index") }
	for i,r:=range g.Radios {
		if r.Selected!=(i==2) { t.Fatalf("radio %d exclusivity mismatch",i) }
	}
}
