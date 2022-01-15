package terminal

const (
	BACK_BLACK     = "\033[40m"
	BACK_RED       = "\033[41m"
	BACK_GREEN     = "\033[42m"
	BACK_YELLOW    = "\033[43m"
	BACK_BLUE      = "\033[44m"
	BACK_PURPLE    = "\033[45m"
	BACK_CYAN      = "\033[46m"
	BACK_WHITE     = "\033[47m"
	BACK_DARKGRAY  = "\033[100m"
	BACK_LIGHTBLUE = "\033[104m"
)

// BackgroundBlack returns a string with the background color set to black.
func BackgroundBlack(message string) string {
	return BACK_BLACK + message + RESET
}

// BackgroundRed returns a string with the background color set to red.
func BackgroundRed(message string) string {
	return BACK_RED + message + RESET
}

// BackgroundGreen returns a string with the background color set to green.
func BackgroundGreen(message string) string {
	return BACK_GREEN + message + RESET
}

// BackgroundYellow returns a string with the background color set to yellow.
func BackgroundYellow(message string) string {
	return BACK_YELLOW + message + RESET
}

// BackgroundBlue returns a string with the background color set to blue.
func BackgroundBlue(message string) string {
	return BACK_BLUE + message + RESET
}

// BackgroundPurple returns a string with the background color set to purple.
func BackgroundPurple(message string) string {
	return BACK_PURPLE + message + RESET
}

// BackgroundCyan returns a string with the background color set to cyan.
func BackgroundCyan(message string) string {
	return BACK_CYAN + message + RESET
}

// BackgroundWhite returns a string with the background color set to white.
func BackgroundWhite(message string) string {
	return BACK_WHITE + message + RESET
}

// BackgroundLightBlue returns a string with the background color set to light blue.
func BackgroundLightBlue(message string) string {
	return BACK_LIGHTBLUE + message + RESET
}

// BackgroundDarkGray returns a string with the background color set to dark gray.
func BackgroundDarkGray(message string) string {
	return BACK_DARKGRAY + message + RESET
}
