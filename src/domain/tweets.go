package domain

import "math/rand"

type Tweet struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Message string `json:"message" validate:"max=140"`
}

func NewTweet(userID int, message string) Tweet {
	id := rand.Intn(1000)
	return Tweet{
		ID:      id,
		UserID:  userID,
		Message: message,
	}
}
