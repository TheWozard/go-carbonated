package components

import (
	carbon "github.com/TheWozard/go-carbonated"
	tea "github.com/charmbracelet/bubbletea"
)

func NewButton(text string, components ...carbon.Component) Button {
	return Button{
		Text:       text,
		Components: components,
	}
}

type Button struct {
	Text       string
	Components []carbon.Component
}

func (b Button) ComponentUpdate(msg *carbon.Msg, cmd *carbon.Cmd) carbon.Component {
	if key, ok := msg.Get().(tea.KeyMsg); ok {
		switch key.String() {
		case "enter", " ":
			msg.Consume()
			cmd.Push(b.Components...)
		}
	}
	return b
}

func (b Button) View() string {
	return b.Text
}
