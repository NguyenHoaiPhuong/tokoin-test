package utils

import (
	"bufio"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

// ReadInput : read input from terminal
func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from input: ", err)
	}
	return t
}

// EnterSearchTermAndValue : enter search term and value
func EnterSearchTermAndValue() (key, value string) {
	fmt.Println(aurora.BrightCyan("Enter search term"))
	key = ReadInput()

	fmt.Println(aurora.BrightCyan("Enter search value"))
	value = ReadInput()

	fmt.Println()

	return key, value
}
