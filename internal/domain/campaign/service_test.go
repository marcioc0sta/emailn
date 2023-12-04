package campaign

import (
	"emailn/internal/contract"
	internalErrors "emailn/internal/internal-errors"
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

func (r *repositoryMock) Get() ([]Campaign, error) {
	return nil, nil
}

func (r *repositoryMock) GetBy(id string) (*Campaign, error) {
	args := r.Called(id)
	return args.Get(0).(*Campaign), args.Error(1)
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Campaign 1",
		Content: "Body 1",
		Emails:  []string{"email1@email.com", "email2@email.com"},
	}
	service = ServiceImp{}
)

func Test_Create_Campaign(t *testing.T) {
	// arrange
	repositoryMock := new(repositoryMock)
	service.Repository = repositoryMock
	repositoryMock.On("Save", mock.Anything).Return(nil)
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
	assert.False(t, errors.Is(internalErrors.ErrInternal, err))
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
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error"))
	// act
	_, err := service.Create(newCampaign)
	// assert
	assert.NotNil(t, err.Error())
	assert.True(t, errors.Is(internalErrors.ErrInternal, err))
	repositoryMock.AssertExpectations(t)
}

func Test_GetById(t *testing.T) {
	// arrange
	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	repositoryMock := new(repositoryMock)
	service.Repository = repositoryMock
	repositoryMock.On("GetBy", mock.Anything).Return(campaign, nil)
	// act
	campaignResponse, _ := service.GetBy(campaign.ID)
	// assert
	assert.Equal(t, campaign.ID, campaignResponse.ID)
}
