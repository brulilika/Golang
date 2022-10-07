package main

import (
	"Golang/M3/models"
	"database/sql"
	"html/template"
	"net/http"

	/*
		O _ antes da importação serve para mantê-la e permitir sua
		utilização durante o tempo de execução
	*/
	_ "github.com/lib/pq"
)

/*
Forma de buscar todos os arquivos que renderizam tela sem ter de
informar a rota de cada um e criar uma variável para cada uma
*/
var temp = template.Must(template.ParseGlob("templates/*.html"))

func databaseConnection() *sql.DB {
	connection := "user=postgres dbname=loja password=tamake28 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	//Ao encontrar uma "/" executará a função index
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := databaseConnection()

	dbProducts, err := db.Query("SELECT * FROM PRODUTOS")

	if err != nil {
		panic(err.Error())
	}

	prod := models.Produto{}
	produtos := []models.Produto{}

	for dbProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var valor float64

		err = dbProducts.Scan(&id, &nome, &descricao, &valor, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		prod.Nome = nome
		prod.Descricao = descricao
		prod.Valor = valor
		prod.Quantidade = int32(quantidade)

		produtos = append(produtos, prod)
	}
	//Lembrar que é necessário "embeddar" o html para que o Golang consiga ler ele
	temp.ExecuteTemplate(w, "Index", produtos)

	/*
		defer faz com que ele aguarda toda a execução para posteriormente
		realizar o fechamento do banco
	*/
	defer db.Close()
}
