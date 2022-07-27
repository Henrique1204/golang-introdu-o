package main

import "fmt"

func main() {
	const nome string = "Paulo"
	const idade int = 20
	const versao float32 = 1.2

	fmt.Println("Olá, Sr.", nome, ", sua idade é:", idade)
	fmt.Println("Este programa está na versão", versao)
}
