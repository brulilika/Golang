package main

import (
	"Golang/M3/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}
