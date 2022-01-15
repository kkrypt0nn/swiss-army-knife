package terminal

const (
	BOLD          = "\033[1m"
	DIM           = "\033[2m"
	UNDERLINE     = "\033[4m"
	BLINK         = "\033[5m"
	INVERSE       = "\033[7m"
	STRIKETHROUGH = "\033[9m"
)

// Bold returns a string with the effect set to bold.
func Bold(message string) string {
	return BOLD + message + RESET
}

// Dim returns a string with the effect set to dim.
func Dim(message string) string {
	return DIM + message + RESET
}

// Underline returns a string with the effect set to underline.
func Underline(message string) string {
	return UNDERLINE + message + RESET
}

// Blink returns a string with the effect set to blink.
func Blink(message string) string {
	return BLINK + message + RESET
}

// Inverse returns a string with the effect set to inverse.
func Inverse(message string) string {
	return INVERSE + message + RESET
}

// Strikethrough returns a string with the effect set to strikethrough.
func Strikethrough(message string) string {
	return STRIKETHROUGH + message + RESET
}
