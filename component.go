package carbon

import tea "github.com/charmbracelet/bubbletea"

// Component is an interface that represents a view that can be updated and rendered.
type Component interface {
	Update(tea.Msg, Cmd) (Component, Cmd)
	View() string
}
