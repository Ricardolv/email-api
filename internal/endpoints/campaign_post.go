package endpoints

import (
	"email-api/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, req *http.Request) (interface{}, int, error) {
	var request contract.NewCampaign
	render.DecodeJSON(req.Body, &request)
	email := req.Context().Value("email").(string)
	request.CreatedBy = email
	id, err := h.CampaignService.Create(request)
	return map[string]string{"id": id}, 201, err
}
