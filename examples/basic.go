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
		components.Text("Root"),
		components.Text("Active 1"),
		components.Text("Active 2"),
		components.Text("Active 3"),
		components.Text("Active 4"),
	))

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
