package campaign

import (
	"emailn/internal/contract"
	internalErrors "emailn/internal/internal-errors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	Get() ([]contract.NewCampaign, error)
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internalErrors.ErrInternal
	}

	return campaign.ID, nil
}
