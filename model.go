package carbon

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// NewModel creates a new model with the given active component.
func NewModel(active Component) Model {
	return Model{Active: active}
}

// Model is the main model for the application. It handles focus changes and updates to the current component.
// Provides default input handling for the application.
type Model struct {
	Active Component

	// Size is stored in the model so when components are focused they can update their size.
	Size tea.WindowSizeMsg
}

func (m Model) Init() tea.Cmd {
	// On start we need to focus the current active component.
	return func() tea.Msg { return FocusMsg{} }
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := &Cmd{}
	return cmd.Transition(m.componentUpdate(&Msg{msg: msg}, cmd)), cmd.Cmd()
}

// ComponentUpdate allows the model to be used as a component itself.
func (m Model) ComponentUpdate(msg *Msg, cmd *Cmd) Component {
	return cmd.Transition(m.componentUpdate(msg, cmd))
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
	return Model{Active: m.Active.ComponentUpdate(msg, cmd)}
}

func (m Model) View() string {
	return m.Active.View()
}
