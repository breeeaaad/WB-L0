package router

import (
	"github.com/breeeaaad/WB-L0/internal/router/handlers"
	"github.com/gin-gonic/gin"
)

func Router(h *handlers.Handlers) {
	r := gin.Default()
	r.GET("/get_order", h.GetOrder)
	r.Run()
}
