package main

import (
	"fmt"
	"log"
	"sandbox/nats-basic/pb"

	"github.com/golang/protobuf/proto"
	nats "github.com/nats-io/go-nats"
)

var nc *nats.Conn

var users map[string]string

func main() {
	var err error
	nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
	}
	defer nc.Close()

	users = map[string]string{
		"1": "Bob",
		"2": "John",
		"3": "Dan",
		"4": "Kate",
	}

	nc.QueueSubscribe("UsernameById", "UsernameByProviders", replyWithUserID)
	select {}
}

func replyWithUserID(m *nats.Msg) {
	myUser := pb.User{}
	if err := proto.Unmarshal(m.Data, &myUser); err != nil {
		log.Println(err)
		return
	}

	myUser.Name = users[myUser.Id]
	data, err := proto.Marshal(&myUser)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Replying to ", m.Reply)
	nc.Publish(m.Reply, data)
}
