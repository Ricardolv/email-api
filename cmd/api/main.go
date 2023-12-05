package main

import (
	"email-api/internal/domain/campaign"
	"email-api/internal/endpoints"
	"email-api/internal/infrastructure/mail"
	"email-api/internal/infrastructure/persistence"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	db := persistence.NewDb()
	campaignService := campaign.ServiceImp{
		Repository: &persistence.CampaignRepository{Db: db},
		SendMail:   mail.SendMail,
	}

	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	router.Route("/campaigns", func(r chi.Router) {
		r.Use(endpoints.Auth)
		r.Post("/", endpoints.HandlerError(handler.CampaignPost))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetById))
		r.Delete("/{id}", endpoints.HandlerError(handler.CampaignDelete))
		r.Patch("/{id}", endpoints.HandlerError(handler.CampaignStart))
	})

	http.ListenAndServe(":3000", router)

}
