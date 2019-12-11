package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/NguyenHoaiPhuong/tokoin-test/golang/jsonfunc"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/model"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/utils"
	"github.com/logrusorgru/aurora"
)

var (
	users                 model.Users
	userByID              model.UserByID              // Sorted user by ID
	userByURL             model.UserByURL             // Sorted user by URL
	userByExternalID      model.UserByExternalID      // Sorted user by ExternalID
	userByName            model.UserByName            // Sorted user by its name
	userByAlias           model.UserByAlias           // Sorted user by its alias
	userByCreatedAt       model.UserByCreatedAt       // Sorted user by its CreatedAt
	usersByActive         model.UsersByActive         // Sorted users by Active
	usersByOrganizationID model.UsersByOrganizationID // Sorted users by OrganizationID
)

func initUser() {
	users = make(model.Users, 0)
	userByID = make(model.UserByID, 0)
	userByURL = make(model.UserByURL, 0)
	userByExternalID = make(model.UserByExternalID, 0)
	userByName = make(model.UserByName, 0)
	userByAlias = make(model.UserByAlias, 0)
	userByCreatedAt = make(model.UserByCreatedAt, 0)
	usersByActive = make(model.UsersByActive, 0)
	usersByOrganizationID = make(model.UsersByOrganizationID, 0)
}

// LoadUserData : load user
func LoadUserData() {
	initUser()

	err := jsonfunc.ReadFromFile(cfg.URLs.UserURL, &users)
	utils.CheckError(err)

	for _, user := range users {
		appendToUserByID(user)
		appendToUserByURL(user)
		appendToUserByExternalID(user)
		appendToUserByName(user)
		appendToUserByAlias(user)
		appendToUserByCreatedAt(user)
		appendToUsersByActive(user)
		appendToUsersByOrganizationID(user)
	}
}

func appendToUserByID(user *model.User) {
	userByID[user.ID] = user
}

func appendToUserByURL(user *model.User) {
	userByURL[user.URL] = user
}

func appendToUserByExternalID(user *model.User) {
	userByExternalID[user.ExternalID] = user
}

func appendToUserByName(user *model.User) {
	// IMPORTANT: user name is unique???
	userByName[user.Name] = user
}

func appendToUserByAlias(user *model.User) {
	userByAlias[user.Alias] = user
}

func appendToUserByCreatedAt(user *model.User) {
	userByCreatedAt[user.CreatedAt] = user
}

func appendToUsersByActive(user *model.User) {
	if _, ok := usersByActive[user.Active]; !ok {
		usersByActive[user.Active] = make(model.UserByName, 0)
	}
	sortedUsers := usersByActive[user.Active]
	sortedUsers[user.Name] = user
	usersByActive[user.Active] = sortedUsers
}

func appendToUsersByOrganizationID(user *model.User) {
	if _, ok := usersByOrganizationID[user.OrganizationID]; !ok {
		usersByOrganizationID[user.OrganizationID] = make(model.UserByName, 0)
	}
	sortedUsers := usersByOrganizationID[user.OrganizationID]
	sortedUsers[user.Name] = user
	usersByOrganizationID[user.OrganizationID] = sortedUsers
}

// SearchUser : search user
func SearchUser() {
	key, value := utils.EnterSearchTermAndValue()

	var user *model.User
	var isFound bool

	switch key {
	case "_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			utils.CheckError(err)
			return
		}
		user, isFound = userByID[id]
		PrintSingleUser(user, isFound, key, value)
	case "url":
		user, isFound = userByURL[value]
		PrintSingleUser(user, isFound, key, value)
	case "external_id":
		user, isFound = userByExternalID[value]
		PrintSingleUser(user, isFound, key, value)
	case "name":
		user, isFound = userByName[value]
		PrintSingleUser(user, isFound, key, value)
	case "alias":
		user, isFound = userByAlias[value]
		PrintSingleUser(user, isFound, key, value)
	case "created_at":
		user, isFound = userByCreatedAt[value]
		PrintSingleUser(user, isFound, key, value)
	case "active":
		active := false
		if value == "true" {
			active = true
		} else if value != "false" {
			err := errors.New("Active value must be either true or false")
			utils.CheckError(err)
			return
		}
		sortedUsers, isFound := usersByActive[active]
		PrintMultipleUsers(sortedUsers, isFound, key, value)
	case "organization_id":
		orgID, err := strconv.Atoi(value)
		if err != nil {
			utils.CheckError(err)
			return
		}
		sortedUsers, isFound := usersByOrganizationID[orgID]
		PrintMultipleUsers(sortedUsers, isFound, key, value)
	default:
		fmt.Println(aurora.Red("Searching term " + key + " hasn't been supported yet"))
	}

	utils.PrintSeparation()
}

// SearchUserTickets : search all tickets related to this user
func SearchUserTickets(user *model.User) {
	id := user.ID

	sortedTickets := ticketsByAssigneeID[id]
	idx := 1
	for ticketSubject := range sortedTickets {
		fmt.Println(aurora.BrightCyan("Assignee Ticket "+strconv.Itoa(idx)), ":", aurora.BrightGreen(ticketSubject))
		idx++
	}

	sortedTickets = ticketsBySubmitterID[id]
	idx = 1
	for ticketSubject := range sortedTickets {
		fmt.Println(aurora.BrightCyan("Submitter Ticket "+strconv.Itoa(idx)), ":", aurora.BrightGreen(ticketSubject))
		idx++
	}
}

// SearchUserOrganization : search organization related to this user
func SearchUserOrganization(user *model.User) {
	orgID := user.OrganizationID
	org, isFound := organizationByID[orgID]
	if !isFound {
		utils.NoResult(org, "_id", strconv.Itoa(orgID))
	} else {
		fmt.Println(aurora.BrightCyan("Organization name"), ":", aurora.BrightGreen(org.Name))
	}
}

// PrintSingleUser : print 1 user
func PrintSingleUser(user *model.User, isFound bool, key, value string) {
	utils.PrintObject(user, isFound, key, value)
	if isFound {
		SearchUserTickets(user)
		SearchUserOrganization(user)
	}
}

// PrintMultipleUsers : print a slice of users
func PrintMultipleUsers(sortedUsers model.UserByName, isFound bool, key, value string) {
	idx := 1
	for _, user := range sortedUsers {
		fmt.Println(aurora.BrightYellow("User " + strconv.Itoa(idx)))
		utils.PrintObject(user, isFound, key, value)
		if isFound {
			SearchUserTickets(user)
			SearchUserOrganization(user)
		}
		idx++
	}
}
