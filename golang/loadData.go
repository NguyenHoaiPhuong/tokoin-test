package main

import (
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/jsonfunc"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/model"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/utils"
)

var (
	users             model.Users
	usersByID         model.UserByID         // Sorted user by ID
	usersByExternalID model.UserByExternalID // Sorted user by ExternalID
)

var tickets model.Tickets
var organizations model.Organizations

func initUser() {
	users = make(model.Users, 0)
	usersByID = make(model.UserByID, 0)
	usersByExternalID = make(model.UserByExternalID, 0)
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
	// fmt.Println(users[0].Email)
	for _, user := range users {
		id := user.ID
		externalID := user.ExternalID
		usersByID[id] = user
		usersByExternalID[externalID] = user
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
