package services

import (
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

func (s *TweetsService) CreateTweet(userID int, message string) (domain.Tweet, error) {
	tweet := domain.NewTweet(userID, message)
	if _, err := s.tweetsRepository.CreateTweet(tweet); err != nil {
		return domain.Tweet{}, err
	}

	return tweet, nil
}
