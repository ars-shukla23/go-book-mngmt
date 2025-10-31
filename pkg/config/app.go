package config

//The whole point of this file will be to return a variable DB which will help the other files interact with the database


import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect(){
	d, err:=gorm.Open("mysql","")
	if err!=nil{
		panic(err)
	}
	db=d
}
func GetDB() *gorm.DB{
	return db
}

