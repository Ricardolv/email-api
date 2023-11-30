package persistence

import "email-api/internal/domain/campaign"

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (c *CampaignRepository) Create(campaign *campaign.Campaign) error {
	c.campaigns = append(c.campaigns, *campaign)
	return nil
}

func (c *CampaignRepository) Update(campaign *campaign.Campaign) error {
	// tx := c.Db.Save(campaign)
	// return tx.Error

	return nil
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	// var campaigns []campaign.Campaign
	// tx := c.Db.Find(&campaigns)
	// return campaigns, tx.Error
	return nil, nil
}

func (c *CampaignRepository) GetBy(id string) (*campaign.Campaign, error) {
	// var campaign campaign.Campaign
	// tx := c.Db.Preload("Contacts").First(&campaign, "id = ?", id)
	// return &campaign, tx.Error
	return nil, nil
}

func (c *CampaignRepository) Delete(campaign *campaign.Campaign) error {
	// tx := c.Db.Select("Contacts").Delete(campaign)
	// return tx.Error
	return nil
}
