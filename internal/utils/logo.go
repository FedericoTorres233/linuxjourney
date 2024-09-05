package utils

import (
	"os"
	"path/filepath"
)

func LoadAsciiArt() string {
	ascii_art, _ := os.ReadFile(filepath.Join(".", "public", "ascii.txt"))
	return "\n\n" + string(ascii_art) + "\n\n"
}
