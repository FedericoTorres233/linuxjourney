package game

type Model struct {
	Page     int
	State    int
	Choices  []string
	Cursor   int
	Selected map[int]struct{}
}

func StartJourney() string {
	text := "\n\n\n\n\n\n\n\n\n"
	text += "Welcome to your journey through linux!\n\n\n"

	text += "Press '→' to continue through your journey.\n"
	text += "Press '<-' to go back to the previous section\n"
	text += "Press 'q' to go to main menu\n"
	return text
}
