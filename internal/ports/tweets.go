package ports

import "github.com/juanignaciorc/microbloggin-pltf/internal/domain"

type TweetRepository interface {
	CreateTweet(tweet domain.Tweet) (domain.Tweet, error)
}

type TweetService interface {
	CreateTweet(userID int, message string) (domain.Tweet, error)
}
