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

func adicionarEspacoLinha() {
	fmt.Println("----------------------------")
}

func exibirIntroducao() {
	nome := "Paulo"
	versao := 1.2

	adicionarEspacoLinha()

	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	adicionarEspacoLinha()
}

func exibirMenu() {
	fmt.Println("1 - Iniciar Monitoramento.")
	fmt.Println("2 - Exibir logs.")
	fmt.Println("0 - Sair do programa.")

	adicionarEspacoLinha()
}

func lerComando() int {
	var comandoLido int

	fmt.Scan(&comandoLido)

	return comandoLido
}

func registraLog(site string, status bool) {
	arquivo, erro := os.OpenFile("./src/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if erro != nil {
		fmt.Println(erro)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func monitarSite(site string) {
	res, erro := http.Get(site)

	if erro != nil {
		fmt.Println(erro)

		adicionarEspacoLinha()

		return
	}

	statusSucesso := res.StatusCode >= 200 || res.StatusCode < 300

	if statusSucesso {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code [", res.Status, "]")
	}

	adicionarEspacoLinha()

	registraLog(site, statusSucesso)
}

func lerSitesDoArquivo() []string {
	arquivo, erro := os.Open("./src/sites.txt")

	if erro != nil {
		fmt.Println(erro)

	}

	leitor := bufio.NewReader(arquivo)
	sliceLinhas := []string{}

	for {
		linha, erro := leitor.ReadString('\n')
		linhaSemQuebra := strings.TrimSpace(linha)

		sliceLinhas = append(sliceLinhas, linhaSemQuebra)

		if erro == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sliceLinhas
}

func iniciarMonitoramento() {
	adicionarEspacoLinha()

	fmt.Println("Monitorando...")

	adicionarEspacoLinha()

	sites := lerSitesDoArquivo()

	const vezesDeMonitoramento = 5
	const cincoSegundos = 5 * time.Second

	for i := 1; i <= vezesDeMonitoramento; i++ {
		for index, site := range sites {
			indexMaisUm := index + 1

			fmt.Println("Testando o site", indexMaisUm, ":", site)

			monitarSite(site)
		}

		if i != vezesDeMonitoramento {
			fmt.Println("Esperando... (x", i, ")")
			adicionarEspacoLinha()

			time.Sleep(cincoSegundos)
		}
	}
}

func exibirLogs() {
	adicionarEspacoLinha()

	arquivo, erro := ioutil.ReadFile("./src/log.txt")

	if erro != nil {
		fmt.Println(erro)
	}

	fmt.Println(string(arquivo))

	adicionarEspacoLinha()
}

func main() {
	exibirIntroducao()

	for {
		exibirMenu()

		comandoEscolhido := lerComando()

		switch comandoEscolhido {
		case 1:
			iniciarMonitoramento()
		case 2:
			exibirLogs()
		case 0:
			os.Exit(0)
		default:
			os.Exit(-1)
		}
	}
}
