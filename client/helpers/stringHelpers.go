package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TakeInput(inputPromptText string) (input string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(inputPromptText)
	input, err = reader.ReadString('\n')
	return
}

func NormalizeString(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}