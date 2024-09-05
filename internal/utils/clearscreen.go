package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	var cmd *exec.Cmd

	// Determine the command based on the operating system
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else { // Unix-like systems, including Linux and macOS
		cmd = exec.Command("clear")
	}

	// Set the output to the terminal
	cmd.Stdout = os.Stdout
	cmd.Run()
}
