package main

import (
	"fmt"
	"strconv"

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

// SearchUser : search user
func SearchUser() {
	key, value := EnterSearchTermAndValue()

	switch key {
	case "_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			utils.CheckError(err)
		}
		if user, ok := userByID[id]; !ok {
			fmt.Println(aurora.Red("User with specific ID"), aurora.Red(id), aurora.Red("doesn't exist."))
		} else {
			utils.PrintPtrStructObject(user)
		}
	case "external_id":
		externalID := value
		if user, ok := userByExternalID[externalID]; !ok {
			fmt.Println(aurora.Red("User with specific ExternalID"), aurora.Red(externalID), aurora.Red("doesn't exist."))
		} else {
			utils.PrintPtrStructObject(user)
		}
	case "name":
		name := value
		if user, ok := userByName[name]; !ok {
			fmt.Println(aurora.Red("User with specific name"), aurora.Red(name), aurora.Red("doesn't exist."))
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
