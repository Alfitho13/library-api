package models

import "gorm.io/gorm"

type Buku struct{
	gorm.Model
	Judul_buku string `json:"judul_buku"`
	Penulis string	`json:"penulis"`
	Penerbit string	`json:"penerbit"`
	Tahun_terbit int `json:"tahun_terbit"`
}