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

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IdIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)
	campaign := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreateOn, now)
}
