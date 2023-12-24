package handlers

import (
	"github.com/breeeaaad/WB-L0/internal/repository/cache"
)

type Handlers struct {
	c *cache.Cache
}

func New(c *cache.Cache) *Handlers {
	return &Handlers{
		c: c,
	}
}
