package storage

import "github.com/bxcodec/faker/v4"

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BooksRepository struct {
	Books []Book
}

func NewBooksRepository() *BooksRepository {
	return &BooksRepository{}
}

func (r *BooksRepository) GetAll() []Book {
	return r.Books
}

func (r *BooksRepository) CreateMocks(amount int) []Book {
	books := make([]Book, amount)

	for i := 0; i < amount; i++ {
		books[i] = Book{
			ID:          i + 1,
			Title:       faker.Word(),
			Description: faker.Sentence(),
		}
	}

	r.Books = books

	return books
}
