package cache

import (
	"sync"

	"github.com/breeeaaad/WB-L0/internal/models"
)

type Cache struct {
	sync.RWMutex
	order map[string]models.Order
}

func New() *Cache {
	return &Cache{
		order: make(map[string]models.Order),
	}
}
