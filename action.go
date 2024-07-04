package carbon

import tea "github.com/charmbracelet/bubbletea"

// NewEmptyAction returns an empty action.
func NewEmptyAction() Cmd {
	return Cmd{}
}

// Transition is a function that will modify the current focus of the application.
type Transition func(Model) Model

// Swap will swap the current focus of the application with the new focus after the transition.
func (t Transition) Swap(m Model) (Model, Cmd) {
	act := NewEmptyAction()
	m, act = m.PartialUpdate(BlurMsg{}, act)
	m = m.Clear()
	return m.PartialUpdate(FocusMsg{}, act)
}

// Cmd is a collection of actions to be taken either asynchronously or against the model.
type Cmd struct {
	transition Transition
	async      []tea.Cmd
}

// Act will execute the actions against the model and provide any asynchronous commands to be executed.
func (a Cmd) Act(m Model) (Model, tea.Cmd) {
	transition := a.transition
	cmds := a.async
	for transition != nil {
		m, a = transition.Swap(m)
		transition = a.transition
		cmds = append(cmds, a.async...)
	}
	return m, tea.Batch(cmds...)
}

// Async will append asynchronous commands to the current set.
func (a Cmd) Async(cmds ...tea.Cmd) Cmd {
	a.async = append(a.async, cmds...)
	return a
}

// Clear will clear the current models active stack back to the root.
func (a Cmd) Clear() Cmd {
	a.transition = Model.Clear
	return a
}

// Pop will pop the current active component off the stack.
func (a Cmd) Pop() Cmd {
	a.transition = Model.Pop
	return a
}
