package models

import (
	"Golang/M3/banco"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Valor      float64
	Quantidade int32
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

	/*
		defer faz com que ele aguarda toda a execução para posteriormente
		realizar o fechamento do banco
	*/
	defer db.Close()

	return produtos
}
