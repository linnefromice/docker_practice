package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
	"os"
)

func redisConnect() redis.Conn {
	const IP_PORT = "127.0.0.1:6379"
	
	// Redisに接続
	c, err := redis.Dial("tcp", IP_PORT)
	if err != nil {
		panic(err)
	}
	return c
}

func redisSet(key string, value string, c redis.Conn) {
	c.Do("SET", key, value)
}

func redisGet(key string, c redis.Conn) string {
	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return s
}

func main() {
	c := redisConnect()
	defer c.Close()
	var key = "KEY_SAMPLE"
	var value = "VALUE_SAMPLE"
	redisSet(key, value, c)
	output := redisGet(key, c)
	fmt.Println(output)
}
