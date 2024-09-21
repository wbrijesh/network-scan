package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ConfirmBeforeRunning(prompt string, fn func()) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s (Y/n): ", prompt)
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		response = strings.TrimSpace(strings.ToLower(response))

		switch response {
		case "y", "yes", "":
			fn()
			return
		case "n", "no":
			fmt.Println("Operation cancelled.")
			return
		default:
			fmt.Println("Invalid response. Please answer with 'y' or 'n'.")
		}
	}
}
