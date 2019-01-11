package main

import (
	"fmt"
	"log"
	"sandbox/nats-basic/pb"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
)

var nc *nats.Conn

func main() {
	var err error
	nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
	}
	defer nc.Close()

	nc.QueueSubscribe("TimeTeller", "TimeTellers", replyWithTime)
	select {}
}

func replyWithTime(m *nats.Msg) {
	currTime := pb.Time{
		Time: time.Now().Format(time.RFC3339),
	}
	data, err := proto.Marshal(&currTime)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Reply to ", m.Reply)
	nc.Publish(m.Reply, data)
}
