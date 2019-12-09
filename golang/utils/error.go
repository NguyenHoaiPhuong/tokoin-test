package utils

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

// CheckError check error
func CheckError(err error) {
	if err != nil {
		fmt.Println(aurora.Red(err))
	}
}
