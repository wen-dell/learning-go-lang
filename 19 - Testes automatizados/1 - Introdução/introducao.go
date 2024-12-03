package main

import (
	"fmt"
	"testes-automatizados/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rua dos Bobos")
	fmt.Println(tipoEndereco)
}
