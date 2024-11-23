package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 2
	subtracao := 10 - 2
	multiplicacao := 3 * 5
	divisao := 30 / 3
	restoDivisao := 10 % 5
	fmt.Println(soma)
	fmt.Println(subtracao)
	fmt.Println(multiplicacao)
	fmt.Println(divisao)
	fmt.Println(restoDivisao)

	// ATRIBUICAO
	var numero int8 = 2
	fmt.Println(numero)

	// RELACIONAIS
	fmt.Println(1 > 3)
	fmt.Println(5 < 2)
	fmt.Println(5 >= 2)
	fmt.Println(5 <= 2)
	fmt.Println(5 == 2)
	fmt.Println(5 != 2)

	// LOGICOS
	fmt.Println(true && true)
	fmt.Println(false || false)
	fmt.Println(!false)

	var texto string
	if numero > 1 {
		texto = "Maior do que 1"
	} else {
		texto = "Menor do que 1"
	}

	fmt.Println(texto)

	// OPERADORES UNARIOS
	valor := 10
	valor++
	valor += 15
	valor--
	valor -= 20
	valor *= -1
	valor /= 2
	valor %= 2
	fmt.Println(valor)
}
