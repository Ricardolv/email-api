package campaign

import "time"

type Contact struct {
	Email string `json:"email"`
}

type Campaign struct {
	ID        string
	Name      string    `json:"name"`
	CreatedOn time.Time `json:"created_on"`
	Content   string    `json:"content"`
	Contacts  []Contact `json:"contacts"`
}

func NewCampaign(name string, content string, emails []string) *Campaign {

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		ID:        "1",
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
	}
}
