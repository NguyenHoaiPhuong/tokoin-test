package main

import (
	"fmt"

	"github.com/NguyenHoaiPhuong/tokoin-test/golang/config"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/utils"
	"github.com/logrusorgru/aurora"
)

var cfg *config.Configurations

func init() {
	cfg = config.SetupConfig("./resources", "default.toml")
}

func main() {
	LoadData()

	cont := true
	for cont {
		introduction()
		sel := utils.ReadInput()
		switch sel {
		case "quit":
			cont = false
		case "1":
			Search()
		case "2":
			SearchableFields()
		default:
			fmt.Println(aurora.Red("Selected option is invalid. Please try again."))
		}
	}
}

func introduction() {
	fmt.Println("Type", aurora.Magenta("quit"), "to exit at any time, Press", aurora.Magenta("Enter"), "to continue")
	fmt.Println(aurora.BrightYellow("\tSelect search options:"))
	fmt.Println(aurora.BrightYellow("\t* Press 1 to search"))
	fmt.Println(aurora.BrightYellow("\t* Press 2 to view a list of searchable fields"))
	fmt.Println(aurora.BrightYellow("\t* Type quit to exit"))
}

// LoadData : read json files and store data into heap
func LoadData() {
	LoadUserData()
	LoadTicketData()
	LoadOrganizationData()
}

// Search : print search results
func Search() {
	fmt.Println(aurora.BrightCyan("Select 1) Users or 2) Tickets or 3) Organizations"))
	sel := utils.ReadInput()
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

// SearchableFields : print searchable fields
func SearchableFields() {

}
