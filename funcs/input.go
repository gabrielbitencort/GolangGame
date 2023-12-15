package funcs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func Input(prompt string) string {
	return input(prompt)
}
