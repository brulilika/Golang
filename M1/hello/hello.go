package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Uso de "MagicNumbers"
const monitoramentos = 2
const delay = 5

func main() {
	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	/*
		Em Golang, ao declarar uma variavel vc pode fazer de duas formas:
			var nome string
		desta forma, caso não atribuido um valor, esta é instanciada vazia
		ou da forma abaixo
	*/
	nome := "Bruna"
	versao := 1.0
	fmt.Println("Olá,", nome)
	fmt.Println("Programa na versão:", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	/*
		Quando se instancia um Array em Golang é preciso pre-determinar seu tamanho e realizar as atribuições
		Para se ter Arrays sem tamanho pre-determinado,
		Utiliza-se "Slice", que é tipo um Array
	*/
	//Instanciação da variavel tipo Slice com atribuição a partir do retorno da função
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		//Uso do range para fazer tipo um forEach
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	//Verificação de existencia ou não de erro
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println("Site:", site, " ", resp.StatusCode)

	if resp.StatusCode == 200 {
		registraLog(site, true)
	} else {
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		//Caracter que vai ser utilizado para determinar o fim de uma linha
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)
		//Erro de END OF FILE
		if err == io.EOF {
			break
		}

	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {

	//Abertura de arquivo com parâmetros para leitura
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}
