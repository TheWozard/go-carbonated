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

// Update allies the transition and consumes it
func (c *Cmd) Transition(m Model) Model {
	if c.transition != nil {
		m = c.transition(m)
		c.Reset()
		return m
	}
	return m
}

// Cmds returns the combined set of commands for the model.
func (c *Cmd) Cmd() tea.Cmd {
	if c.async == nil {
		return nil
	}
	return tea.Batch(c.async...)
}

// Async will append asynchronous commands to the current set.
func (a *Cmd) Async(cmds ...tea.Cmd) {
	a.async = append(a.async, cmds...)
}

// Set sets the component to transition to. Ignores passed nils.
// To clear call Reset.
func (a *Cmd) Set(active Component) {
	if active == nil {
		return
	}
	a.transition = func(m Model) Model {
		m = m.componentUpdate(&Msg{msg: BlurMsg{}}, a)
		m.Active = active
		return m.componentUpdate(&Msg{msg: FocusMsg{}}, a)
	}
}

// Reset will reset the current transition.
func (a *Cmd) Reset() {
	a.transition = nil
}
