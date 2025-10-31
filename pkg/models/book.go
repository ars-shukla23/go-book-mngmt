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
	db=config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b) //Checks if the record is new(ID==0)
	db.Create(&b) 
	return b //We are returning Book object as GORM will add the ID to the object
}

func GetAllBooks() []Book{
   var Books []Book
   db.Find(&Books)
   return Books
}

func GetBookbyID(id int64) (*Book, *gorm.DB){
	var getBook Book
	db:=db.Where("ID=?",id).Find(&getBook)
	return &getBook,db
}
