package carbon

import tea "github.com/charmbracelet/bubbletea"

type FocusMsg struct{}

func FocusCmd() tea.Msg {
	return FocusMsg{}
}

type BlurMsg struct{}

func BlurCmd() tea.Msg {
	return BlurMsg{}
}
