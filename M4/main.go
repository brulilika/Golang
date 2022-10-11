package main

import (
	"Golang/M4/database"
	"Golang/M4/routes"
	"fmt"
)

func main() {
	database.DatabaseConnection()
	fmt.Println("Iniciando...")
	routes.HandleResquest()
}
