package main

import (
	"email-api/internal/contract"
	"email-api/internal/domain/campaign"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	service := campaign.Service{}

	r.Post("/campaigns", func(w http.ResponseWriter, req *http.Request) {
		var request contract.NewCampaign

		err := render.DecodeJSON(req.Body, &request)
		if err != nil {
			println(err)
		}

		id, err := service.Create(request)
		if err != nil {
			render.Status(req, 400)
			render.JSON(w, req, map[string]string{"error": err.Error()})
			return
		}

		render.Status(req, 201)
		render.JSON(w, req, map[string]string{"id": id})
	})

	http.ListenAndServe(":3000", r)

}
