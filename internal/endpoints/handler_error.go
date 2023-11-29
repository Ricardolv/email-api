package endpoints

import (
	internalerrors "email-api/internal/internal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type EndpointFunc func(w http.ResponseWriter, req *http.Request) (interface{}, int, error)

func HandlerError(endpointFunc EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		objt, status, err := endpointFunc(w, req)
		if err != nil {
			if errors.Is(err, internalerrors.ErrInternal) {
				render.Status(req, 500)
			} else {
				render.Status(req, 400)
			}
			render.JSON(w, req, map[string]string{"error": err.Error()})
			return
		}

		render.Status(req, status)

		if objt != nil {
			render.JSON(w, req, objt)
		}

	})
}
