package main

import (
	"fmt"
	"net/http"
	"os"
)

func exibeIntroducao() {
	nome := "Paulo"
	versao := 1.2

	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibirMenu() {
	fmt.Println("1 - Iniciar Monitoramento.")
	fmt.Println("2 - Exibir logs.")
	fmt.Println("0 - Sair do programa.")
}

func lerComando() int {
	var comandoLido int

	fmt.Scan(&comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	siteUrl := "https://acessibilidade-senai.vercel.app"

	res, _ := http.Get(siteUrl)

	if res.StatusCode >= 200 || res.StatusCode < 300 {
		fmt.Println("Site:", siteUrl, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", siteUrl, "está com problemas. Status code [", res.Status, "]")
	}
}

func main() {
	exibeIntroducao()

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
