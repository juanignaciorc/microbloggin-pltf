package domain

import (
	"math/rand"
)

type User struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Followers []int   `json:"followers"`
	Follwing  []int   `json:"following"`
	Tweets    []Tweet `json:"tweets"`
}

func NewUser(name, email string) User {
	id := rand.Intn(1000)

	return User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}