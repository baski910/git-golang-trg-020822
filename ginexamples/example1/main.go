package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	fmt.Println("server started at 8010")
	log.Fatal(router.Run(":8010"))
}
