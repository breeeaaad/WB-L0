package handlers

import (
	"github.com/gin-gonic/gin"
)

type uri struct {
	uid string `uri:"uid" binding:"required"`
}

func (h *Handlers) GetOrder(c *gin.Context) {
	var uid uri
	if err := c.BindUri(&uid); err != nil {
		c.JSON(400, err.Error())
		return
	}
	order, err := h.c.Get(uid.uid)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, order)
}
