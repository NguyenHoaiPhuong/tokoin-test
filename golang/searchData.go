package main

import (
	"fmt"

	"github.com/NguyenHoaiPhuong/tokoin-test/golang/utils"
	"github.com/logrusorgru/aurora"
)

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

// EnterSearchTermAndValue : enter search term and value
func EnterSearchTermAndValue() (key, value string) {
	fmt.Println(aurora.BrightCyan("Enter search term"))
	key = utils.ReadInput()

	fmt.Println(aurora.BrightCyan("Enter search value"))
	value = utils.ReadInput()

	return key, value
}

// SearchOrganization : search organization
func SearchOrganization() {

}

// SearchableFields : print searchable fields
func SearchableFields() {

}
