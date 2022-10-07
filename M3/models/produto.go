package models

import (
	"Golang/M3/banco"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Valor      float64
	Quantidade int
}

func BuscaProdutos() []Produto {
	db := banco.DatabaseConnection()

	dbProducts, err := db.Query("SELECT * FROM PRODUTOS")

	if err != nil {
		panic(err.Error())
	}

	prod := Produto{}
	produtos := []Produto{}

	for dbProducts.Next() {
		var id int
		var quantidade int
		var nome, descricao string
		var valor float64

		err = dbProducts.Scan(&id, &nome, &descricao, &valor, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		prod.Id = id
		prod.Nome = nome
		prod.Descricao = descricao
		prod.Valor = valor
		prod.Quantidade = quantidade

		produtos = append(produtos, prod)
	}

	/*
		defer faz com que ele aguarda toda a execução para posteriormente
		realizar o fechamento do banco
	*/
	defer db.Close()

	return produtos
}

func CriaNovoProduto(nome, descricao string, valor float64, quantidade int) {
	db := banco.DatabaseConnection()

	insert, err := db.Prepare("INSERT INTO PRODUTOS(nome, descricao, valor, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(nome, descricao, valor, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := banco.DatabaseConnection()

	delete, err := db.Prepare("DELETE FROM PRODUTOS WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
	defer db.Close()

}
