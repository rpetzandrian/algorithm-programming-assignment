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

// func Max[T comparable](a, b T) T {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func CheckForExitInput[T any](input T, nextStep func()) {
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
