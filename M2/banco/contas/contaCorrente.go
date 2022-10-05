package contas

import (
	"Golang/M2/banco/clientes"
	"fmt"
)

type ContaCorrente struct {
	/*
		Assim como na instanciação de variaveis,
		os valores de cada campo da struct inicia com
		o ZeroValue de cada tipo
	*/
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int     //propriedade publica
	saldo         float64 //propriedade privada
}

func (c *ContaCorrente) Sacar(valor float64) (string, float64) {
	allow := valor <= c.saldo && valor > 0
	if allow {
		c.saldo -= valor
		return "Saque realizado.", c.saldo
	} else {
		return "Valor inválido.", c.saldo
	}

}

func (c *ContaCorrente) Depositar(valor float64) {
	allow := valor > 0
	if allow {
		c.saldo += valor
		fmt.Println("Deposito realizado")
	} else {
		fmt.Println("Valor inválido")
	}
}

/*
Como vamos dentro da função vai fazer a alteração do conteúdo do apontamento da conta destino
é preciso pegar o alocamento da variável
*/
func (c *ContaCorrente) Transferir(contaDestino *ContaCorrente, valor float64) bool {
	allow := valor <= c.saldo && valor > 0
	if allow {
		c.saldo -= valor
		contaDestino.saldo += valor
		return true
	} else {
		return false
	}
}

/*
Com a ideia de encapsular uma propriedade não permitindo que esta seja alterada a qualquer momento
sem regras a serem aplicadas
*/
func (c *ContaCorrente) MostrarSaldo() float64 {
	return c.saldo
}
