package auris

import (
	"testing"

	"fyne.io/fyne/v2"
)

func TestChamferProducesSignaturePolygon(t *testing.T) {
	points := NewChamfer(Panel, BorderBright, BevelMD).points(fyne.NewSize(100, 50))
	if len(points) != 6 {
		t.Fatalf("expected 6 points, got %d", len(points))
	}
	want := []fyne.Position{
		{X: 10, Y: 0}, {X: 100, Y: 0}, {X: 100, Y: 40},
		{X: 90, Y: 50}, {X: 0, Y: 50}, {X: 0, Y: 10},
	}
	for i := range want {
		if points[i] != want[i] {
			t.Fatalf("point %d: got %v want %v", i, points[i], want[i])
		}
	}
}

func TestChamferClampsCutToHalfSmallestDimension(t *testing.T) {
	points := NewChamfer(Panel, Border, 99).points(fyne.NewSize(20, 10))
	if points[0].X != 5 || points[2].Y != 5 {
		t.Fatalf("cut was not clamped correctly: %v", points)
	}
}
