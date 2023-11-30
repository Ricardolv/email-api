package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// DELETE /campaigns/delete/{id}
func (h *Handler) CampaignDelete(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	err := h.CampaignService.Delete(id)
	return nil, 200, err
}
