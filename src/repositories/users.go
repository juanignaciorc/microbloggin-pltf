package repositories

import (
	"encoding/json"

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

func (db *InMemoryDB) CreateUser(user domain.User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	db.data[user.ID] = userBytes
	return nil
}
