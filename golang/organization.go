package main

import (
	"fmt"
	"strconv"

	"github.com/NguyenHoaiPhuong/tokoin-test/golang/jsonfunc"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/model"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/utils"
	"github.com/logrusorgru/aurora"
)

var (
	organizations            model.Organizations
	organizationByID         model.OrganizationByID
	organizationByExternalID model.OrganizationByExternalID
	organizationByName       model.OrganizationByName
)

func initOrganization() {
	organizations = make(model.Organizations, 0)
	organizationByID = make(model.OrganizationByID, 0)
	organizationByExternalID = make(model.OrganizationByExternalID, 0)
	organizationByName = make(model.OrganizationByName, 0)
}

// LoadOrganizationData : load organization
func LoadOrganizationData() {
	initOrganization()

	err := jsonfunc.ReadFromFile(cfg.URLs.OrganizationURL, &organizations)
	utils.CheckError(err)
}

// ConsolidateOrganizationData : store organizations in maps
func ConsolidateOrganizationData() {
	for _, organization := range organizations {
		appendToOrganizationByID(organization)
		appendToOrganizationByExternalID(organization)
		appendToOrganizationByName(organization)
	}
}

func appendToOrganizationByID(organization *model.Organization) {
	organizationByID[organization.ID] = organization
}

func appendToOrganizationByExternalID(organization *model.Organization) {
	organizationByExternalID[organization.ExternalID] = organization
}

func appendToOrganizationByName(organization *model.Organization) {
	organizationByName[organization.Name] = organization
}

// SearchOrganization : search organization
func SearchOrganization() {
	key, value := utils.EnterSearchTermAndValue()

	var org *model.Organization
	var isFound bool

	switch key {
	case "_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			utils.CheckError(err)
			return
		}
		org, isFound = organizationByID[id]
		utils.PrintObject(org, isFound, key, value)
	case "external_id":
		org, isFound = organizationByExternalID[value]
		utils.PrintObject(org, isFound, key, value)
	case "name":
		org, isFound = organizationByName[value]
		utils.PrintObject(org, isFound, key, value)
	default:
		fmt.Println(aurora.Red("Searching term " + key + " hasn't been supported yet"))
	}

	if isFound {
		SearchOrganizationTickets(org)
		SearchOrganizationUsers(org)
	}

	utils.PrintSeparation()
}

// SearchOrganizationTickets : search all tickets related to this organization
func SearchOrganizationTickets(org *model.Organization) {
	id := org.ID
	sortedTickets := ticketsByOrganizationID[id]
	idx := 1
	for ticketSubject := range sortedTickets {
		fmt.Println(aurora.BrightCyan("Ticket "+strconv.Itoa(idx)), ":", aurora.BrightGreen(ticketSubject))
		idx++
	}
}

// SearchOrganizationUsers : search all users related to this organization
func SearchOrganizationUsers(org *model.Organization) {
	id := org.ID
	sortedUsers := usersByOrganizationID[id]
	idx := 1
	for userName := range sortedUsers {
		fmt.Println(aurora.BrightCyan("User "+strconv.Itoa(idx)), ":", aurora.BrightGreen(userName))
		idx++
	}
}
