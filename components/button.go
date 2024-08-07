package components

import (
	carbon "github.com/TheWozard/go-carbonated"
	tea "github.com/charmbracelet/bubbletea"
)

func NewButton(text string, confirm, reject carbon.Component) *Button {
	return &Button{
		Text:    text,
		Confirm: confirm,
		Reject:  reject,
	}
}

type Button struct {
	Text    string
	Confirm carbon.Component
	Reject  carbon.Component
}

func (b Button) ComponentUpdate(msg *carbon.Msg, cmd *carbon.Cmd) carbon.Component {
	if key, ok := msg.Get().(tea.KeyMsg); ok {
		switch key.String() {
		case "enter", " ":
			msg.Consume()
			cmd.Set(b.Confirm)
		case "esc", "backspace":
			msg.Consume()
			cmd.Set(b.Reject)
		}
	}
	return b
}

func (b Button) View() string {
	return b.Text
}
