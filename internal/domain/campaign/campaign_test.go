package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign 1"
	content  = "Body 1"
	contacts = []string{"email1@e.com", "email2@e.com"}
)

func Test_NewCampaign_Create_Campaign(t *testing.T) {
	// act
	campaign, _ := NewCampaign(name, content, contacts)
	// assert
	assert.Equal(t, name, campaign.Name)
	assert.Equal(t, content, campaign.Content)
	assert.Equal(t, 2, len(campaign.Contacts))
}

func Test_NewCampaign_ID_notEmpty(t *testing.T) {
	// act
	campaign, _ := NewCampaign(name, content, contacts)
	// assert
	assert.NotEmpty(t, campaign.ID)
}

func Test_NewCampaign_CreatedOn_Must_Be_Now(t *testing.T) {
	// arrange
	now := time.Now().Add(-time.Minute)
	// act
	campaign, _ := NewCampaign(name, content, contacts)
	// assert
	assert.NotEmpty(t, campaign.CreatedOn)
	assert.True(t, campaign.CreatedOn.After(now))
}

func Test_NewCampaign_Must_Have_Name(t *testing.T) {
	// act
	_, err := NewCampaign("", content, contacts)
	// assert
	assert.EqualError(t, err, "name is required")
}

func Test_NewCampaign_Must_Have_Name_With_Minimum_5_Characters(t *testing.T) {
	// act
	_, err := NewCampaign("1234", content, contacts)
	// assert
	assert.EqualError(t, err, "name must be greater than 5")
}

func Test_NewCampaign_Must_Have_Name_With_Maximum_24_Characters(t *testing.T) {
	// arrange
	fake := faker.New()
	// act
	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)
	// assert
	assert.EqualError(t, err, "name must be less than 24")
}

func Test_NewCampaign_Must_Have_Content(t *testing.T) {
	// act
	_, err := NewCampaign(name, "", contacts)
	// assert
	assert.EqualError(t, err, "content is required")
}

func Test_NewCampaign_Must_Have_Content_With_Minimum_5_Characters(t *testing.T) {
	// act
	_, err := NewCampaign(name, "1234", contacts)
	// assert
	assert.EqualError(t, err, "content must be greater than 5")
}

func Test_NewCampaign_Must_Have_Content_With_Maximum_1024_Characters(t *testing.T) {
	// arrange
	fake := faker.New()
	// act
	_, err := NewCampaign(name, fake.Lorem().Text(1100), contacts)
	// assert
	assert.EqualError(t, err, "content must be less than 1024")
}

func Test_NewCampaign_Must_Have_Contacts(t *testing.T) {
	// act
	_, err := NewCampaign(name, content, []string{})
	// assert
	assert.EqualError(t, err, "contacts must be greater than 1")
}

func Test_NewCampaign_Must_Have_Contacts_With_Valid_Email(t *testing.T) {
	// act
	_, err := NewCampaign(name, content, []string{"invalid-email"})
	// assert
	assert.EqualError(t, err, "email is not a valid email")
}

func Test_NewCampaign_Must_Have_At_Least_One_Contact(t *testing.T) {
	// act
	_, err := NewCampaign(name, content, nil)
	// assert
	assert.EqualError(t, err, "contacts must be greater than 1")
}
