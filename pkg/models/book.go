package models

import(
	"github.com/jinzhu/gorm"
	"github.com/ars-shukla23/go-book-mngmt/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Title string `gorm:"" json:"title"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db=config.getDB()
	db.AutoMigrate(&Book{})
}