package main

import (
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/jsonfunc"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/model"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/utils"
)

var organizations model.Organizations

func initOrganization() {
	organizations = make(model.Organizations, 0)
}

// LoadOrganizationData : load organization
func LoadOrganizationData() {
	err := jsonfunc.ReadFromFile(cfg.URLs.OrganizationURL, &organizations)
	utils.CheckError(err)
	// fmt.Println(organizations[0].Details)
}

// SearchOrganization : search organization
func SearchOrganization() {

}
