package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/NguyenHoaiPhuong/tokoin-test/golang/utils"
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

// SelectOption : read input text from terminal.
// If input = quit, return 0
// If input = 1, return 1
// If input = 2, return 2
// else, return 3
func SelectOption() int8 {
	fmt.Println("Type", aurora.Magenta("quit"), "to exit at any time, Press", aurora.Magenta("Enter"), "to continue")
	fmt.Println(aurora.BrightYellow("\tSelect search options:"))
	fmt.Println(aurora.BrightYellow("\t* Press 1 to search"))
	fmt.Println(aurora.BrightYellow("\t* Press 2 to view a list of searchable fields"))
	fmt.Println(aurora.BrightYellow("\t* Type quit to exit"))

	t := ReadInput()
	switch t {
	case "quit":
		return 0
	case "1":
		return 1
	case "2":
		return 2
	default:
		return 3
	}
}

// Search : print search results
func Search() {
	fmt.Println(aurora.BrightCyan("Select 1) Users or 2) Tickets or 3) Organizations"))
	sel := ReadInput()
	switch sel {
	case "1":
		SearchUser()
	case "2":
		SearchTicket()
	case "3":
		SearchOrganization()
	default:
		fmt.Println(aurora.Red("Selected option is invalid."))
	}
}

// EnterSearchTermAndValue : enter search term and value
func EnterSearchTermAndValue() (key, value string) {
	fmt.Println(aurora.BrightCyan("Enter search term"))
	key = ReadInput()

	fmt.Println(aurora.BrightCyan("Enter search value"))
	value = ReadInput()

	return key, value
}

// SearchUser : search user
func SearchUser() {
	key, value := EnterSearchTermAndValue()

	switch key {
	case "_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			utils.CheckError(err)
		}
		if user, ok := usersByID[id]; !ok {
			fmt.Println(aurora.Red("User with specific ID"), aurora.Red(id), aurora.Red("doesn't exist."))
		} else {
			utils.PrintPtrStructObject(user)
		}
	case "external_id":
		externalID := value
		if user, ok := usersByExternalID[externalID]; !ok {
			fmt.Println(aurora.Red("User with specific ExternalID"), aurora.Red(externalID), aurora.Red("doesn't exist."))
		} else {
			utils.PrintPtrStructObject(user)
		}
	default:
		fmt.Println(aurora.Red("Search term"), aurora.Red(key), aurora.Red("hasn't been supported yet"))
	}
}

// SearchTicket : search ticket
func SearchTicket() {

}

// SearchOrganization : search organization
func SearchOrganization() {

}

// SearchableFields : print searchable fields
func SearchableFields() {

}
