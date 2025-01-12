package campaign

import "github.com/kaetaen/emailn/internal/domain/contract"

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) error {
	return nil
}

func teste() {

}
