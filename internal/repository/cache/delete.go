package cache

import "errors"

func (c *Cache) Delete(uid string) error {
	c.Lock()
	defer c.Unlock()
	if _, found := c.order[uid]; !found {
		return errors.New("Order uid not found")
	}
	delete(c.order, uid)
	return nil
}
