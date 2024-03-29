package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')
	HandleError(err)

	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimSuffix(str, "\r")

	return str
}
