package nats

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/breeeaaad/WB-L0/internal/models"
	"github.com/breeeaaad/WB-L0/internal/repository/database"
	"github.com/nats-io/stan.go"
)

func Subscribe(sc stan.Conn, wg *sync.WaitGroup, subject string, d *database.Database) error {
	var order models.Order
	sub, err := sc.Subscribe(subject, func(msg *stan.Msg) {
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			log.Print("error with unmarshaling json")
		}
		wg.Done()
	})
	if err != nil {
		return err
	}
	wg.Wait()
	if err := d.Create(order); err != nil {
		return err
	}
	d.C.Set(order)
	err = sub.Unsubscribe()
	if err != nil {
		return err
	}
	return nil
}
