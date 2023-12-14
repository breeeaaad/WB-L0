package router

import (
	"github.com/breeeaaad/WB-L0/internal/router/handlers"
	"github.com/gin-gonic/gin"
)

func Router(h *handlers.Handlers) {
	r := gin.Default()
	r.GET("/main/order/:order_uid", h.GetOrder)
	r.Run()
}
