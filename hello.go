package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	for {
		// exibeNomes()
		exibeIntroducao()
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit((-1))
		}
	}

}

// func devolveNomeEIdade() (string, int) {
// 	nome := "Daniel"
// 	idade := 36
// 	return nome, idade
// }

func exibeIntroducao() {
	nome := "Daniel"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão ", versao)
}
func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O endereço da variável comando é ", &comandoLido)
	fmt.Println("Comando digitado: ", comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	// sites := []string{
	// 	"https://random-status-code.herokuapp.com/",
	// 	"https://alura.com.br",
	// 	"https://caelum.com.br"}

	sites := leSitesDoArquivo()
	fmt.Println(sites)

	// for i := 0; i < len(sites); i++ {	// 	fmt.Println(sites[i])// }

	for j := 0; j < monitoramentos; j++ {
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
	if err != nil {
		fmt.Println("Ocorreu um erro ao testar o site: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, " foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site", site, " está com problrmas. Status Code: ", resp.StatusCode)
		registraLog(site, false)
	}
}

func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"}
	nomes = append(nomes, "Aparecida")
	nomes = append(nomes, "Maria")
	nomes = append(nomes, "Joana")
	nomes = append(nomes, "Silvana")
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println(len(nomes))
	fmt.Println(cap(nomes))

}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	// arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		if err == io.EOF {
			break
			// fmt.Println("Ocorreu algum erro na leitura do arquivo: ", err)
		}
		linha = strings.TrimSpace(linha) //remove \n
		sites = append(sites, linha)
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("err")
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Erro ao ler log: ", err)
	}
	fmt.Println(string(arquivo))
}
