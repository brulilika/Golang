package main

import "fmt"

//Null em Go é nil

type ContaCorrente struct {
	/*
		Assim como na instanciação de variaveis,
		os valores de cada campo da struct inicia com
		o ZeroValue de cada tipo
	*/
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
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

func main() {
	minhaConta := ContaCorrente{"Bruna", 123, 321, 100.00}
	minhaConta2 := ContaCorrente{"Bruna", 123, 321, 100.00}
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
	fmt.Println(resp, "Novo saldo:", valor)

	minhaConta.Depositar(5.0)
	fmt.Println("Novo saldo:", minhaConta.saldo)

	if minhaConta.Transferir(&minhaConta2, 10.0) {
		fmt.Println("Novo saldo Conta 1:", minhaConta.saldo)
		fmt.Println("Novo saldo Conta 2:", minhaConta2.saldo)
	} else {
		fmt.Println("Não foi possível fazer a transferência")
	}

}
