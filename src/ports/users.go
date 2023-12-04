package ports

import "github.com/juanignaciorc/microbloggin-pltf/src/domain"

type UsersRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	GetUser(id int) (domain.User, error)
	CreateTweet(tweet domain.Tweet) (domain.Tweet, error)
	FollowUser(userID, followedID int) error
	GetUserTimeline(userID int) ([]domain.Tweet, error)
}

type UserService interface {
	CreateUser(name, mail string) (domain.User, error)
	GetUser(id int) (domain.User, error)
	CreateTweet(userID int, message string) (domain.Tweet, error)
	FollowUser(userID, followedID int) error
	GetUserTimeline(userID int) ([]domain.Tweet, error)
}
