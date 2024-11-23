package main

import "fmt"

func main() {
	var variavel1 string = "variavel1"
	variavel2 := "variavel2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "variavel3"
		variavel4 string = "variavel4"
	)

	const constante1 = "Constante"

	fmt.Println(variavel3, variavel4)
	fmt.Println(constante1)
}
