package endpoints

import (
	"context"
	"net/http"
	"strings"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	jwtgo "github.com/dgrijalva/jwt-go"

	"github.com/go-chi/render"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		tokenString := req.Header.Get("Authorization")
		if tokenString == "" {
			render.Status(req, 401)
			render.JSON(w, req, map[string]string{"error": "request does not contain an authorization header"})
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		provider, err := oidc.NewProvider(req.Context(), "http://localhost:8180/auth/realms/provider")
		if err != nil {
			render.Status(req, 500)
			render.JSON(w, req, map[string]string{"error": "error to connect to the provider"})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: "email-api"})
		_, err = verifier.Verify(req.Context(), tokenString)
		if err != nil {
			render.Status(req, 401)
			render.JSON(w, req, map[string]string{"error": "invalid token"})
			return
		}

		token, _ := jwtgo.Parse(tokenString, nil)
		claims := token.Claims.(jwtgo.MapClaims)
		email := claims["email"]

		ctx := context.WithValue(req.Context(), "email", email)

		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
