package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sandbox/nats-streaming-events/pb"

	stan "github.com/nats-io/go-nats-streaming"
)

var conn stan.Conn

var users map[string]string

func logCloser(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("close error: %s", err)
	}
}

func main() {
	conn, err := stan.Connect(
		"test-cluster",
		"111",
		stan.NatsURL("nats://localhost:4222"),
	)
	if err != nil {
		log.Print(err)
	}
	defer logCloser(conn)

	sub, _ := conn.Subscribe("foo", func(msg *stan.Msg) {
		// Print the value and whether it was redelivered.
		fmt.Printf("seq = %d [redelivered = %v]\n", msg.Sequence, msg.Redelivered)
		message := pb.TextMessage{}
		if err := json.Unmarshal(msg.Data, &message); err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Got standard message: \"", message.Body, "\" with the id ", message.Id)
	}, stan.DurableName("i-will-remember"), stan.SetManualAckMode())
	if err != nil {
		log.Print(err)
	}

	defer logCloser(sub)

	select {}
}
