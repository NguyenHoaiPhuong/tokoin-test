package model

// Organization struct
type Organization struct {
	ID            int      `json:"_id"`
	URL           string   `json:"url"`
	ExternalID    string   `json:"external_id"`
	Name          string   `json:"name"`
	DomainNames   []string `json:"domain_names"`
	CreatedAt     string   `json:"created_at"`
	Details       string   `json:"details"`
	SharedTickets bool     `json:"shared_tickets"`
	Tags          []string `json:"tags"`
}

// Organizations : slice of Organization
type Organizations []*Organization

// OrganizationByID : sorted organization by its ID
type OrganizationByID map[int]*Organization

// OrganizationByExternalID : sorted organization by its ExternalID
type OrganizationByExternalID map[string]*Organization

// OrganizationByName : sorted organization by its Name
type OrganizationByName map[string]*Organization
