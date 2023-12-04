package services

import (
	"github.com/juanignaciorc/microbloggin-pltf/internal/domain"
	"github.com/juanignaciorc/microbloggin-pltf/internal/ports"
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

func (s *UserService) FollowUser(userID, followedID int) error {
	if err := s.userRepository.FollowUser(userID, followedID); err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUserTimeline(userID int) ([]domain.Tweet, error) {
	tweets, err := s.userRepository.GetUserTimeline(userID)
	if err != nil {
		return nil, err
	}

	return tweets, nil
}
