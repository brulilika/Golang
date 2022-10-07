package routes

import (
	"Golang/M4/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.GetAllPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetByIdPersonalidades).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
