package utils

import (
	"bufio"
	"fmt"
	"os"
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
