package endpoints

import (
	"email-api/internal/contract"
	internalerrors "email-api/internal/internal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, req *http.Request) {
	var request contract.NewCampaign
	render.DecodeJSON(req.Body, &request)

	id, err := h.CampaignService.Create(request)
	if err != nil {

		if errors.Is(err, internalerrors.ErrInternal) {
			render.Status(req, 500)
		} else {
			render.Status(req, 400)
		}
		render.JSON(w, req, map[string]string{"error": err.Error()})
		return
	}
	render.Status(req, 201)
	render.JSON(w, req, map[string]string{"id": id})
}
