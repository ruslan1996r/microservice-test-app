package storage

import (
	"github.com/bxcodec/faker/v4"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRepository struct {
	Users []User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetAll() []User {
	return r.Users
}

func (r *UserRepository) CreateMocks(amount int) []User {
	users := make([]User, amount)

	for i := 0; i < amount; i++ {
		users[i] = User{
			ID:   i + 1,
			Name: faker.Name(),
		}
	}

	r.Users = users

	return users
}
