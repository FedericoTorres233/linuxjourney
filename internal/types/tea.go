package types

// model holds the state of our application
type Model struct {
	Choices  []string
	Cursor   int
	Selected map[int]struct{}
}