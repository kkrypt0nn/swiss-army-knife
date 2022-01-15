# Terminal

This package provides some helper functions to be able to send formatted output to the terminal. It uses the ANSI escape codes to color the output.

## Example

```go
func main() {
    if terminal.AreColorsSupported() {
        fmt.Println(terminal.BackgroundBlack(terminal.Green("Terminal supports colors")))
    } else {
        fmt.Println("Terminal does not support colors")
    }
}
```
