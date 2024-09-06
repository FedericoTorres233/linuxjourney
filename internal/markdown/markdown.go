package markdown

import (
	"os"
)

func ParseMarkdown(f string) (string, error) {

	mdContent, err := os.ReadFile(f)
	if err != nil {
		return "", err
	}

	return string(mdContent), nil
}
