package util

import (
	"os"
	"os/exec"
)

func ClearScreen() {
	cmd := exec.Command("clear") // Use "clear" command
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func CheckForExitInput(input string, nextStep func()) {
	if input == "cancel" {
		nextStep()
	}
}
