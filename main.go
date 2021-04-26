package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	var (
		key      = "abc"
		value int64 = 22
		vKey      = "cbd"
	)
	r, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer r.Close()

	req, err := r.Do("set", key, value)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(req)
	resp, err := redis.String(r.Do("get", vKey))
	if err != nil && err != redis.ErrNil {
		fmt.Println(err)
	}
	if err == redis.ErrNil {
		fmt.Println("No data")
	}

	fmt.Println(resp)

}