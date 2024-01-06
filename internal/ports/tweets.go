package ports

import (
	"context"
	"github.com/juanignaciorc/microbloggin-pltf/internal/domain"
)

type TweetRepository interface {
	CreateTweet(ctx context.Context, tweet domain.Tweet) (domain.Tweet, error)
}

type TweetService interface {
	CreateTweet(userID string, message string) (domain.Tweet, error)
}
