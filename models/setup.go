package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDatabase(){
db,error := gorm.Open(mysql.Open("root:@tcp(127.0.0.1)/library?parseTime=true"))
	if error!= nil{
		panic(error)
	}

	// Make a table book if there no table book
	db.AutoMigrate(&Buku{})
	DB = db
}