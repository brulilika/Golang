package routes

import (
	"Golang/M3/controllers"
	"net/http"
)

func CarregaRotas() {
	//Ao encontrar uma "/" executará a função index
	http.HandleFunc("/", controllers.Index)
}
