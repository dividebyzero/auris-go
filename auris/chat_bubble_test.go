package auris

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestChatBubbleVariantsRender(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	cases := []ChatBubble{
		{Author: "Rin", Role: ChatAssistant, Message: "Ready.", Timestamp: "13:14:03"},
		{Author: "Zero", Role: ChatUser, Message: "Show me.", Timestamp: "13:14:28"},
		{Author: "Rin", Role: ChatAssistant, State: ChatSending, Timestamp: "13:15:22"},
		{Author: "Zero", Role: ChatUser, State: ChatError, Message: "Failed to send.", Timestamp: "13:15:40"},
	}
	for _, tc := range cases {
		if NewChatBubble(tc) == nil {
			t.Fatalf("chat bubble must render: %#v", tc)
		}
	}
}
