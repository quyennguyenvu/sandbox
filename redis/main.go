package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
)

const prefix string = "user:"

// User is a simple user struct for this example
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	pool := newPool()
	conn := pool.Get()
	defer conn.Close()

	if len(os.Args) != 3 {
		log.Printf("not enough parameters, now has %d\n", len(os.Args))
		return
	}
	if err := ping(conn); err != nil {
		log.Println(err)
	}

	key := os.Args[1]
	val := os.Args[2]
	if err := set(conn, key, val); err != nil {
		log.Println(err)
	}
	result, err := get(conn, key)
	if err != nil {
		log.Println(err)
	}
	log.Println(key + " has password: " + result)

	// struct
	user := User{
		Username: key,
		Password: val,
	}
	if err = setStruct(conn, &user); err != nil {
		log.Println(err)
	}
	userData, err := getStruct(conn, key)
	if err != nil {
		log.Println(err)
	}
	log.Println("User data:")
	log.Printf("%+v\n", userData)

}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				log.Fatal(err)
			}
			return c, err
		},
	}
}

func ping(c redis.Conn) error {
	// Send PING command to Redis
	pong, err := c.Do("PING")
	if err != nil {
		return err
	}
	s, err := redis.String(pong, err)
	if err != nil {
		return err
	}
	log.Printf("PING response = %s\n", s)
	return nil
}

func set(c redis.Conn, key, val string) error {
	if _, err := c.Do("SET", key, val); err != nil {
		return err
	}
	return nil
}

func get(c redis.Conn, search string) (string, error) {
	s, err := redis.String(c.Do("GET", search))
	if err != nil {
		return "", err
	}
	return s, nil
}

// Struct
func setStruct(c redis.Conn, user *User) error {
	json, err := json.Marshal(&user)
	if err != nil {
		return err
	}
	if _, err = c.Do("SET", prefix+user.Username, json); err != nil {
		return err
	}
	return nil
}

func getStruct(c redis.Conn, username string) (user *User, err error) {
	s, err := redis.String(c.Do("GET", prefix+username))
	if err == redis.ErrNil {
		log.Println("User does not exist")
	} else if err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(s), &user); err != nil {
		return nil, err
	}
	return user, nil
}
