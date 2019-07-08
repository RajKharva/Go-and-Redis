package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main(){

	conn, err := redis.Dial("tcp", ":6379")
	if err != nil{
		panic(err)
	}

	defer conn.Close()

	conn.Do("SET", "counter", 1)

	countVal, _ := redis.Int(conn.Do("GET", "counter"))
	fmt.Printf("Counter Value: %#v\n", countVal)

	conn.Do("INCR", "counter")

	countVal, _ = redis.Int(conn.Do("GET", "counter"))
	fmt.Printf("Counter Value after INCR operation: %#v\n", countVal)
}
