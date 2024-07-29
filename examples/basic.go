package main

import (
	"fmt"
	"os"

	carbon "github.com/TheWozard/go-carbonated"
	"github.com/TheWozard/go-carbonated/components"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(carbon.NewModel(
		components.NewBatch(1,
			components.Text("Rocket Launch Simulation"),
			carbon.NewModel(
				components.Text("Blastoff 🚀"),
				components.Text("1"),
				components.Text("2"),
				components.Text("3"),
				components.Text("4"),
				components.Text("5"),
				components.Text("Press Space to launch!"),
			),
		),
	))

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
