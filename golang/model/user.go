package model

// User struct
type User struct {
	ID             int      `json:"_id"`
	URL            string   `json:"url"`
	ExternalID     string   `json:"external_id"`
	Name           string   `json:"name"`
	Alias          string   `json:"alias"`
	CreatedAt      string   `json:"created_at"`
	Active         bool     `json:"active"`
	Verified       bool     `json:"verified"`
	Shared         bool     `json:"shared"`
	Locale         string   `json:"locale"`
	Timezone       string   `json:"timezone"`
	LastLoginAt    string   `json:"last_login_at"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	Signature      string   `json:"signature"`
	OrganizationID int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	Suspended      bool     `json:"suspended"`
	Role           string   `json:"role"`
}

// Users : slice of users
type Users []*User

// UserByID : sorted user by its ID
type UserByID map[int]*User

// UserByURL : sorted user by its URL
type UserByURL map[string]*User

// UserByExternalID : sorted user by its external ID
type UserByExternalID map[string]*User

// UserByName : sorted user by its name
type UserByName map[string]*User

// UserByAlias : sorted user by its alias
type UserByAlias map[string]*User

// UserByCreatedAt : sorted user by its alias
type UserByCreatedAt map[string]*User

// UsersByActive : map of users and their property Active. Key is the property Active.
type UsersByActive map[bool]UserByName

// UsersByOrganizationID : map of users and OrganizationID. Key is OrganizationID.
type UsersByOrganizationID map[int]UserByName
