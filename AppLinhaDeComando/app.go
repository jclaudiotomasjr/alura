package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	monitoramentos = 2
	delay          = 2
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
	//sites := []string{"https://golang.org/", "https://ge.globo.com", "https://alura.com.br", "https://udemy.com"}

	sites := lerSitesDoArquivo()

	for j := 0; j <= monitoramentos; j++ {
		for i, site := range sites {
			fmt.Println("Testando looping", j+1, "site", i+1, ":", site)
			testaSite(site)
		}
	}
}

//função que ler todos os sites que contém no arquivo sites.txt e retorna para função testaSite
func lerSitesDoArquivo() []string {
	var sites []string
	//arquivo, err := ioutil.ReadFile("sites.txt")
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

//função usada para testar os sites que contém no arquivo sites.txt
func testaSite(s string) {
	resp, err := http.Get(s)
	if err != nil {
		fmt.Println("ocorreu um erro", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site", s, "OK! Status Code:", resp.StatusCode, "\n ")
		registraLog(s, true)
		time.Sleep(time.Second * delay)
	} else {
		fmt.Println("Site", s, "Fora! Status Code:", resp.StatusCode, "\n ")
		registraLog(s, false)
		time.Sleep(time.Second * delay)
	}
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func main() {
	exibeIntro()
	registraLog("https://jvis.com.br", false)
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
