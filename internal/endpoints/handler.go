package endpoints

import "email-api/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
