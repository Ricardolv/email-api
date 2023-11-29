package main

import (
	"email-api/internal/domain/campaign"
	"email-api/internal/endpoints"
	"email-api/internal/infrastructure/persistence"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	campaignService := campaign.Service{
		Repository: &persistence.CampaignRepository{},
	}

	handler := endpoints.Handler{
		CampaignService: campaignService,
	}

	router.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	router.Get("/campaigns", endpoints.HandlerError(handler.CampaignGet))

	http.ListenAndServe(":3000", router)

}