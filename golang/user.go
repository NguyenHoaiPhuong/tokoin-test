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
	users                 model.Users
	userByID              model.UserByID              // Sorted user by ID
	userByExternalID      model.UserByExternalID      // Sorted user by ExternalID
	userByName            model.UserByName            // Sorted user by its name
	usersByOrganizationID model.UsersByOrganizationID // Sorted users by OrganizationID
)

func initUser() {
	users = make(model.Users, 0)
	userByID = make(model.UserByID, 0)
	userByExternalID = make(model.UserByExternalID, 0)
	userByName = make(model.UserByName, 0)
	usersByOrganizationID = make(model.UsersByOrganizationID, 0)
}

// LoadUserData : load user
func LoadUserData() {
	initUser()

	err := jsonfunc.ReadFromFile(cfg.URLs.UserURL, &users)
	utils.CheckError(err)

	for _, user := range users {
		appendToUserByID(user)
		appendToUserByExternalID(user)
		appendToUserByName(user)
	}
}

func appendToUserByID(user *model.User) {
	userByID[user.ID] = user
}

func appendToUserByExternalID(user *model.User) {
	userByExternalID[user.ExternalID] = user
}

func appendToUserByName(user *model.User) {
	// IMPORTANT: user name is unique???
	userByName[user.Name] = user

	if _, ok := usersByOrganizationID[user.OrganizationID]; !ok {
		usersByOrganizationID[user.OrganizationID] = make(model.UserByName, 0)
	}
	sortedUsers := usersByOrganizationID[user.OrganizationID]
	sortedUsers[user.Name] = user
	usersByOrganizationID[user.OrganizationID] = sortedUsers
}

// SearchUser : search user
func SearchUser() {
	key, value := EnterSearchTermAndValue()

	switch key {
	case "_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			utils.CheckError(err)
			return
		}
		user, ok := userByID[id]
		utils.PrintObject(user, ok, key, value)
	case "external_id":
		user, ok := userByExternalID[value]
		utils.PrintObject(user, ok, key, value)
	case "name":
		user, ok := userByName[value]
		utils.PrintObject(user, ok, key, value)
	default:
		fmt.Println(aurora.Red("Searching term " + key + " hasn't been supported yet"))
	}
}
