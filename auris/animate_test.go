package auris

import "testing"

func TestAnimateReducedMotionCompletesSynchronously(t *testing.T) {
	calls:=0
	last:=float32(0)
	Animate(Motion{Reduced:true},DurationSlow,func(p float32){ calls++; last=p })
	if calls!=1 || last!=1 { t.Fatalf("reduced motion must complete once, got calls=%d last=%f",calls,last) }
}

func TestAnimateValueReducedMotionUsesExactTarget(t *testing.T) {
	got:=float32(0)
	AnimateValue(Motion{Reduced:true},DurationNormal,4,28,func(v float32){ got=v })
	if got!=28 { t.Fatalf("expected exact target 28, got %f",got) }
}
