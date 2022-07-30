package main

import (
	"fmt"
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

func main() {
	exibeIntroducao()
	exibirMenu()

	comandoEscolhido := lerComando()

	switch comandoEscolhido {
	case 1:
		fmt.Println("Iniciar Monitoramento.")
	case 2:
		fmt.Println("Exibir logs.")
	case 0:
		os.Exit(0)
	default:
		os.Exit(-1)
	}
}
