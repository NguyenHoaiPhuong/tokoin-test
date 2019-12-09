package model

// Ticket struct
type Ticket struct {
	ID             string   `json:"_id"`
	URL            string   `json:"url"`
	ExternalID     string   `json:"external_id"`
	CreatedAt      string   `json:"created_at"`
	Type           string   `json:"type"`
	Subject        string   `json:"subject"`
	Description    string   `json:"description"`
	Priority       string   `json:"priority"`
	Status         string   `json:"status"`
	SubmitterID    int      `json:"submitter_id"`
	AssigneeID     int      `json:"assignee_id"`
	OrganizationID int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	HasIncidents   bool     `json:"has_incidents"`
	DueAt          string   `json:"due_at"`
	Via            string   `json:"via"`
}

// Tickets : slice of Ticket
type Tickets []*Ticket

// TicketByID : sorted ticket by its ID
type TicketByID map[int]*Ticket

// TicketByExternalID : sorted ticket by its ExternalID
type TicketByExternalID map[string]*Ticket

// TicketBySubject : sorted ticket by its subject
type TicketBySubject map[string]*Ticket

// TicketsByOrganizationID : map of tickets and OrganizationID. Key is OrganizationID.
type TicketsByOrganizationID map[int]TicketBySubject

// TicketsBySubmitterID : sorted ticket by its ExternalID
type TicketsBySubmitterID map[int]TicketBySubject
