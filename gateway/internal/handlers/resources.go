package handlers

import (
	"encoding/json"
	"gateway/internal/middlewares"
	"gateway/internal/services"
	"gateway/internal/types"
	"net/http"
)

type IResourcesService interface {
	GetUsers(path string) (*[]types.User, error)
	GetBooks(path string) (*[]types.Book, error)
}

type ResourcesHandlers struct {
	resourcesService IResourcesService
	FetchService     services.IFetchService
}

func NewResourcesHandlers(
	resourcesService IResourcesService,
	fetchService services.IFetchService,
) *ResourcesHandlers {
	return &ResourcesHandlers{
		resourcesService,
		fetchService,
	}
}

func (h *ResourcesHandlers) InstallRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/users", middlewares.AuthOnly(h.GetUsers, h.FetchService))
	mux.HandleFunc("/api/v1/books", h.GetBooks)
}

func (h *ResourcesHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.resourcesService.GetUsers("/api/v1/users")
	if err != nil {
		http.Error(w, "Get users error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}
}

func (h *ResourcesHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.resourcesService.GetBooks("/api/v1/books")
	if err != nil {
		http.Error(w, "Get books error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}
}
