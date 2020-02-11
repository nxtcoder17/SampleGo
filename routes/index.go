package routes

import (
	"SampleCrud/db"
	"SampleCrud/services"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init() {
	router = gin.Default()
	router.POST("/register", func(c *gin.Context) {
		// Fetching keys from context
		user := db.User{
			Name: c.PostForm("name"),
			Email: c.PostForm("email"),
			Password: c.PostForm("password"),
		};
		services.Register(user)
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen
}

