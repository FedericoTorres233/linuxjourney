package main

import (
	"federicotorres233/mylinuxjourney/internal/terminal"
	"federicotorres233/mylinuxjourney/internal/utils"
	"fmt"
	"os"

	"golang.org/x/term"
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
		choices:  []string{"Start", "Continue", "Progress", "Exit"},
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
			if m.cursor != len(m.choices)-1 { // Exit selected
				_, ok := m.selected[m.cursor]
				if ok {
					// 
					delete(m.selected, m.cursor)
				} else {
					m.selected[m.cursor] = struct{}{}
				}
				break
			}

			return m, tea.Quit
		}
	}
	return m, nil
}

// View renders the UI
func (m model) View() string {

	s := string(utils.LoadAsciiArt())

	// Get the terminal width
    width, _, err := term.GetSize(int(os.Stdout.Fd()))
    if err != nil {
        width = 0 // Default width if we can't get the terminal size
    }

	for i, choice := range m.choices {
		// Choice is every option in the loop
		// Checked is a boolean that checks status
		// Cursor is cursor

		cursor := " "
		if m.cursor == i {
			cursor = "▸" // cursor
		}

		s += terminal.DisplayMainOption(width, cursor, choice)
		
	}
	return s
}

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
