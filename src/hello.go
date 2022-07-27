package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Paulo"
	idade := 20

	versao := 1.2

	fmt.Println("Olá, Sr.", nome, ", sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O nome da variável versão é", reflect.TypeOf(versao))
}
