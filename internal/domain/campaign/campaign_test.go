package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)

	name := "Campaign X"
	content := "This is a test campaign."
	contacts := []string{"julia_contact@google.com", "paul_contact"}

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}
