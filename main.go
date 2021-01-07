package main

import (
	"fmt"
	"net"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Printf("Error when sending Dial with message: %v", err.Error())
	}

	redisConnection := redis.NewConn(connection, 60*time.Second, 5*time.Second)

	fmt.Println("Sending AUTH command")

	if _, err := redisConnection.Do("AUTH", "callumosborn"); err != nil {
		fmt.Printf("Error when sending AUTH command with message: %v", err.Error())
		redisConnection.Close()
	}

	fmt.Println("Sending SET Command")

	if _, err := redisConnection.Do("SET", "connections", "10"); err != nil {
		fmt.Printf("Error when sending SET command with message: %v", err.Error())
		redisConnection.Close()
	}

	err = redisConnection.Close()

	if err != nil {
		fmt.Printf("Error when attempting to close redis connection with message: %v", err.Error())
	}

	err = connection.Close()

	if err != nil {
		fmt.Printf("Error when attempting to close connection with message: %v", err.Error())
	}

	fmt.Println("Connection closed")
}
