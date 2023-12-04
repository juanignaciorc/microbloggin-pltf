package repositories

import (
	"encoding/json"
	"github.com/juanignaciorc/microbloggin-pltf/internal/domain"
)

func (db *InMemoryDB) CreateTweet(tweet domain.Tweet) (domain.Tweet, error) {
	userID := tweet.UserID
	user, err := db.GetUser(userID)
	if err != nil {
		return domain.Tweet{}, err
	}

	user.Tweets = append(user.Tweets, tweet)

	userBytes, err := json.Marshal(user)
	if err != nil {
		return domain.Tweet{}, err
	}

	db.data[userID] = userBytes

	return tweet, nil
}
