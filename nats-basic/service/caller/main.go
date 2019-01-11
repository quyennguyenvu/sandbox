package main

import (
	"fmt"
	"net/http"
	"sandbox/nats-basic/pb"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
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

	m := mux.NewRouter()
	m.HandleFunc("/{id}", handleUserWithTime)
	http.ListenAndServe(":3000", m)
}

func handleUserWithTime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	myUser := pb.User{Id: vars["id"]}
	currTime := pb.Time{}
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		data, err := proto.Marshal(&myUser)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			fmt.Println("Problem with parsing the user id")
			return
		}
		msg, err := nc.Request("UsernameById", data, 100*time.Millisecond)
		if err == nil && msg != nil {
			myUserWithName := pb.User{}
			if err := proto.Unmarshal(msg.Data, &myUserWithName); err == nil {
				myUser = myUserWithName
			}
		}

		wg.Done()
	}()

	go func() {
		msg, err := nc.Request("TimeTeller", nil, 100*time.Millisecond)
		if err == nil && msg != nil {
			receivedTime := pb.Time{}
			if err := proto.Unmarshal(msg.Data, &receivedTime); err == nil {
				currTime = receivedTime
			}
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Fprintln(w, "Hello ", myUser.Name, " with id ", myUser.Id, ", the time is ", currTime.Time)
}
