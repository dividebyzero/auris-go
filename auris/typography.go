package auris

const (
	TrackingBody    float32 = 0.5
	TrackingLabel   float32 = 1.5
	TrackingHeading float32 = 1.8
	TrackingButton  float32 = 1.44
)

// TypeRole documents the three-font Auris typography contract. Fyne's Theme
// API exposes one resource per TextStyle, so exact family assignment belongs
// in Auris-native text widgets rather than pretending bold/mono map 1:1.
type TypeRole int
const (
	TypeDisplay TypeRole = iota // Rajdhani in the Flutter reference
	TypeBody                    // Exo 2
	TypeData                    // Share Tech Mono
)
