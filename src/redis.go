package main

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

var (
	redisConn redis.Conn
)

func OpenRedisConnection() {

	address := "localhost:6379"
	conn, e := redis.Dial("tcp", address)
	if e != nil {
		log.Fatal("REDIS  : service connection failed to start.")
	}
	log.Printf("REDIS  : Service connection started : Host=%s", address)
	redisConn = conn

}

func CloseRedisConnection() {

	e := redisConn.Close()
	if e != nil {
		log.Println("REDIS  : Service connection failed to close:", e)
	} else {
		log.Println("REDIS  : Service connection stopped.")
	}

}