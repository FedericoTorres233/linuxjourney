package terminal

import (
	"fmt"
	"strings"
)

func DisplayMainOption(w int, cursor string, choice string) string {

	text := fmt.Sprintf("%s [ %s ] \n", cursor, choice)

	// Calculate padding needed to center the text
	chpad := CenterText(w, text)
	return chpad
}

func CenterText(w int, text string) string {
	// Calculate padding needed to center the text
	padding := strings.Repeat(" ", (w-len(text))/10)
	return padding + text
}
