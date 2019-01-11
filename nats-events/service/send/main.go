package main

import (
	"fmt"
	"sandbox/nats-events/pb"
	"time"

	"github.com/nats-io/go-nats"
	natsp "github.com/nats-io/go-nats/encoders/protobuf"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
	}
	ec, err := nats.NewEncodedConn(nc, natsp.PROTOBUF_ENCODER)
	if err != nil {
		fmt.Println(err)
	}
	defer ec.Close()

	for i := 0; i < 5; i++ {
		myMessage := pb.TextMessage{
			Id:   int32(i),
			Body: "Hello over standard",
		}
		if err := ec.Publish("Messaging.Text.Standard", &myMessage); err != nil {
			fmt.Println(err)
		}
	}

	for i := 5; i < 10; i++ {
		myMessage := pb.TextMessage{
			Id:   int32(i),
			Body: "Hello, please respond!",
		}
		res := pb.TextMessage{}
		if err := ec.Request("Messaging.Text.Respond", &myMessage, &res, 200*time.Millisecond); err != nil {
			fmt.Println(err)
		}
		fmt.Println(res.Body, " with id ", res.Id)
	}

	sendChannel := make(chan *pb.TextMessage)
	ec.BindSendChan("Messaging.Text.Channel", sendChannel)
	for i := 10; i < 15; i++ {
		myMessage := pb.TextMessage{
			Id:   int32(i),
			Body: "Hello over channel!",
		}
		sendChannel <- &myMessage
	}
}
