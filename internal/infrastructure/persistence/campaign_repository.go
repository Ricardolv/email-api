package persistence

import "email-api/internal/domain/campaign"

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (r *CampaignRepository) Create(campaign *campaign.Campaign) error {
	r.campaigns = append(r.campaigns, *campaign)
	return nil
}
