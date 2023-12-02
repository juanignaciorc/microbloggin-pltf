package ports

import (
	"context"

	"github.com/juanignaciorc/microbloggin-pltf/internal/core/domain"
)

type UsersRepository interface {
	GetUser(id int) (domain.User, error)
	CreateUser(user domain.User) error
}

type UserService interface {
	GetUser(ctx context.Context, id int) (domain.User, error)
	CreateUser(id int, name string) error
}