package carbon

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Cmd is a collection of actions to be taken either asynchronously or against the model.
type Cmd struct {
	// transition
	transition func(Model) Model
	async      []tea.Cmd
}

// Act will execute the actions against the model and provide any asynchronous commands to be executed.
func (c *Cmd) Act(m Model) (Model, tea.Cmd) {
	if c.transition != nil {
		// Swap the focus
		m = m.componentUpdate(&Msg{msg: BlurMsg{}}, c)
		m = c.transition(m)
		m = m.componentUpdate(&Msg{msg: FocusMsg{}}, c)
	}
	return m, tea.Batch(c.async...)
}

// Async will append asynchronous commands to the current set.
func (a *Cmd) Async(cmds ...tea.Cmd) {
	a.async = append(a.async, cmds...)
}

// Clear will clear the current models active stack back to the root.
func (a *Cmd) Clear() {
	a.transition = Model.Clear
}

// Pop will pop the current active component off the stack.
func (a *Cmd) Pop() {
	a.transition = Model.Pop
}

// Push will push a new component onto the active stack.
func (a *Cmd) Push(c ...Component) {
	a.transition = func(m Model) Model {
		return m.Push(c...)
	}
}
