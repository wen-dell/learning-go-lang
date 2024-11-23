package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"Davi", "Silva", 29, 165}
	e1 := estudante{p1, "biomedicina", "UFRN"}
	fmt.Println(p1)
	fmt.Println(e1)
}
