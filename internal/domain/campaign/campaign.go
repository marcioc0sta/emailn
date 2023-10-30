package campaign

import (
	"time"

	"github.com/rs/xid"
)

type Contatct struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contatct
}

func NewCampaign(name string, content string, emails []string) *Campaign {
	contacts := make([]Contatct, len(emails))
	for idx, email := range emails {
		contacts[idx].Email = email
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
	}
}
