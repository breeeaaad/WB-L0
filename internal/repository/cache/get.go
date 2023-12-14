package cache

import (
	"errors"

	"github.com/breeeaaad/WB-L0/internal/models"
)

func (c *Cache) Get(uid string) (models.Order, error) {
	c.RLock()
	defer c.RUnlock()
	if order, found := c.order[uid]; found {
		return order, nil
	}
	return models.Order{}, errors.New("Order with that uid was not found")
}
