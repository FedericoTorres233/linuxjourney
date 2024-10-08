package game

import (
	"federicotorres233/mylinuxjourney/internal/markdown"
	"federicotorres233/mylinuxjourney/internal/utils"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	startchapter int = 1
	startpage    int = 1
	startScreen  int = iota
	gameScreen
)

// model holds the state of our application
type model struct {
	chapter  int
	page     int
	state    int
	choices  []string
	cursor   int
	selected map[int]struct{}
}

// initialize the model
func InitialModel() model {
	return model{
		chapter:  startchapter,
		page:     startpage,
		state:    startScreen,
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
			if m.state == startScreen {
				return m, tea.Quit
			} else {
				m.state = startScreen
			}
		case "h", "left":
			if m.page <= 1 {
				m.state = startScreen
				m.page = 1
				break
			}
			m.page--
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", "l", "right":
			if m.state == startScreen {

				switch m.cursor {
				case 0:
					// start selected
					m.state = gameScreen
					return m, nil
				case len(m.choices) - 1:
					// Exit selected
					return m, tea.Quit
				}

				_, ok := m.selected[m.cursor]
				if ok {
					delete(m.selected, m.cursor)
				} else {
					m.selected[m.cursor] = struct{}{}
				}

			} else {
				// Outside start screen
				m.page++
			}
		}
	}
	return m, nil
}

// View renders the UI
func (m model) View() string {

	s := string(utils.LoadAsciiArt())

	/* Get the terminal width
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 0 // Default width if we can't get the terminal size
	}*/

	chapters := make(map[int]map[int]string)

	for i := 1; i <= 5; i++ {
		chapters[i] = make(map[int]string)
		for f := 1; f <= 4; f++ {
			chapters[i][f] = fmt.Sprintf("public/pages/chapter%d/%d.md", i, f)
		}
	}

	switch m.state {
	case startScreen:
		for i, choice := range m.choices {
			// Choice is every option in the loop
			// Checked is a boolean that checks status
			// Cursor is cursor

			cursor := " "
			if m.cursor == i {
				cursor = "▸" // cursor
			}

			// Don't center for now
			//s += terminal.DisplayMainOption(width, cursor, choice)
			s += fmt.Sprintf("%s { %s } \n", cursor, choice)

		}
		return s
	case gameScreen:

		//utils.ClearScreen()
		padding := "\n\n\n"
		_, err := markdown.ParseMarkdown(chapters[m.chapter][m.page])
		if err != nil {
			m.page = 1
			m.chapter++
		}
		txt, _ := markdown.ParseMarkdown(chapters[m.chapter][m.page])
		return utils.FormatOutput(padding + txt)
	}
	return ""

}
