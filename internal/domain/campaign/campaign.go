package campaign

import (
	internalErrors "emailn/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"required,email"`
}

const (
	Pending  = "pending"
	Started  = "started"
	Finished = "finished"
)

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"required,min=5,max=24"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"required,min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))
	for idx, email := range emails {
		contacts[idx].Email = email
	}
	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
		Status:    Pending,
	}
	err := internalErrors.ValidateStruct(campaign)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}
