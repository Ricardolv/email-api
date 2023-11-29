package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, req *http.Request) (interface{}, int, error) {

	campaigns, err := h.CampaignService.Repository.FindAll()
	return campaigns, 200, err
}
