package terminal

import "os"

const RESET = "\033[0m"

func AreColorsSupported() bool {
	if term := os.Getenv("TERM"); term == "" {
		return false
	}
	return true
}
