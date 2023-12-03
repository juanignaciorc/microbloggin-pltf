package repositories

import (
	"encoding/json"
	"fmt"

	"github.com/juanignaciorc/microbloggin-pltf/src/domain"
)

type InMemoryDB struct {
	data map[int][]byte
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: make(map[int][]byte),
	}
}

func (db *InMemoryDB) CreateUser(user domain.User) (domain.User, error) {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return domain.User{}, err
	}

	db.data[user.ID] = userBytes
	var createdUser domain.User
	err = json.Unmarshal(userBytes, &createdUser)
	return createdUser, err
}

func (db *InMemoryDB) GetUser(id int) (domain.User, error) {
	userBytes, ok := db.data[id]
	if !ok {
		return domain.User{}, fmt.Errorf("user with id %d not found", id)
	}

	var user domain.User
	err := json.Unmarshal(userBytes, &user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

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

func (db *InMemoryDB) FollowUser(userID, followedID int) error {
	user, err := db.GetUser(userID)
	if err != nil {
		return err
	}

	followedUser, err := db.GetUser(followedID)
	if err != nil {
		return err
	}

	followedUser.Followers = append(user.Followers, userID)
	user.Followeds = append(user.Followeds, followedID)

	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	followedUserBytes, err := json.Marshal(followedUser)
	if err != nil {
		return err
	}

	db.data[userID] = userBytes
	db.data[followedID] = followedUserBytes

	return nil
}
