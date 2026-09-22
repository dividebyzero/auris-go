package auris

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// ChatRole controls the alignment/accent language of a chat bubble.
type ChatRole int

const (
	ChatAssistant ChatRole = iota
	ChatUser
)

// ChatState controls transient delivery presentation.
type ChatState int

const (
	ChatReady ChatState = iota
	ChatSending
	ChatError
)

// ChatBubble describes one Auris-styled conversation message.
type ChatBubble struct {
	Author    string
	Role      ChatRole
	State     ChatState
	Message   string
	Timestamp string
	Width     float32
}

// NewChatBubble renders a chamfered Auris conversation cell. Assistant bubbles
// use the warm primary ramp; user bubbles use the cool slate secondary accent.
func NewChatBubble(spec ChatBubble) fyne.CanvasObject {
	s := CurrentScheme()
	width := spec.Width
	if width <= 0 {
		width = 360
	}

	accent := s.PrimaryActive
	role := "ASSISTANT"
	if spec.Role == ChatUser {
		accent = s.Secondary
		role = "USER"
	}
	if spec.State == ChatError {
		accent = s.Danger
	}

	author := strings.ToUpper(spec.Author)
	if author == "" {
		author = role
	}
	header := DataText(author+"  //  "+role, 11, accent)

	message := spec.Message
	if spec.State == ChatSending && message == "" {
		message = "•  •  •   Sending…"
	}
	body := canvas.NewText(message, s.TextBright)
	body.TextSize = 14
	body.TextStyle = fyne.TextStyle{}

	footerText := spec.Timestamp
	if spec.State == ChatError {
		if footerText != "" {
			footerText += "   //   FAILED"
		} else {
			footerText = "FAILED"
		}
	}
	footer := DataText(footerText, 10, s.TextMid)

	inside := container.NewVBox(header, body)
	if footerText != "" {
		inside.Add(footer)
	}
	padded := container.NewPadded(inside)

	height := padded.MinSize().Height + 8
	if height < 72 {
		height = 72
	}
	return NewDepthContainer(padded, fyne.NewSize(width, height), s.Bevel.MD, s.SurfaceInset, accent, DepthSubtle)
}
