package main

import (
	"context"
	"log"
	"sync"

	"github.com/breeeaaad/WB-L0/internal/configs"
	"github.com/breeeaaad/WB-L0/internal/nats"
	"github.com/breeeaaad/WB-L0/internal/repository/database"
	"github.com/breeeaaad/WB-L0/internal/router"
	"github.com/breeeaaad/WB-L0/internal/router/handlers"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", uuid.NewString(), stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Print(err.Error())
	}
	defer func() {
		err := sc.Close()
		if err != nil {
			log.Print(err.Error())
		}
	}()
	c := configs.Dbconfig()
	defer func() {
		if err := c.Close(context.Background()); err != nil {
			log.Print(err.Error())
		}
	}()
	d := database.New(c)
	h := handlers.New(d.C)
	if err := d.Alldata(); err != nil {
		log.Print(err.Error())
	}
	if sc != nil {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			err := nats.Subscribe(sc, &wg, "order", d)
			if err != nil {
				log.Print(err.Error())
			}
		}()
	}
	router.Router(h)
}
