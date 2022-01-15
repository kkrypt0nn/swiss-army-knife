package terminal

const (
	BLACK  = "\033[30m"
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	BLUE   = "\033[34m"
	PURPLE = "\033[35m"
	CYAN   = "\033[36m"
	WHITE  = "\033[37m"
)

// Black returns a string with the foreground color set to black.
func Black(message string) string {
	return BLACK + message + RESET
}

// Red returns a string with the foreground color set to red.
func Red(message string) string {
	return RED + message + RESET
}

// Green returns a string with the foreground color set to green.
func Green(message string) string {
	return GREEN + message + RESET
}

// Yellow returns a string with the foreground color set to yellow.
func Yellow(message string) string {
	return YELLOW + message + RESET
}

// Blue returns a string with the foreground color set to blue.
func Blue(message string) string {
	return BLUE + message + RESET
}

// Purple returns a string with the foreground color set to purple.
func Purple(message string) string {
	return PURPLE + message + RESET
}

// Cyan returns a string with the foreground color set to cyan.
func Cyan(message string) string {
	return CYAN + message + RESET
}

// White returns a string with the foreground color set to white.
func White(message string) string {
	return WHITE + message + RESET
}
