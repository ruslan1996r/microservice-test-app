package handlers

import (
	"fmt"
	"net/http"
)

type IAuthService interface {
	Login(path string) (string, error)
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
	mux.HandleFunc("/api/v1/users/sign_in", h.Login)
}

func (h *AuthHandlers) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		token, err := h.authService.Login("/api/v1/token")
		if err != nil {
			http.Error(w, "Login error", http.StatusBadRequest)
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
