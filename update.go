package main

import tea "github.com/charmbracelet/bubbletea"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "up":
			m.handleCleanStageUpKey()
			m.handleUpKey()
		case "left":
			m.handleCleanStageLeftKey()
		case "right":
			m.handleCleanStageRightKey()
		case "down":
			m.handleCleanStageDownKey()
			m.handleDownKey()
		case "enter":
			m.handleCleanStageEnter() //TODO: order matters here, which seems like a bad idea
			m.handleEnter()
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd

}
