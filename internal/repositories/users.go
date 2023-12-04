package repositories

import (
	"encoding/json"
	"fmt"

	"github.com/juanignaciorc/microbloggin-pltf/internal/domain"
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
	user.Follwing = append(user.Follwing, followedID)

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

func (db *InMemoryDB) GetUserTimeline(userID int) ([]domain.Tweet, error) {
	followedUsers, err := db.GetFollowedUsers(userID)
	if err != nil {
		return nil, err
	}

	var userTimeline []domain.Tweet
	for _, followedID := range followedUsers {
		user, err := db.GetUser(followedID)
		if err != nil {
			return nil, err
		}

		userTimeline = append(userTimeline, user.Tweets...)
	}

	return userTimeline, nil
}

func (db *InMemoryDB) GetFollowedUsers(userID int) ([]int, error) {
	user, err := db.GetUser(userID)
	if err != nil {
		return nil, err
	}

	return user.Follwing, nil
}
