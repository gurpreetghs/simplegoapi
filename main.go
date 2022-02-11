package main

import (
	"github.com//gurpreetghs/simplegoapi/gin-bookstore/controllers"
	"github.com//gurpreetghs/simplegoapi/gin-bookstore/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	
	models.ConnectDatabase()


	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	
	r.Run()
}
