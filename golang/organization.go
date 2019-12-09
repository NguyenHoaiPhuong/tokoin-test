package main

import (
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/jsonfunc"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/model"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/utils"
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

}

func appendToOrganizationByExternalID(organization *model.Organization) {

}

func appendToOrganizationByName(organization *model.Organization) {

}

// SearchOrganization : search organization
func SearchOrganization() {

}
