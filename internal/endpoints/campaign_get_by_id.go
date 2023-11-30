package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// GET /campaigns/{id}
func (h *Handler) CampaignGetById(w http.ResponseWriter, req *http.Request) (interface{}, int, error) {
	id := chi.URLParam(req, "id")
	campaign, err := h.CampaignService.GetBy(id)
	if err == nil && campaign == nil {
		return nil, http.StatusNotFound, err
	}
	return campaign, 200, err
}
