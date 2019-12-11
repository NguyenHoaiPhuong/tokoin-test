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
	tickets                 model.Tickets
	ticketByID              model.TicketByID
	ticketByExternalID      model.TicketByExternalID
	ticketBySubject         model.TicketBySubject
	ticketsByOrganizationID model.TicketsByOrganizationID
	ticketsBySubmitterID    model.TicketsBySubmitterID
	ticketsByAssigneeID     model.TicketsByAssigneeID
)

func initTicket() {
	tickets = make(model.Tickets, 0)
	ticketByID = make(model.TicketByID, 0)
	ticketByExternalID = make(model.TicketByExternalID, 0)
	ticketBySubject = make(model.TicketBySubject, 0)
	ticketsByOrganizationID = make(model.TicketsByOrganizationID, 0)
	ticketsBySubmitterID = make(model.TicketsBySubmitterID, 0)
	ticketsByAssigneeID = make(model.TicketsByAssigneeID, 0)
}

// LoadTicketData : load ticket
func LoadTicketData() {
	initTicket()

	err := jsonfunc.ReadFromFile(cfg.URLs.TicketURL, &tickets)
	utils.CheckError(err)

	for _, ticket := range tickets {
		appendToTicketByID(ticket)
		appendToTicketByExternalID(ticket)
		appendToTicketBySubject(ticket)
		appendToTicketsByOrganizationID(ticket)
		appendToTicketsBySubmitterID(ticket)
		appendToTicketsByAssigneeID(ticket)
	}
}

func appendToTicketByID(ticket *model.Ticket) {
	ticketByID[ticket.ID] = ticket
}

func appendToTicketByExternalID(ticket *model.Ticket) {
	ticketByExternalID[ticket.ExternalID] = ticket
}

func appendToTicketBySubject(ticket *model.Ticket) {
	ticketBySubject[ticket.Subject] = ticket
}

func appendToTicketsByOrganizationID(ticket *model.Ticket) {
	if _, ok := ticketsByOrganizationID[ticket.OrganizationID]; !ok {
		ticketsByOrganizationID[ticket.OrganizationID] = make(model.TicketBySubject, 0)
	}
	sortedTickets := ticketsByOrganizationID[ticket.OrganizationID]
	sortedTickets[ticket.Subject] = ticket
	ticketsByOrganizationID[ticket.OrganizationID] = sortedTickets
}

func appendToTicketsBySubmitterID(ticket *model.Ticket) {
	if _, ok := ticketsBySubmitterID[ticket.SubmitterID]; !ok {
		ticketsBySubmitterID[ticket.SubmitterID] = make(model.TicketBySubject, 0)
	}
	sortedTickets := ticketsBySubmitterID[ticket.SubmitterID]
	sortedTickets[ticket.Subject] = ticket
	ticketsBySubmitterID[ticket.SubmitterID] = sortedTickets
}

func appendToTicketsByAssigneeID(ticket *model.Ticket) {
	if _, ok := ticketsByAssigneeID[ticket.AssigneeID]; !ok {
		ticketsByAssigneeID[ticket.AssigneeID] = make(model.TicketBySubject, 0)
	}
	sortedTickets := ticketsByAssigneeID[ticket.AssigneeID]
	sortedTickets[ticket.Subject] = ticket
	ticketsByAssigneeID[ticket.AssigneeID] = sortedTickets
}

// SearchTicket : search ticket
func SearchTicket() {
	key, value := utils.EnterSearchTermAndValue()

	switch key {
	case "_id":
		ticket, isFound := ticketByID[value]
		utils.PrintObject(ticket, isFound, key, value)
		if isFound {
			SearchTicketUsers(ticket)
			SearchTicketOrganization(ticket)
		}
	case "external_id":
		ticket, isFound := ticketByExternalID[value]
		utils.PrintObject(ticket, isFound, key, value)
		if isFound {
			SearchTicketUsers(ticket)
			SearchTicketOrganization(ticket)
		}
	case "subject":
		ticket, isFound := ticketBySubject[value]
		utils.PrintObject(ticket, isFound, key, value)
		if isFound {
			SearchTicketUsers(ticket)
			SearchTicketOrganization(ticket)
		}
	case "organization_id":
		orgID, err := strconv.Atoi(value)
		if err != nil {
			utils.CheckError(err)
			return
		}
		sortedTickets, isFound := ticketsByOrganizationID[orgID]
		idx := 1
		for _, ticket := range sortedTickets {
			fmt.Println(aurora.BrightYellow("Ticket " + strconv.Itoa(idx)))
			utils.PrintObject(ticket, isFound, key, value)
			if isFound {
				SearchTicketUsers(ticket)
				SearchTicketOrganization(ticket)
			}
			idx++
		}
	case "submitter_id":
		submitterID, err := strconv.Atoi(value)
		if err != nil {
			utils.CheckError(err)
			return
		}
		sortedTickets, isFound := ticketsBySubmitterID[submitterID]
		idx := 1
		for _, ticket := range sortedTickets {
			fmt.Println(aurora.BrightYellow("Ticket " + strconv.Itoa(idx)))
			utils.PrintObject(ticket, isFound, key, value)
			if isFound {
				SearchTicketUsers(ticket)
				SearchTicketOrganization(ticket)
			}
			idx++
		}
	case "assignee_id":
		assigneeID, err := strconv.Atoi(value)
		if err != nil {
			utils.CheckError(err)
			return
		}
		tickets, isFound := ticketsByAssigneeID[assigneeID]
		idx := 1
		for _, ticket := range tickets {
			fmt.Println(aurora.BrightYellow("Ticket " + strconv.Itoa(idx)))
			utils.PrintObject(ticket, isFound, key, value)
			if isFound {
				SearchTicketUsers(ticket)
				SearchTicketOrganization(ticket)
			}
			idx++
		}
	default:
		fmt.Println(aurora.Red("Searching term " + key + " hasn't been supported yet"))
	}

	utils.PrintSeparation()
}

// SearchTicketUsers : search all users related to this ticket
func SearchTicketUsers(ticket *model.Ticket) {
	assigneeID := ticket.AssigneeID
	user, isFound := userByID[assigneeID]
	if !isFound {
		utils.NoResult(user, "_id", strconv.Itoa(assigneeID))
	} else {
		fmt.Println(aurora.BrightCyan("Assignee name"), ":", aurora.BrightGreen(user.Name))
	}

	submitterID := ticket.SubmitterID
	user, isFound = userByID[submitterID]
	if !isFound {
		utils.NoResult(user, "_id", strconv.Itoa(submitterID))
	} else {
		fmt.Println(aurora.BrightCyan("Submitter name"), ":", aurora.BrightGreen(user.Name))
	}
}

// SearchTicketOrganization : search organization related to this ticket
func SearchTicketOrganization(ticket *model.Ticket) {
	orgID := ticket.OrganizationID
	org, isFound := organizationByID[orgID]
	if !isFound {
		utils.NoResult(org, "_id", strconv.Itoa(orgID))
	} else {
		fmt.Println(aurora.BrightCyan("Organization name"), ":", aurora.BrightGreen(org.Name))
	}
}
