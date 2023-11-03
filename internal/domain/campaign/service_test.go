package campaign

import (
	internalErrors "emailn/internal/InternalErrors"
	"emailn/internal/contract"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
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

func Test_Create_Campaign_Domain_Error(t *testing.T) {
	// arrange
	repositoryMock := new(repositoryMock)
	service.Repository = repositoryMock
	newCampaign.Name = ""

	// act
	_, err := service.Create(newCampaign)
	// assert
	assert.NotNil(t, err.Error())

}

func Test_Save_Campaign(t *testing.T) {
	// arrange
	repositoryMock := new(repositoryMock)
	service.Repository = repositoryMock
	newCampaign.Name = "Campaign 1"
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

func Test_Error_Save_Campaign(t *testing.T) {
	// arrange
	repositoryMock := new(repositoryMock)
	service.Repository = repositoryMock
	newCampaign.Name = ""
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error saving campaign"))
	// act
	_, err := service.Create(newCampaign)
	// assert
	assert.NotNil(t, err.Error())
	assert.True(t, errors.Is(internalErrors.ErrInternal, err))
}
