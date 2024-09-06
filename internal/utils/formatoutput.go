package utils

import (
	"strings"
)

const width = 100

func FormatOutput(text string) string {
	var result strings.Builder
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		words := strings.Fields(line)
		lineBuilder := strings.Builder{}
		for _, word := range words {
			if lineBuilder.Len()+len(word) > width {
				if lineBuilder.Len() > 0 {
					result.WriteString(lineBuilder.String())
					result.WriteString("\n")
					lineBuilder.Reset()
				}
				if len(word) > width {
					// If the word itself is longer than width, break it
					for i := 0; i < len(word); i += width {
						result.WriteString(word[i:min(i+width, len(word))])
						result.WriteString("\n")
					}
				} else {
					lineBuilder.WriteString(word)
					lineBuilder.WriteRune(' ')
				}
			} else {
				lineBuilder.WriteString(word)
				lineBuilder.WriteRune(' ')
			}
		}
		if lineBuilder.Len() > 0 {
			result.WriteString(strings.TrimSpace(lineBuilder.String()))
			result.WriteString("\n")
		}
	}
	return result.String()
}
