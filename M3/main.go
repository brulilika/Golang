package main

import (
	"Golang/M3/models"
	"html/template"
	"net/http"
)

/*
Forma de buscar todos os arquivos que renderizam tela sem ter de
informar a rota de cada um e criar uma variável para cada uma
*/
var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	//Ao encontrar uma "/" executará a função index
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	prods := []models.Produto{
		{Nome: "Camiseta", Descricao: "Estampa florida azul", Valor: 39.90, Quantidade: 10},
		{Nome: "Saia", Descricao: "Drapeada preta", Valor: 49.90, Quantidade: 5},
	}
	//Lembrar que é necessário "embeddar" o html para que o Golang consiga ler ele
	temp.ExecuteTemplate(w, "Index", prods)
}
