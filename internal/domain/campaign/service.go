package campaign

import "email-api/internal/contract"

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) error {

	return nil
}
