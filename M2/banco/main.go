package main

import (
	"Golang/M2/banco/clientes"
	"Golang/M2/banco/contas"
	"fmt"
)

// Null em Go é nil
type iConta interface {
	Sacar(valor float64) (string, float64)
}

/*
A necessidade de implementação de determinada função a partir de uma interface
é dada de modo implicito a partir da utilização da função por parte da struct
*/
func pagarBoleto(conta iConta, valor float64) {
	conta.Sacar(valor)
}

func main() {
	minhaConta := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Bruna",
			CPF:       "1234",
			Profissao: "Programadora",
		},
		NumeroAgencia: 123,
		NumeroConta:   321}
	minhaConta.Depositar(500.0)
	minhaConta2 := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Bruna",
		CPF:       "1234",
		Profissao: "Programadora",
	},
		NumeroAgencia: 123,
		NumeroConta:   321}
	minhaConta2.Depositar(250)
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
	fmt.Println("Novo Saldo:", minhaConta.MostrarSaldo())

	if minhaConta.Transferir(&minhaConta2, 10.0) {
		fmt.Println("Novo Saldo Conta 1:", minhaConta.MostrarSaldo())
		fmt.Println("Novo Saldo Conta 2:", minhaConta2.MostrarSaldo())
	} else {
		fmt.Println("Não foi possível fazer a transferência")
	}

	pagarBoleto(&minhaConta, 100)
	fmt.Println("Novo Saldo Conta 1:", minhaConta.MostrarSaldo())
}
