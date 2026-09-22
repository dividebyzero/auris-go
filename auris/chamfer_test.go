package auris

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func TestChamferProducesSignaturePolygon(t *testing.T) {
	obj := NewChamfer(Panel, BorderBright, BevelMD).Object(fyne.NewSize(100, 50))
	p, ok := obj.(*canvas.Polygon)
	if !ok { t.Fatalf("expected polygon, got %T", obj) }
	if len(p.Points) != 6 { t.Fatalf("expected 6 points, got %d", len(p.Points)) }
	want := []fyne.Position{
		{X:10,Y:0}, {X:100,Y:0}, {X:100,Y:40},
		{X:90,Y:50}, {X:0,Y:50}, {X:0,Y:10},
	}
	for i := range want {
		if p.Points[i] != want[i] { t.Fatalf("point %d: got %v want %v", i, p.Points[i], want[i]) }
	}
}

func TestChamferClampsCutToHalfSmallestDimension(t *testing.T) {
	obj := NewChamfer(Panel, Border, 99).Object(fyne.NewSize(20, 10)).(*canvas.Polygon)
	if obj.Points[0].X != 5 || obj.Points[2].Y != 5 {
		t.Fatalf("cut was not clamped correctly: %v", obj.Points)
	}
}
