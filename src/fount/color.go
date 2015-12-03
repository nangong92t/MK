package fount

import (
	"fmt"
)

const (
	Black = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func Color(str,color string) string {
	switch color {
		case "black":
			return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", Black, str)
		case "red":
			return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", Red, str)
		case "green":
			return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", Green, str)
		case "yellow":
			return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", Yellow, str)
		case "blue":
			return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", Blue, str)
		case "magenta":
			return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", Magenta, str)
		case "cyan":
			return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", Cyan, str)
		case "white":
			return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", White, str)
		default:
			return str
	}
}