package auris

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// Text creates an Auris text object whose TextStyle resolves through Theme.Font
// to the correct embedded family.
func Text(value string, role TypeRole, size float32, c color.Color) *canvas.Text {
	t := canvas.NewText(value, c)
	t.TextSize = size
	switch role {
	case TypeDisplay:
		t.TextStyle = fyne.TextStyle{Bold:true}
	case TypeData:
		t.TextStyle = fyne.TextStyle{Monospace:true}
	default:
		t.TextStyle = fyne.TextStyle{}
	}
	return t
}

func DisplayText(value string, size float32, c color.Color) *canvas.Text {
	return Text(value, TypeDisplay, size, c)
}
func BodyText(value string, size float32, c color.Color) *canvas.Text {
	return Text(value, TypeBody, size, c)
}
func DataText(value string, size float32, c color.Color) *canvas.Text {
	return Text(value, TypeData, size, c)
}
