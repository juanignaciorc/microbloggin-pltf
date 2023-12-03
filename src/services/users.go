package services

import (
	"github.com/juanignaciorc/microbloggin-pltf/src/domain"
	"github.com/juanignaciorc/microbloggin-pltf/src/ports"
)

type UserService struct {
	userRepository ports.UsersRepository
}

func NewUserService(userRepository ports.UsersRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(name, mail string) error {
	user := domain.NewUser(name, mail)
	if err := s.userRepository.CreateUser(user); err != nil {
		return err
	}

	return nil
}
