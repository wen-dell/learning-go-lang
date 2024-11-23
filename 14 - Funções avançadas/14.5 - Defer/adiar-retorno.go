package main

import "fmt"

func isAlunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Entrando na função para verificar se aluno está aprovado")
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isAlunoAprovado(7, 8))
}
