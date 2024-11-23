package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario
	u.nome = "Wendell"
	u.idade = 27
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	usuario2 := usuario{"Wendell", 27, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Rafael"}
	fmt.Println(usuario3)
}
