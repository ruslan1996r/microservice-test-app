package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"resources/internal/storage"
)

type IUserRepository interface {
	GetAll() []storage.User
}

type UsersHandler struct {
	usersRepo IUserRepository
}

func NewUsersHandler(userRepo IUserRepository) *UsersHandler {
	return &UsersHandler{
		userRepo,
	}
}

func (h *UsersHandler) InstallRoutes(r gin.IRouter) {
	r.GET("/api/v1/users", h.GetAll)
}

func (h *UsersHandler) GetAll(ctx *gin.Context) {
	users := h.usersRepo.GetAll()

	ctx.JSON(http.StatusOK, users)
}
