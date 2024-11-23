package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	variavel1++

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int

	fmt.Println(variavel3, ponteiro)

	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro) // desreferenciação
}
