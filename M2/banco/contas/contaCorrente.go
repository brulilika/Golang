package contas

import "fmt"

type ContaCorrente struct {
	/*
		Assim como na instanciação de variaveis,
		os valores de cada campo da struct inicia com
		o ZeroValue de cada tipo
	*/
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) (string, float64) {
	allow := valor <= c.Saldo && valor > 0
	if allow {
		c.Saldo -= valor
		return "Saque realizado.", c.Saldo
	} else {
		return "Valor inválido.", c.Saldo
	}

}

func (c *ContaCorrente) Depositar(valor float64) {
	allow := valor > 0
	if allow {
		c.Saldo += valor
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
	allow := valor <= c.Saldo && valor > 0
	if allow {
		c.Saldo -= valor
		contaDestino.Saldo += valor
		return true
	} else {
		return false
	}

}
