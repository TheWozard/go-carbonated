package carbon

import tea "github.com/charmbracelet/bubbletea"

// FocusMsg is a custom tea.Msg that is sent when a component is focused.
type FocusMsg struct {
	Size tea.WindowSizeMsg
}

// BlurMsg is a custom tea.Msg that blurs the active component.
type BlurMsg struct{}

type Msg struct {
	msg      tea.Msg
	consumed bool
}

// Get returns the message if it has not been consumed.
func (m *Msg) Get() tea.Msg {
	if m.consumed {
		return nil
	}
	return m.msg
}

// Consume marks the message as consumed.
func (m *Msg) Consume() {
	m.consumed = true
}

// Message returns the message even if it has been consumed.
func (m *Msg) Message() tea.Msg {
	return m.msg
}
