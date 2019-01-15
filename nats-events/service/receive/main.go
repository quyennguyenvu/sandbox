package main

import (
	"fmt"
	"sandbox/nats-events/pb"

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

	ec.Subscribe("Messaging.Text.Standard", func(m *pb.TextMessage) {
		fmt.Println("Got standard message: \"", m.Body, "\" with the id ", m.Id)
	})

	ec.Subscribe("Messaging.Text.Respond", func(subject, reply string, m pb.TextMessage) {
		fmt.Println("Got ask for response message: \"", m.Body, "\" with the id ", m.Id)
		newMessage := pb.TextMessage{
			Id:   m.Id,
			Body: "Responding",
		}
		ec.Publish(reply, &newMessage)
	})

	receiveChannel := make(chan *pb.TextMessage)
	ec.BindRecvChan("Messaging.Text.Channel", receiveChannel)

	for m := range receiveChannel {
		fmt.Println("Got channel'ed message: \"", m.Body, "\" with the id ", m.Id)
	}
}
