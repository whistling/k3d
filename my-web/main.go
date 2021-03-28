package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v4"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world, this is v2 version")
	})

	r.GET("/read", read)
	r.Run(":8080")
}

func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "123456",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

func read(c *gin.Context) {
	client := createClient()
	incr := client.Incr("read_times")
	c.String(200, "times: %d", incr)
}
