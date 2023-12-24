package main

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", uuid.NewString(), stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() {
		err := sc.Close()
		if err != nil {
			log.Fatalf("Error while closing connection to the nuts-streaming server: %s", err.Error())
		}
	}()
	data, err := os.ReadFile("publisher/model.json")
	if err != nil {
		log.Print(err.Error())
	}
	if err := sc.Publish("order", data); err != nil {
		log.Print(err.Error())
	}
}
