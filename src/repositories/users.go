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
