package handlers

import (
	"github.com/breeeaaad/WB-L0/internal/repository/cache"
	"github.com/breeeaaad/WB-L0/internal/repository/database"
)

type Handlers struct {
	c *cache.Cache
	d *database.Database
}

func New(c *cache.Cache, d *database.Database) *Handlers {
	return &Handlers{
		c: c,
		d: d,
	}
}
