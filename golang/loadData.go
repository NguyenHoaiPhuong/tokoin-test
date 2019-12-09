package main

import (
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/jsonfunc"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/model"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/utils"
)

var tickets model.Tickets
var organizations model.Organizations

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
