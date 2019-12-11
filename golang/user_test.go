package main

import (
	"testing"

	"github.com/NguyenHoaiPhuong/tokoin-test/golang/model"
)

func mockUsers() {
	user := &model.User{
		ID:             19,
		URL:            "http://initech.tokoin.io.com/api/v2/users/19.json",
		ExternalID:     "68e35e26-7b1f-46ec-a9e5-3edcbcf2aeb9",
		Name:           "Francis Rodrigüez",
		Alias:          "Mr Lea",
		CreatedAt:      "2016-05-05T01:56:24 -10:00",
		Active:         false,
		Verified:       false,
		Shared:         false,
		Locale:         "zh-CN",
		Timezone:       "Brazil",
		LastLoginAt:    "2012-05-25T01:55:34 -10:00",
		Phone:          "8275-873-442",
		Signature:      "Don't Worry Be Happy!",
		OrganizationID: 124,
		Tags: []string{
			"Vicksburg",
			"Kilbourne",
			"Gorham",
			"Gloucester",
		},
		Suspended: false,
		Role:      "agent",
	}
	users = append(users, user)

	user = &model.User{
		ID:             20,
		URL:            "http://initech.tokoin.io.com/api/v2/users/20.json",
		ExternalID:     "1c8b9696-b5c5-4f5a-b077-5d85dd8b22f4",
		Name:           "Lee Davidson",
		Alias:          "Miss Hayes",
		CreatedAt:      "2016-04-07T11:08:24 -10:00",
		Active:         false,
		Verified:       false,
		Shared:         true,
		Locale:         "en-AU",
		Timezone:       "Norway",
		LastLoginAt:    "2014-05-04T09:18:56 -10:00",
		Phone:          "9544-832-980",
		Signature:      "Don't Worry Be Happy!",
		OrganizationID: 104,
		Tags: []string{
			"Riegelwood",
			"Snyderville",
			"Roland",
			"National",
		},
		Suspended: true,
		Role:      "end-user",
	}
}

func TestConsolidateUserData(t *testing.T) {
	initUser()
	mockUsers()
	ConsolidateUserData()

	// Testing
	name := "Francis Rodrigüez"
	if _, ok := userByName[name]; !ok {
		t.Errorf("User %s must exist", name)
	}

	name = "akagi"
	if _, ok := userByName[name]; ok {
		t.Errorf("User %s must NOT exist", name)
	}
}
