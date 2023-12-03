package domain

import (
	"math/rand"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Followers []int  `json:"followers"`
}

func NewUser(name, email string) User {
	id := rand.Intn(1000)

	return User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}