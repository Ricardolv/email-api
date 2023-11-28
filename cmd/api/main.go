package main

import (
	"email-api/internal/contract"
	"email-api/internal/domain/campaign"
	"email-api/internal/infrastructure/persistence"
	internalerrors "email-api/internal/internal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	service := campaign.Service{
		Repository: &persistence.CampaignRepository{},
	}

	router.Post("/campaigns", func(w http.ResponseWriter, req *http.Request) {
		var request contract.NewCampaign
		render.DecodeJSON(req.Body, &request)

		println("name: ", request.Name)
		println("content: ", request.Content)
		println("emails: ", len(request.Emails))

		id, err := service.Create(request)
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
	})

	http.ListenAndServe(":3000", router)

}
