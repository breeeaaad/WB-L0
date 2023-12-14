package cache

import (
	"sync"

	"github.com/breeeaaad/WB-L0/internal/models"
	"github.com/breeeaaad/WB-L0/internal/repository/database"
)

type Cache struct {
	sync.RWMutex
	order map[string]models.Order
	db    *database.Database
}

func New(db *database.Database) *Cache {
	return &Cache{
		order: make(map[string]models.Order),
		db:    db,
	}
}
