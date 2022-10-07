package controllers

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

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := models.BuscaProdutos()

	//Lembrar que é necessário "embeddar" o html para que o Golang consiga ler ele
	temp.ExecuteTemplate(w, "Index", produtos)

}
