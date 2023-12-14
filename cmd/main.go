package main

import (
	"context"
	"log"

	"github.com/breeeaaad/WB-L0/internal/configs"
	"github.com/breeeaaad/WB-L0/internal/repository/cache"
	"github.com/breeeaaad/WB-L0/internal/repository/database"
	"github.com/breeeaaad/WB-L0/internal/router"
	"github.com/breeeaaad/WB-L0/internal/router/handlers"
)

func main() {
	c := configs.Dbconfig()
	defer func() {
		if err := c.Close(context.Background()); err != nil {
			log.Fatalf("Error while closing connection to the postgre server: %s", err.Error())
		}
	}()
	d := database.New(c)
	r := cache.New(d)
	h := handlers.New(r, d)
	router.Router(h)
}
