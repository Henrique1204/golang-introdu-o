package main

import "fmt"

func main() {
	nome := "Paulo"
	versao := 1.2

	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar Monitoramento.")
	fmt.Println("2 - Exibir logs.")
	fmt.Println("0 - Sair do programa.")

	var comando int

	fmt.Scan(&comando)

	switch comando {
	case 1:
		fmt.Println("Iniciar Monitoramento.")
	case 2:
		fmt.Println("Exibir logs.")
	case 0:
		fmt.Println("Sair do programa.")
	default:
		fmt.Println("Não conheço o comando", comando)
	}
}
