package auris

import (
	"testing"
	"time"
)

func TestReducedMotionCompletesImmediately(t *testing.T) {
	m:=Motion{Reduced:true}
	if m.Progress(0,DurationSlow)!=1 { t.Fatal("reduced motion must skip transitions") }
}

func TestMotionProgressBounds(t *testing.T) {
	m:=Motion{}
	if m.Progress(-time.Millisecond,DurationNormal)!=0 { t.Fatal("negative elapsed must clamp to zero") }
	if m.Progress(DurationNormal,DurationNormal)!=1 { t.Fatal("completed transition must equal one") }
	p:=m.Progress(DurationNormal/2,DurationNormal)
	if p<.49 || p>.51 { t.Fatalf("midpoint should remain centered, got %f",p) }
}

func TestLerpClampsProgress(t *testing.T) {
	if Lerp(4,28,-1)!=4 || Lerp(4,28,2)!=28 { t.Fatal("lerp must clamp progress") }
}
