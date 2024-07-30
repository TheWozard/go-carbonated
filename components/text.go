package components

import (
	carbon "github.com/TheWozard/go-carbonated"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// / NewText creates a new text component.
func NewText(text string) Text {
	return Text(text)
}

// Text is a simple viewable text component.
type Text string

func (t Text) ComponentUpdate(msg *carbon.Msg, cmd *carbon.Cmd) carbon.Component {
	return t
}

func (t Text) View() string {
	return string(t)
}

// NewDynamicText creates a new dynamic text component.
func NewDynamicText(get func() string) DynamicText {
	return DynamicText(get)
}

// DynamicText is a text component that can change its text.
type DynamicText func() string

func (d DynamicText) ComponentUpdate(msg *carbon.Msg, cmd *carbon.Cmd) carbon.Component {
	return d
}

func (d DynamicText) View() string {
	return d()
}

// NewStyled creates a new styled wrapping component that styles the wrapped view.
func NewStyled(contents carbon.Component, style lipgloss.Style) Styled {
	return Styled{
		Contents: contents,
		Style:    style,
	}
}

// Styled is a component that wraps another component and applies a style to it.
type Styled struct {
	Contents carbon.Component
	Style    lipgloss.Style
}

func (s Styled) ComponentUpdate(msg *carbon.Msg, cmd *carbon.Cmd) carbon.Component {
	switch typed := msg.Get().(type) {
	case carbon.FocusMsg:
		s.Style = s.Style.Width(typed.Size.Width)
	case tea.WindowSizeMsg:
		s.Style = s.Style.Width(typed.Width)
	}
	s.Contents = s.Contents.ComponentUpdate(msg, cmd)
	return s
}

func (s Styled) View() string {
	return s.Style.Render(s.Contents.View())
}
