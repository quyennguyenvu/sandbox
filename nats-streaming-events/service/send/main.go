package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sandbox/nats-streaming-events/pb"

	stan "github.com/nats-io/go-nats-streaming"
)

func main() {
	sc, err := stan.Connect(
		"test-cluster",
		"test-client",
	)
	if err != nil {
		log.Println(err)
		return
	}
	defer logCloser(sc)

	for i := 0; i < 5; i++ {
		myMessage := pb.TextMessage{
			Id:   int32(i),
			Body: "Hello over standard",
		}
		msg, err := json.Marshal(&myMessage)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		if err := sc.Publish("foo", msg); err != nil {
			fmt.Println(err)
		}
	}
}

func logCloser(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("close error: %s", err)
	}
}
