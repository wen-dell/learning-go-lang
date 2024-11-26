package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("Dentro do método salvar")
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()
	fmt.Println(usuario1.maiorDeIdade())
	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}
