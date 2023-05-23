package controller

import (
	"fmt"
	"library-api/models"

	"github.com/gin-gonic/gin"
)

// Routing for showing all data book
func GetAllBook(c *gin.Context){
	var buku []models.Buku
	models.DB.Find(&buku)
	c.JSON(200,buku)
}



// Routing for searching data book based on ID
func GetBookId(c *gin.Context) {
 id := c.Param("id")
 var buku models.Buku
 if err := models.DB.Where("id = ?", id).First(&buku).Error; err != nil {
	c.JSON(404, gin.H{"message" : "the book you are looking for doesn't exist"})
 } else {
    c.JSON(200, buku)
 }
}


// Routing for adding a new data book
func CreateBook(c *gin.Context){
	var buku models.Buku
	c.BindJSON(&buku)
	models.DB.Create(&buku)
	c.JSON(200,buku)
	
}


//Routing for update data book
func UpdateBook(c *gin.Context) {
 var buku models.Buku
 id := c.Param("id")
 if err := models.DB.Where("id = ?", id).First(&buku).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
}
 c.BindJSON(&buku)
 models.DB.Save(&buku)
 c.JSON(200, gin.H{"message" : "Buku Update succesfully"})
 c.JSON(200, buku)
 
}


// Routing for delete data book through ID
func DeleteBook(c *gin.Context){
	id := c.Param("id")
	var buku models.Buku
	models.DB.First(&buku,id)
	if buku.ID != 0 {
		models.DB.Delete(&buku)
		c.JSON(200, gin.H{"message" : "buku" + id + "deleted succesfully"})
	}else{
		c.JSON(404, gin.H{"message" : "Oops Sorry, buku not found"})
	}
}