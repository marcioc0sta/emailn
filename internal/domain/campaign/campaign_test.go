package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign 1"
	content  = "Body"
	contacts = []string{"email1@e.com", "email2@e.com"}
)

func Test_NewCampaign_Create_Campaign(t *testing.T) {
	// arrange
	assert := assert.New(t)
	// act
	campaign, _ := NewCampaign(name, content, contacts)
	// assert
	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(2, len(campaign.Contacts))
}

func Test_NewCampaign_ID_notEmpty(t *testing.T) {
	// arrange
	assert := assert.New(t)
	// act
	campaign, _ := NewCampaign(name, content, contacts)
	// assert
	assert.NotEmpty(campaign.ID)
}

func Test_NewCampaign_CreatedOn_Must_Be_Now(t *testing.T) {
	// arrange
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)
	// act
	campaign, _ := NewCampaign(name, content, contacts)
	// assert
	assert.NotEmpty(campaign.CreatedOn)
	assert.True(campaign.CreatedOn.After(now))
}

func Test_NewCampaign_Must_Have_Name(t *testing.T) {
	// arrange
	assert := assert.New(t)
	// act
	_, err := NewCampaign("", content, contacts)
	// assert
	assert.EqualError(err, "name is required")
}

func Test_NewCampaign_Must_Have_Content(t *testing.T) {
	// arrange
	assert := assert.New(t)
	// act
	_, err := NewCampaign(name, "", contacts)
	// assert
	assert.EqualError(err, "content is required")
}

func Test_NewCampaign_Must_Have_Contacts(t *testing.T) {
	// arrange
	assert := assert.New(t)
	// act
	_, err := NewCampaign(name, content, []string{})
	// assert
	assert.EqualError(err, "at least one email is required")
}
