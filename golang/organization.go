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

	switch key {
	case "_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			utils.CheckError(err)
			return
		}
		org, ok := organizationByID[id]
		utils.PrintObject(org, ok, key, value)
	case "external_id":
		org, ok := organizationByExternalID[value]
		utils.PrintObject(org, ok, key, value)
	case "name":
		org, ok := organizationByName[value]
		utils.PrintObject(org, ok, key, value)
	default:
		fmt.Println(aurora.Red("Searching term " + key + " hasn't been supported yet"))
		fmt.Println()
	}
}
