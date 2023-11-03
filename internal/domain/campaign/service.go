package campaign

import (
	internalErrors "emailn/internal/InternalErrors"
	"emailn/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", internalErrors.ErrInternal
	}
	s.Repository.Save(campaign)

	return campaign.ID, nil
}
