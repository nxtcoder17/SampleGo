package routes

import (
	"SampleCrud/db"
	"SampleCrud/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var router *gin.Engine

func Init() {
	router = gin.Default()
	router.POST("/register", func(c *gin.Context) {
		var user db.User
		err := c.BindJSON(&user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
		data := services.Register(user)
		fmt.Println(data)
		c.JSON(200, gin.H{
			"msg": "hello",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen
}

