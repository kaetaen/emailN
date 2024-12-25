package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "This is a test campaign."
	contacts = []string{"julia_contact@google.com", "paul_contact@mail.com"}
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IdIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)
	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreateOn, now)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaign("", content, contacts)

	assert.Equal("name is required", error.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaign(name, "", contacts)

	assert.Equal("content is required", error.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaign(name, content, []string{})

	assert.Equal("contacts is required", error.Error())
}
