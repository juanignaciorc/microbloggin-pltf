package postgre_db

import (
	"context"
	"github.com/google/uuid"
	"github.com/juanignaciorc/microbloggin-pltf/internal/domain"
	"log"
)

type TweetRepository struct {
	db *DB
}

func NewTweetRepository(db *DB) *TweetRepository {
	return &TweetRepository{
		db,
	}
}

func (tr *TweetRepository) CreateTweet(ctx context.Context, tweet domain.Tweet) (domain.Tweet, error) {
	tweet.ID = uuid.New()

	result, err := tr.db.connPool.Exec(ctx, "INSERT INTO tweets (id, user_id, content) VALUES ($1, $2, $3)", tweet.ID, tweet.UserID, tweet.Message)
	if err != nil {
		return domain.Tweet{}, err
	}

	rowsAffected := result.RowsAffected()
	if err != nil {
		return domain.Tweet{}, err
	}

	if rowsAffected != 1 {
		log.Printf("Expected to affect 1 row, affected %d", rowsAffected)
		return domain.Tweet{}, err
	}
	return tweet, nil
}
