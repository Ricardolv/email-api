package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, req *http.Request) {

	render.Status(req, 200)
	render.JSON(w, req, h.CampaignService.Repository.FindAll())
}
