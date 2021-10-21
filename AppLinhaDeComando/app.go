package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	monitoramentos = 4
	delay          = 3
	colocacao      = "º"
)

func exibeIntro() {

	var nome string
	fmt.Print("Digite seu nome: ")
	fmt.Scan(&nome)
	versao := 1.1
	fmt.Println("****************************************")
	fmt.Println("*Bem-vindo!", nome)
	fmt.Println("*Estamos na versão do programa", versao)
}

func carregaComando() int {
	var comandoCarregado int
	fmt.Print("Digite seu comando: ")
	fmt.Scan(&comandoCarregado)
	fmt.Println("Comando Digitado ", comandoCarregado, "\n ")
	return comandoCarregado
}

func exibeOpcoes() {
	fmt.Println("****************************************")
	fmt.Println("Digite a opção que você deseja executar.")
	fmt.Println("1 - Monitorar Sites")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	fmt.Println("****************************************\n ")
}

func startMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://golang.org/", "https://ge.globo.com", "https://alura.com.br", "https://udemy.com"}

	for j := 0; j <= monitoramentos; j++ {
		for i, site := range sites {
			fmt.Println("Testando looping", j+1, "site", i+1, ":", site)
			testaSite(site)
		}
	}
}

func testaSite(s string) {
	resp, _ := http.Get(s)
	if resp.StatusCode == 200 {
		fmt.Println("Site", s, "OK! Status Code:", resp.StatusCode, "\n ")
		time.Sleep(time.Second * delay)
	} else {
		fmt.Println("Site", s, "Fora! Status Code:", resp.StatusCode, "\n ")
		time.Sleep(time.Second * delay)
	}
}

func main() {
	exibeIntro()

	for {
		exibeOpcoes()

		comando := carregaComando()
		switch comando {
		case 1:
			startMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa.")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido.")
			os.Exit(-1)
		}
	}

}
