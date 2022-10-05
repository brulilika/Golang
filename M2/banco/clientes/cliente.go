package clientes

/*
Em Golang não se tem herança, o que se tem é composição
*/

type Titular struct {
	Nome      string
	CPF       string
	Profissao string
}
