package components

import (
	"strings"

	carbon "github.com/TheWozard/go-carbonated"
	tea "github.com/charmbracelet/bubbletea"
)

func NewWrapper(input int, contents ...carbon.Component) Wrapper {
	return Wrapper{
		Contents: contents,
		Input:    input,
	}
}

// Wrapper wraps a slice of components into a single component.
// Only the component at inputs receives input events.
// All other events are shared.
type Wrapper struct {
	Contents []carbon.Component
	Input    int
}

func (b Wrapper) ComponentUpdate(msg *carbon.Msg, cmd *carbon.Cmd) carbon.Component {
	switch msg.Get().(type) {
	case tea.KeyMsg:
		b.Contents[b.Input] = b.Contents[b.Input].ComponentUpdate(msg, cmd)
	default:
		for i := range b.Contents {
			b.Contents[i] = b.Contents[i].ComponentUpdate(msg, cmd)
		}
	}
	return b
}

func (b Wrapper) View() string {
	var s strings.Builder
	for i := range b.Contents {
		if i > 0 {
			s.WriteString("\n")
		}
		s.WriteString(b.Contents[i].View())
	}
	return s.String()
}
