package main

import (
	"library-api/controller"
	"library-api/models"

	"github.com/gin-gonic/gin"
)


func main(){
	r := gin.Default()
	models.ConnectDatabase()
 	r.GET("/library/", controller.GetAllBook)
 	r.GET("/library/:id", controller.GetBookId)
	r.POST("library/", controller.CreateBook)
	r.PUT("/library/:id",controller.UpdateBook)
	r.DELETE("/library/:id",controller.DeleteBook)
//Run server on port 8080
	r.Run(":8080")

}


