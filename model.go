package carbon

import (
	tea "github.com/charmbracelet/bubbletea"
)

func NewModel(root Component) Model {
	return Model{
		Active: []Component{},
		Root:   root,
	}
}

type Model struct {
	Active []Component
	Root   Component

	// Size is stored in the model so when components are focused they can update their size.
	Size tea.WindowSizeMsg
}

func (m Model) Init() tea.Cmd {
	return FocusCmd
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m, act := m.PartialUpdate(msg, NewEmptyAction())
	return act.Act(m)
}

// PartialUpdate is a helper method that allows you to update the model with a message and an action, without
// committing to the action. This is useful when you want to update the model but still have more actions to collect.
func (m Model) PartialUpdate(msg tea.Msg, act Cmd) (Model, Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, act.Async(tea.Quit)
		case "esc":
			// Set an action, if the current component does not override it, it will be cleared.
			act = act.Clear()
		}
	case tea.WindowSizeMsg:
		// Size is stored in the model so when components are focused they can update their size.
		m.Size = msg
	case FocusMsg:
		// When we change focus, make sure the new component is focused.
		m, act = m.PartialUpdate(m.Size, act)
	}
	var c Component
	c, act = m.Current().Update(msg, act)
	return m.Set(c), act
}

func (m Model) View() string {
	return m.Current().View()
}

func (m Model) Current() Component {
	if len(m.Active) > 0 {
		return m.Active[len(m.Active)-1]
	}
	return m.Root
}

func (m Model) Set(c Component) Model {
	if len(m.Active) > 0 {
		m.Active[len(m.Active)-1] = c
	} else {
		m.Root = c
	}
	return m
}

func (m Model) Clear() Model {
	m.Active = []Component{}
	return m
}

func (m Model) Pop() Model {
	if len(m.Active) > 0 {
		m.Active = m.Active[:len(m.Active)-1]
	}
	return m
}

func (m Model) Push(c ...Component) Model {
	m.Active = append(m.Active, c...)
	return m
}
