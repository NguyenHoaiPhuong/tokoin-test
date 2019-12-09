package main

import (
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/jsonfunc"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/model"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/utils"
)

var (
	users             model.Users
	usersByID         map[int]*model.User    // Sorted user by ID
	usersByExternalID map[string]*model.User // Sorted user by ExternalID
)

var tickets model.Tickets
var organizations model.Organizations

// LoadData : read json files and store data into heap
func LoadData() {
	LoadUserData()
	LoadTicketData()
	LoadOrganizationData()
}

// LoadUserData : load user
func LoadUserData() {
	err := jsonfunc.ReadFromFile(cfg.URLs.UserURL, &users)
	utils.CheckError(err)
	// fmt.Println(users[0].Email)

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
