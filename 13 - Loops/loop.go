package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Printf("i está valendo %d\n", i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j")
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, valor := range nomes {
		fmt.Printf("Indice %d\n", indice)
		fmt.Printf("Valor %s\n", valor)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Printf("Indice {%d} Código ASCII {%d}\n", indice, letra)
		fmt.Printf("Indice {%d} Letra {%s}\n", indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Wendell",
		"sobrenome": "Alves",
	}

	for chave, valor := range usuario {
		fmt.Printf("Chave {%s}\n", chave)
		fmt.Printf("Valor {%s}\n", valor)
	}
}
