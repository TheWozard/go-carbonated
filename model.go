package carbon

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// NewModel creates a new Model with the given root component and active components.
func NewModel(root Component, active ...Component) Model {
	return Model{
		Active: active,
		Root:   root,
	}
}

// Model is the main model for the application. It maintains the current focus and state of the application.
type Model struct {
	// Active is a stack of components that are currently active.
	// The last component in the stack is the current component.
	Active []Component
	// Root is the main component of the application, this component cannot be removed.
	Root Component

	// Size is stored in the model so when components are focused they can update their size.
	Size tea.WindowSizeMsg
}

func (m Model) Init() tea.Cmd {
	// On start we need to focus the current active component.
	return func() tea.Msg { return FocusMsg{} }
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := &Cmd{}
	return cmd.Act(m.componentUpdate(&Msg{msg: msg}, cmd))
}

// ComponentUpdate allows the model to be used as a component itself.
func (m Model) ComponentUpdate(msg *Msg, cmd *Cmd) Component {
	model, _ := cmd.Act(m.componentUpdate(msg, cmd))
	cmd.Reset()
	return model
}

// Internal implementation of ComponentUpdate that returns a Model instead of a Component.
func (m Model) componentUpdate(msg *Msg, cmd *Cmd) Model {
	switch typed := msg.Get().(type) {
	case tea.KeyMsg:
		if typed.String() == "ctrl+c" {
			// To preserve all output, we write a newline before exiting as the terminal may be in a weird state.
			os.Stdout.WriteString("\n")
			cmd.Async(tea.Quit)
			return m
		}
	case tea.WindowSizeMsg:
		// Size is stored in the model so when components are focused they can update their size.
		m.Size = typed
	case FocusMsg:
		// Update focus message with the current size if it is missing or invalid.
		if typed.Size.Width != m.Size.Width || typed.Size.Height != m.Size.Height {
			msg = &Msg{msg: FocusMsg{Size: m.Size}}
		}
	}
	c := m.Current().ComponentUpdate(msg, cmd)
	// If the message was not consumed, see if a default can handle it.
	switch typed := msg.Get().(type) {
	case tea.KeyMsg:
		if typed.String() == "esc" {
			// We could pop, but we don't know how the stack is being used.
			cmd.Clear()
			msg.Consume()
		}
	}
	return m.Set(c)
}

func (m Model) View() string {
	return m.Current().View()
}

// Returns the current component in the stack.
func (m Model) Current() Component {
	if len(m.Active) > 0 {
		return m.Active[len(m.Active)-1]
	}
	return m.Root
}

// Sets the current active component in the stack.
func (m Model) Set(c Component) Model {
	if len(m.Active) > 0 {
		m.Active[len(m.Active)-1] = c
	} else {
		m.Root = c
	}
	return m
}

// Clear will clear the current models active stack back to the root.
func (m Model) Clear() Model {
	m.Active = []Component{}
	return m
}

// Pop will pop the current active component off the stack.
func (m Model) Pop() Model {
	if len(m.Active) > 0 {
		m.Active = m.Active[:len(m.Active)-1]
	}
	return m
}

// Push will push a new component onto the active stack.
func (m Model) Push(c ...Component) Model {
	m.Active = append(m.Active, c...)
	return m
}
