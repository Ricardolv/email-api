package endpoints

import (
	"net/http"
	"strings"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/render"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		token := req.Header.Get("Authorization")
		if token == "" {
			render.Status(req, 401)
			render.JSON(w, req, map[string]string{"error": "request does not contain an authorization header"})
			return
		}

		token = strings.Replace(token, "Bearer ", "", 1)
		provider, err := oidc.NewProvider(req.Context(), "http://localhost:8180/auth/realms/provider")
		if err != nil {
			render.Status(req, 500)
			render.JSON(w, req, map[string]string{"error": "error to connect to the provider"})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: "email-api"})
		_, err = verifier.Verify(req.Context(), token)
		if err != nil {
			render.Status(req, 401)
			render.JSON(w, req, map[string]string{"error": "invalid token"})
			return
		}

		next.ServeHTTP(w, req)
	})
}
