package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"gateway/internal/services"
)

const AuthUrl = "AUTH_URL"

func AuthOnly(next http.HandlerFunc, fetchService services.IFetchService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Unauthorized - No Token Provided", http.StatusUnauthorized)
			return
		}

		url := fmt.Sprintf("%s%s", os.Getenv(AuthUrl), "/api/v1/check_token")
		rawPayload := fmt.Sprintf(`{"token":"%s"}`, token)
		tokenPayload := []byte(rawPayload)

		tokenStatus, err := fetchService.Post(url, tokenPayload)

		if err != nil {
			http.Error(w, "Token validation error", http.StatusUnauthorized)
			return
		}

		if tokenStatus != "true" {
			http.Error(w, "Invalid token provided", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
