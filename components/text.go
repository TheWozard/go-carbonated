package components

import (
	carbon "github.com/TheWozard/go-carbonated"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Text is a simple viewable text component.
type Text string

func (t Text) ComponentUpdate(msg *carbon.Msg, cmd *carbon.Cmd) carbon.Component {
	if key, ok := msg.Get().(tea.KeyMsg); ok {
		switch key.String() {
		case "esc", "backspace", "enter", " ":
			msg.Consume()
			cmd.Pop()
		}
	}
	return t
}

func (t Text) View() string {
	return string(t)
}

func NewStyledText(text string, style lipgloss.Style) StyledText {
	return StyledText{
		Text:  text,
		Style: style,
	}
}

// StyledText is a viewable text component with a style. Automatically resizes to fit the width of the parent component.
type StyledText struct {
	Text  string
	Style lipgloss.Style
}

func (t StyledText) ComponentUpdate(msg *carbon.Msg, cmd *carbon.Cmd) carbon.Component {
	switch typed := msg.Get().(type) {
	case tea.KeyMsg:
		switch typed.String() {
		case "esc", "backspace", "enter", " ":
			msg.Consume()
			cmd.Pop()
		}
	case carbon.FocusMsg:
		t.Style = t.Style.Width(typed.Size.Width)
	case tea.WindowSizeMsg:
		t.Style = t.Style.Width(typed.Width)
	}
	return t
}

func (t StyledText) View() string {
	return t.Style.Render(t.Text)
}
