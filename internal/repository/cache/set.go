package cache

import (
	"github.com/breeeaaad/WB-L0/internal/models"
)

func (c *Cache) Set(order models.Order) {
	c.Lock()
	defer c.Unlock()
	c.order[order.OrderUid] = order
}
