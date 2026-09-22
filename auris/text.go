package auris

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// Text creates an Auris text object with the correct embedded family.
// This is the preferred primitive for custom Auris widgets.
func Text(value string, role TypeRole, size float32, c color.Color) *canvas.Text {
	t := canvas.NewText(value, c)
	t.TextSize = size
	t.TextStyle = fyne.TextStyle{}
	t.TextFont = FontFor(role)
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
