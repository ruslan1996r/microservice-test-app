package types

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
