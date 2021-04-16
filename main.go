package main

import (
	"log"

	"github.com/charleyvibez/learnjwt/database"
	"github.com/charleyvibez/learnjwt/models"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return r
}

func main() {
	err := database.InitDatabase()
	if err != nil {
		log.Fatalln("could not create database", err)
	}

	database.GlobalDB.AutoMigrate(&models.User{})

	r := setupRouter()
	r.Run(":8080")
}
