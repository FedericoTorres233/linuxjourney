package view

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// View renders the UI
func (m tea.model) View() string {
	s := "\nWhat would you like to see?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = "▸" // cursor!
		}

		checked := " "
		if _, ok := m.selected[i]; ok || (i == len(m.choices)-1 && len(m.selected) == len(m.choices)-1) {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	return s
}