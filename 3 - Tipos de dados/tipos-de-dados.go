package main

import "fmt"

func main() {
	var numero int16 = 100
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// int32 = rune
	var numero3 rune = 100
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numero5 float32 = 123.32
	fmt.Println(numero5)

	var numero6 float64 = 1230000.32
	fmt.Println(numero6)

	var str string = "Texto"
	fmt.Println(str)
}
