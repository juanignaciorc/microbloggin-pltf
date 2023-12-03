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

func (s *UserService) CreateUser(name, mail string) (domain.User, error) {
	user := domain.NewUser(name, mail)
	if _, err := s.userRepository.CreateUser(user); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *UserService) GetUser(id int) (domain.User, error) {
	user, err := s.userRepository.GetUser(id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *UserService) CreateTweet(userID int, message string) (domain.Tweet, error) {
	tweet := domain.NewTweet(userID, message)
	if _, err := s.userRepository.CreateTweet(tweet); err != nil {
		return domain.Tweet{}, err
	}

	return tweet, nil
}
