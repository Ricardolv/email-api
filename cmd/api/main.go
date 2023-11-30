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

	db := persistence.NewDb()
	campaignService := campaign.ServiceImp{
		Repository: &persistence.CampaignRepository{Db: db},
	}

	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	router.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	router.Get("/campaigns/{id}", endpoints.HandlerError(handler.CampaignGetById))
	router.Delete("/campaigns/{id}", endpoints.HandlerError(handler.CampaignDelete))

	// router.Route("/campaigns", func(r chi.Router) {
	// 	//r.Use(endpoints.Auth)
	// 	r.Post("/", endpoints.HandlerError(handler.CampaignPost))
	// 	r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetById))
	// 	r.Delete("/{id}", endpoints.HandlerError(handler.CampaignDelete))
	// })

	http.ListenAndServe(":3000", router)

}
