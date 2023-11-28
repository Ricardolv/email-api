package persistence

import "email-api/internal/domain/campaign"

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (c *CampaignRepository) Create(campaign *campaign.Campaign) error {
	c.campaigns = append(c.campaigns, *campaign)
	return nil
}

func (c *CampaignRepository) FindAll() []campaign.Campaign {
	return c.campaigns
}
