package main

import (
	"encoding/json"
	"log"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")

	if err != nil {
		log.Fatal(err)
	}

	update, _ := json.Marshal(UpdateMessage{21})

	_, err = c.Do("PUBLISH", "tempoo-update", update)

	if err != nil {
		log.Fatal(err)
	}
}

type UpdateMessage struct {
	Temperature int
}
