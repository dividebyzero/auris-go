package auris

import (
	"testing"

	"fyne.io/fyne/v2"
)

func TestFontResourcesAreEmbedded(t *testing.T) {
	fonts := []fyne.Resource{FontDisplay, FontDisplayStrong, FontDisplayBold, FontBody, FontData}
	for _, f := range fonts {
		if f == nil || len(f.Content()) == 0 { t.Fatal("expected embedded font resource") }
	}
}

func TestFontRolesResolveDistinctFamilies(t *testing.T) {
	if FontFor(TypeDisplay) == FontFor(TypeBody) { t.Fatal("display and body fonts must differ") }
	if FontFor(TypeData) == FontFor(TypeBody) { t.Fatal("data and body fonts must differ") }
}
