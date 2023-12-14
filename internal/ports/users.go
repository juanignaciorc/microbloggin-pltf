package ports

import "github.com/juanignaciorc/microbloggin-pltf/internal/domain"

type UsersRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	GetUser(id string) (domain.User, error)
	FollowUser(userID string, followedID string) error
	GetUserTimeline(userID string) ([]domain.Tweet, error)
}

type UserService interface {
	CreateUser(name, mail string) (domain.User, error)
	GetUser(id string) (domain.User, error)
	FollowUser(userID, followedID string) error
	GetUserTimeline(userID string) ([]domain.Tweet, error)
}
