package campaign

import (
	"github.com/kaetaen/emailn/internal/domain/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	s.Repository.Save(campaign)

	return campaign.ID, nil
}
