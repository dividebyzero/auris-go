package auris

import "time"

const (
	DurationFast   = 120 * time.Millisecond
	DurationNormal = 200 * time.Millisecond
	DurationSlow   = 350 * time.Millisecond
)

// Motion centralizes animation policy. Applications can disable motion to
// honor platform/user accessibility preferences.
type Motion struct { Reduced bool }

func (m Motion) Duration(d time.Duration) time.Duration {
	if m.Reduced { return 0 }
	return d
}
