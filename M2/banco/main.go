package main

import (
	"Golang/M2/banco/contas"
	"fmt"
)

//Null em Go é nil

func main() {
	minhaConta := contas.ContaCorrente{Titular: "Bruna", NumeroAgencia: 123, NumeroConta: 321, Saldo: 100.00}
	minhaConta2 := contas.ContaCorrente{Titular: "Bruna", NumeroAgencia: 123, NumeroConta: 321, Saldo: 100.00}
	/*
		A utilização do * faz com que os ponteiros
		identifiquem as tipagens criadas e aloquem as
		memórias devidamentes na memória
	*/
	//var segundaConta *ContaCorrente
	//segundaConta = new(ContaCorrente)
	fmt.Println(minhaConta)
	//fmt.Println(*segundaConta)
	/*
		Golang ao fazer a comparação entre objetos, considera a
		comparação entre cada um dos campos.
		Mas se fizermos a alocação por ponteiros, ele não olha o
		conteúdo e sim o endereço, sendo preciso usar a notação *
		para realizar a comparação de conteúdo
		segundaConta == segundaConta2 => false
		*segundaConta == *segundaConta2 => true
	*/
	fmt.Println(minhaConta == minhaConta2)

	resp, valor := minhaConta.Sacar(5.50)
	fmt.Println(resp, "Novo Saldo:", valor)

	minhaConta.Depositar(5.0)
	fmt.Println("Novo Saldo:", minhaConta.Saldo)

	if minhaConta.Transferir(&minhaConta2, 10.0) {
		fmt.Println("Novo Saldo Conta 1:", minhaConta.Saldo)
		fmt.Println("Novo Saldo Conta 2:", minhaConta2.Saldo)
	} else {
		fmt.Println("Não foi possível fazer a transferência")
	}

}
