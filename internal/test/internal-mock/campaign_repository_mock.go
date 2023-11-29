package internalmock

import (
	"email-api/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type CampaignRepositoryMock struct {
	mock.Mock
}

func (r *CampaignRepositoryMock) Create(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *CampaignRepositoryMock) Update(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *CampaignRepositoryMock) Get() ([]campaign.Campaign, error) {
	//args := r.Called(campaign)
	return nil, nil
}
