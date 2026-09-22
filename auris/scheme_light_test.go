package auris

import "testing"

func TestLightSchemeIsDistinct(t *testing.T) {
	d,l:=DarkScheme(),LightScheme()
	if d.SurfacePage==l.SurfacePage || d.TextBright==l.TextBright { t.Fatal("light scheme must be visually distinct") }
}
func TestActiveSchemeCanSwitch(t *testing.T) {
	UseLightScheme()
	if CurrentScheme().SurfacePage!=LightScheme().SurfacePage { t.Fatal("light scheme was not activated") }
	UseDarkScheme()
	if CurrentScheme().SurfacePage!=DarkScheme().SurfacePage { t.Fatal("dark scheme was not restored") }
}
