package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskConfirmation(message string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("🐾 %s \033[0;36m(Y/n)\033[0m: ", message)
	input, err := reader.ReadString('\n')
	if err != nil {
		return false
	}
	input = strings.TrimSpace(strings.ToLower(input))
	if input == "" || input == "y" || input == "yes" {
		return true
	}
	return false
}