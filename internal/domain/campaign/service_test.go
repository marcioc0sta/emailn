package campaign

import (
	"emailn/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	println("Mocked Save")
	println(campaign.Name)
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Campaign 1",
		Content: "Body",
		Emails:  []string{"email1@email", "email2@email"},
	}
	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	// arrange
	repositoryMock := new(repositoryMock)
	service.Repository = repositoryMock

	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	// act
	id, _ := service.Create(newCampaign)

	// assert
	assert.NotEmpty(t, id)
	repositoryMock.AssertExpectations(t)
}

func Test_Save_Campaign(t *testing.T) {
	// arrange
	repositoryMock := new(repositoryMock)
	service.Repository = repositoryMock

	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	// act
	service.Create(newCampaign)

	// assert
	repositoryMock.AssertExpectations(t)

}
