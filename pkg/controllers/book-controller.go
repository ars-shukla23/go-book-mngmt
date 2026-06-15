package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/ars-shukla23/go-book-mngmt/pkg/models"
	"github.com/ars-shukla23/go-book-mngmt/pkg/utils"

)

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request){
	newBooks:=models.GetAllBooks()
	res,_:=json.Marshal(newBooks)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBookByID(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	bookId:=vars["BookId"]
	Id,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("Error while parsing")

	}
	bookDetails,_:=models.GetBookByID(Id)
	res,_:=json.Marshal(bookDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	book:=&models.Book{}
	utils.ParseBody(r,book)
	b:=book.CreateBook()
	res,_:=json.Marshal(b)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	bookId:=vars["BookId"]
	Id,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("Error while parsing")
	}
	book:=models.DeleteBook(Id)
	res,_:=json.Marshal(book)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)


}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook=&models.Book{}
	utils.ParseBody(r,updateBook)
	vars:=mux.Vars(r)
	bookId:=vars["BookId"]
	Id,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("Error while parsing")
	}
	bookdetails,db:=models.GetBookByID(Id)
	if updateBook.Title!=""{
		bookdetails.Title=updateBook.Title	
	}
	if updateBook.Author!=""{
		bookdetails.Author=updateBook.Author
	}
	if updateBook.Publication!=""{
		bookdetails.Publication=updateBook.Publication
	}
	db.Save(&bookdetails)
	res,_:=json.Marshal(bookdetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}