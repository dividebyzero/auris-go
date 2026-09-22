package auris

import "time"

const (
	DurationFast   = 120 * time.Millisecond
	DurationNormal = 200 * time.Millisecond
	DurationSlow   = 350 * time.Millisecond
)

type Motion struct{ Reduced bool }

func (m Motion) Duration(d time.Duration) time.Duration {
	if m.Reduced { return 0 }
	return d
}

func Lerp(from, to, progress float32) float32 {
	if progress <= 0 { return from }
	if progress >= 1 { return to }
	return from + (to-from)*progress
}

// EaseInOut is the default Auris transition curve.
func EaseInOut(t float32) float32 {
	if t <= 0 { return 0 }
	if t >= 1 { return 1 }
	if t < .5 { return 2*t*t }
	return 1 - ((-2*t+2)*(-2*t+2))/2
}

func (m Motion) Progress(elapsed, duration time.Duration) float32 {
	d:=m.Duration(duration)
	if d==0 || elapsed>=d { return 1 }
	if elapsed<=0 { return 0 }
	return EaseInOut(float32(elapsed)/float32(d))
}
