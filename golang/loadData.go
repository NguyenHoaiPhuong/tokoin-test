package main

import (
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/jsonfunc"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/model"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/utils"
)

var (
	users                 model.Users
	userByID              model.UserByID         // Sorted user by ID
	userByExternalID      model.UserByExternalID // Sorted user by ExternalID
	userByName            model.UserByName
	usersByOrganizationID model.UsersByOrganizationID // Sorted users by OrganizationID
)

var tickets model.Tickets
var organizations model.Organizations

func initUser() {
	users = make(model.Users, 0)
	userByID = make(model.UserByID, 0)
	userByExternalID = make(model.UserByExternalID, 0)
	userByName = make(model.UserByName, 0)
	usersByOrganizationID = make(model.UsersByOrganizationID, 0)
}

func initTicket() {
	tickets = make(model.Tickets, 0)
}

func initOrganization() {
	organizations = make(model.Organizations, 0)
}

// LoadData : read json files and store data into heap
func LoadData() {
	LoadUserData()
	LoadTicketData()
	LoadOrganizationData()
}

// LoadUserData : load user
func LoadUserData() {
	initUser()

	err := jsonfunc.ReadFromFile(cfg.URLs.UserURL, &users)
	utils.CheckError(err)
	for _, user := range users {
		id := user.ID
		externalID := user.ExternalID
		name := user.Name
		orgID := user.OrganizationID
		userByID[id] = user
		userByExternalID[externalID] = user
		// IMPORTANT: user name is unique???
		userByName[name] = user

		if _, ok := usersByOrganizationID[orgID]; !ok {
			usersByOrganizationID[orgID] = make(model.UserByName, 0)
		}
		sortedUsers := usersByOrganizationID[orgID]
		sortedUsers[name] = user
		usersByOrganizationID[orgID] = sortedUsers
	}
}

// LoadTicketData : load ticket
func LoadTicketData() {
	err := jsonfunc.ReadFromFile(cfg.URLs.TicketURL, &tickets)
	utils.CheckError(err)
	// fmt.Println(tickets[0].CreatedAt)
}

// LoadOrganizationData : load organization
func LoadOrganizationData() {
	err := jsonfunc.ReadFromFile(cfg.URLs.OrganizationURL, &organizations)
	utils.CheckError(err)
	// fmt.Println(organizations[0].Details)
}
