package util

import (
	"os"
	"os/exec"
	"strings"
)

func ClearScreen() {
	cmd := exec.Command("clear") // Use "clear" command
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func CheckForExitInput[T string | int](input T, nextStep func()) {
	switch strInput := any(input).(type) {
	case string:
		if strings.ToLower(strInput) == "cancel" {
			nextStep()
		}
	case int:
		if strInput == -1 {
			nextStep()
		}
	}
}
