package storage

type IStorage interface {
}

type Storage struct {
	Books []Book `json:"books"`
}
