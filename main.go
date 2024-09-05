package main

import (
	"federicotorres233/mylinuxjourney/internal/utils"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// model holds the state of our application
type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

// initialize the model
func initialModel() model {
	return model{
		choices:  []string{"Start Journey", "Continue", "Progress", "Exit"},
		selected: make(map[int]struct{}),
	}
}

// Init can be used for initial commands you might want to run when the program starts
func (m model) Init() tea.Cmd {
	//log.Println("LinuxJourney started")
	return nil
}

// Update is called when messages are received. Here we handle user input.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", "l":
			if m.cursor == len(m.choices)-1 { // "Everything" is selected
				return m, tea.Quit
			} else {
				_, ok := m.selected[m.cursor]
				if ok {
					delete(m.selected, m.cursor)
				} else {
					m.selected[m.cursor] = struct{}{}
				}
			}
		}
	}
	return m, nil
}

// View renders the UI
func (m model) View() string {

	s := string(utils.LoadAsciiArt())

	for i, choice := range m.choices {
		// Choice is every option in the loop
		// Checked is a boolean that checks status
		// Cursor is cursor

		cursor := " "
		if m.cursor == i {
			cursor = "▸" // cursor
		}

		checked := " "
		if _, ok := m.selected[i]; ok || (i == len(m.choices)-1 && len(m.selected) == len(m.choices)-1) {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
