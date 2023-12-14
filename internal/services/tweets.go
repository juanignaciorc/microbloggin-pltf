package services

import (
	"github.com/google/uuid"
	"github.com/juanignaciorc/microbloggin-pltf/internal/domain"
	"github.com/juanignaciorc/microbloggin-pltf/internal/ports"
)

type TweetsService struct {
	tweetsRepository ports.TweetRepository
}

func NewTweetsService(tweetsRepository ports.TweetRepository) TweetsService {
	return TweetsService{
		tweetsRepository: tweetsRepository,
	}
}

func (s *TweetsService) CreateTweet(userID string, message string) (domain.Tweet, error) {
	id := uuid.NewString()

	tweet := domain.Tweet{
		ID:      id,
		UserID:  userID,
		Message: message,
	}
	if _, err := s.tweetsRepository.CreateTweet(tweet); err != nil {
		return domain.Tweet{}, err
	}

	return tweet, nil
}
