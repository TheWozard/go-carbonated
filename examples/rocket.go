package main

import (
	"fmt"
	"os"

	carbon "github.com/TheWozard/go-carbonated"
	"github.com/TheWozard/go-carbonated/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewFocusCounter(content carbon.Component, count *int) carbon.Component {
	return components.Watcher{
		Content: content,
		Watch: func(msg *carbon.Msg) {
			if _, ok := msg.Get().(carbon.FocusMsg); ok {
				*count++
			}
		},
	}
}

func main() {
	// Shared state should be allocated outside of the model, and pointed to by components that use it.
	attempts := 0
	successes := 0
	p := tea.NewProgram(carbon.NewModel(
		components.NewWrapper(1,
			components.NewStyled(components.NewDynamicText(func() string {
				return fmt.Sprintf("Space Launch Simulator - Successes: %d", successes)
			}), lipgloss.NewStyle().Foreground(lipgloss.Color("#AAAAAA"))),
			carbon.NewModel(
				components.NewButton("Press Space/Enter to launch!",
					components.NewStyled(NewFocusCounter(components.Text("Blastoff 🚀"), &successes), lipgloss.NewStyle().Bold(true)),
					components.Text("1"),
					components.Text("2"),
					components.Text("3"),
					components.Text("4"),
					components.Text("5"),
					components.NewStyled(NewFocusCounter(components.DynamicText(func() string {
						return fmt.Sprintf("Launch %d!", attempts)
					}), &attempts), lipgloss.NewStyle().Bold(true)),
				),
			),
		),
	))

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
