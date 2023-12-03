package ports

import "github.com/juanignaciorc/microbloggin-pltf/src/domain"

type UsersRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	GetUser(id int) (domain.User, error)
}

type UserService interface {
	CreateUser(name, mail string) (domain.User, error)
	GetUser(id int) (domain.User, error)
}
