package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Wendell",
		"sobrenome": "Alves",
	}

	fmt.Println(usuario)
}
