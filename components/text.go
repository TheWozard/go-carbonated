package components

import (
	carbon "github.com/TheWozard/go-carbonated"
	tea "github.com/charmbracelet/bubbletea"
)

// Text is a simple viewable text component.
type Text string

func (t Text) ComponentUpdate(msg *carbon.Msg, cmd *carbon.Cmd) carbon.Component {
	if key, ok := msg.Get().(tea.KeyMsg); ok {
		switch key.String() {
		case "esc", "backspace", "enter", "space":
			msg.Consume()
			cmd.Pop()
		}
	}
	return t
}

func (t Text) View() string {
	return string(t)
}
