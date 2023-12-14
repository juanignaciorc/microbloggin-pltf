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
	user := domain.User{
		Name:  name,
		Email: mail,
	}

	createdUser, err := s.userRepository.CreateUser(user)
	if err != nil {
		return domain.User{}, err
	}

	return createdUser, nil
}

func (s *UserService) GetUser(id string) (domain.User, error) {
	user, err := s.userRepository.GetUser(id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *UserService) FollowUser(userID, followedID string) error {
	if err := s.userRepository.FollowUser(userID, followedID); err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUserTimeline(userID string) ([]domain.Tweet, error) {
	tweets, err := s.userRepository.GetUserTimeline(userID)
	if err != nil {
		return nil, err
	}

	return tweets, nil
}
