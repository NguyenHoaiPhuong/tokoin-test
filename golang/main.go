package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/NguyenHoaiPhuong/tokoin-test/golang/config"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/jsonfunc"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/model"
	"github.com/logrusorgru/aurora"
)

var cfg *config.Configurations
var users model.Users
var tickets model.Tickets
var organizations model.Organizations

func init() {
	cfg = config.SetupConfig("./resources", "default.toml")
}

func main() {
	// For testing configurations
	// cfg.Print()

	LoadData()

	cont := true
	for cont {
		sel := ReadInput()
		switch sel {
		case 0:
			cont = false
		case 1:
			fmt.Println("1111")
		case 2:
			fmt.Println("2222")
		case 3:
			fmt.Println(aurora.Red("Selected option is invalid. Please type again."))
		}
	}
}

// ReadInput : read input text from terminal.
// If input = quit, return 0
// If input = 1, return 1
// If input = 2, return 2
// else, return 3
func ReadInput() int8 {
	fmt.Println("Type", aurora.Magenta("quit"), "to exit at any time, Press", aurora.Magenta("Enter"), "to continue")
	fmt.Println(aurora.BrightYellow("\tSelect search options:"))
	fmt.Println(aurora.BrightYellow("\t* Press 1 to search"))
	fmt.Println(aurora.BrightYellow("\t* Press 2 to view a list of searchable fields"))
	fmt.Println(aurora.BrightYellow("\t* Type quit to exit"))

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from input: ", err)
	}
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

}

// SearchableFields : print searchable fields
func SearchableFields() {

}

// LoadData : read json files and store data into heap
func LoadData() {
	err := jsonfunc.ReadFromFile(cfg.URLs.UserURL, &users)
	CheckError(err)
	// fmt.Println(users[0].Email)

	err = jsonfunc.ReadFromFile(cfg.URLs.TicketURL, &tickets)
	CheckError(err)
	// fmt.Println(tickets[0].CreatedAt)

	err = jsonfunc.ReadFromFile(cfg.URLs.OrganizationURL, &organizations)
	CheckError(err)
	// fmt.Println(organizations[0].Details)
}

// CheckError check error
func CheckError(err error) {
	if err != nil {
		fmt.Println(aurora.Red(err))
	}
}
