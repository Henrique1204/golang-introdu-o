package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
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

func monitarSite(site string) {
	res, erro := http.Get(site)

	if erro != nil {
		fmt.Println(erro)

		adicionarEspacoLinha()

		return
	}

	if res.StatusCode >= 200 || res.StatusCode < 300 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code [", res.Status, "]")
	}

	adicionarEspacoLinha()
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

func main() {
	exibirIntroducao()

	for {
		exibirMenu()

		comandoEscolhido := lerComando()

		switch comandoEscolhido {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibir logs.")
		case 0:
			os.Exit(0)
		default:
			os.Exit(-1)
		}
	}
}
