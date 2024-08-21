package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CheckTokenRequest struct {
	Token string `json:"token"`
}

type IAuthService interface {
	Token() (string, error)
	CheckToken(token string) (bool, error)
}

type AuthHandlers struct {
	authService IAuthService
}

func NewAuthHandlers(authService IAuthService) *AuthHandlers {
	return &AuthHandlers{
		authService,
	}
}

func (h *AuthHandlers) InstallRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/token", h.Token)
	mux.HandleFunc("/api/v1/check_token", h.CheckToken)
}

func (h *AuthHandlers) Token(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		token, err := h.authService.Token()
		if err != nil {
			http.Error(w, "Token generation error", http.StatusBadRequest)
			return
		}

		_, err = fmt.Fprintf(w, token)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func (h *AuthHandlers) CheckToken(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var checkTokenReq CheckTokenRequest
		err = json.Unmarshal(body, &checkTokenReq)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		_, err = h.authService.CheckToken(checkTokenReq.Token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "true")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
