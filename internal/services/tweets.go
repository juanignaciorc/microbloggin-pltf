package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/juanignaciorc/microbloggin-pltf/internal/domain"
	ports "github.com/juanignaciorc/microbloggin-pltf/internal/ports/repositories"
)

type TweetsService struct {
	tweetsRepository ports.TweetRepository
}

func NewTweetsService(tweetsRepository ports.TweetRepository) TweetsService {
	return TweetsService{
		tweetsRepository: tweetsRepository,
	}
}

func (s *TweetsService) CreateTweet(ctx context.Context, userID uuid.UUID, message string) (domain.Tweet, error) {
	tweet := domain.Tweet{
		UserID:  userID,
		Message: message,
	}
	if _, err := s.tweetsRepository.CreateTweet(ctx, tweet); err != nil {
		return domain.Tweet{}, err
	}

	return tweet, nil
}
