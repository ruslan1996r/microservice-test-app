package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"resources/internal/storage"
)

type IBookRepository interface {
	GetAll() []storage.Book
}

type BooksHandler struct {
	booksRepo IBookRepository
}

func NewBooksHandler(booksRepo IBookRepository) *BooksHandler {
	return &BooksHandler{
		booksRepo,
	}
}

func (h *BooksHandler) InstallRoutes(r gin.IRouter) {
	r.GET("/api/v1/books", h.GetAll)
}

func (h *BooksHandler) GetAll(ctx *gin.Context) {
	books := h.booksRepo.GetAll()

	ctx.JSON(http.StatusOK, books)
}
