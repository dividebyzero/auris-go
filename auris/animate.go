package auris

import (
	"time"

	"fyne.io/fyne/v2"
)

// Animate drives a deterministic 0..1 Auris transition. The caller owns the
// rendered state; this helper only supplies eased progress on Fyne's UI thread.
func Animate(m Motion, duration time.Duration, frame func(float32)) {
	if frame == nil { return }
	d:=m.Duration(duration)
	if d==0 { frame(1); return }

	start:=time.Now()
	var tick func()
	tick=func() {
		elapsed:=time.Since(start)
		p:=m.Progress(elapsed,duration)
		frame(p)
		if p>=1 { return }
		time.AfterFunc(time.Second/60,func(){ fyne.Do(tick) })
	}
	fyne.Do(tick)
}

// AnimateValue interpolates a scalar and guarantees the exact target value on
// the final frame.
func AnimateValue(m Motion, duration time.Duration, from,to float32, frame func(float32)) {
	if frame==nil { return }
	Animate(m,duration,func(p float32){ frame(Lerp(from,to,p)) })
}
