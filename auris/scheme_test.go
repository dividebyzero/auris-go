package auris

import "testing"

func TestDarkSchemeUsesCanonicalTokens(t *testing.T) {
	s := DarkScheme()
	if s.SurfacePage != Void { t.Fatal("page surface must use Void") }
	if s.SurfacePanel != Panel { t.Fatal("panel surface must use Panel") }
	if s.PrimaryActive != Gold { t.Fatal("active primary must use Gold") }
	if s.Bevel.XS != BevelXS || s.Bevel.XL != BevelXL { t.Fatal("bevel scale must use canonical tokens") }
}

func TestSchemeSemanticStatesAreDistinct(t *testing.T) {
	s := DarkScheme()
	if s.Danger == s.Success { t.Fatal("danger and success must remain visually distinct") }
	if s.Border == s.BorderBright { t.Fatal("resting and bright borders must remain distinct") }
}
