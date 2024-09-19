package utils

import "fmt"

func PrintInColor(text string, color int) {
	fmt.Printf("\x1b[38;5;%dm%s\x1b[0m\n", color, text)
}
