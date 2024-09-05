package markdown

import (
	"os"
)

func ParseMarkdown() (string, error) {

	mdContent, err := os.ReadFile("public/pages/start.md")
	if err != nil {
		return "", err
	}
	
	return string(mdContent), nil
}
