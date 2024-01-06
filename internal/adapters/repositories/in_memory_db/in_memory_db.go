package in_memory_db

import "github.com/google/uuid"

type InMemoryDB struct {
	data map[uuid.UUID][]byte
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: make(map[uuid.UUID][]byte),
	}
}
