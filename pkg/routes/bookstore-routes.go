package routes

import(
	"github.com/gorilla/mux"
	"github.com/ars-shukla23/go-book-mngmt/pkg/controllers"
)

var RegisterBookStoreRoutes=func(router *mux.Router){
	router.HandleFunc("/book/",controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{BookId}",controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{BookId}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{BookId}",controllers.DeleteBook).Methods("DELETE")
	
}