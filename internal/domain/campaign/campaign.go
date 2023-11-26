package campaign

import "time"

type Contact struct {
	Email string `json:"email"`
}

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contact
}
