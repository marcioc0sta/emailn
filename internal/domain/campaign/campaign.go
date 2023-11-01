package campaign

import (
	"errors"
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

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if content == "" {
		return nil, errors.New("content is required")
	}
	if len(emails) == 0 {
		return nil, errors.New("at least one email is required")
	}
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
	}, nil
}
