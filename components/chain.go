package components

import (
	carbon "github.com/TheWozard/go-carbonated"
	tea "github.com/charmbracelet/bubbletea"
)

func NewChain(components ...carbon.Component) Chain {
	return Chain{
		Components: components,
	}
}

type Chain struct {
	Components []carbon.Component
}

func (c Chain) ComponentUpdate(msg *carbon.Msg, cmd *carbon.Cmd) carbon.Component {
	c.Components[0] = c.Components[0].ComponentUpdate(msg, cmd)
	if key, ok := msg.Get().(tea.KeyMsg); ok {
		switch key.String() {
		case "enter", " ":
			msg.Consume()
			c.Components = c.Components[1:]
			if len(c.Components) == 1 {
				cmd.Set(c.Components[0])
			} else {
				// Use Set to trigger focus change in the model.
				cmd.Set(c)
			}
		case "esc", "backspace":
			msg.Consume()
			cmd.Set(c.Components[len(c.Components)-1])
		}
	}
	return c
}

func (c Chain) View() string {
	return c.Components[0].View()
}
