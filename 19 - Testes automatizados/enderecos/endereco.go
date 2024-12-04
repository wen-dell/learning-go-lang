package enderecos

import (
	"strings"
)

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}
	primeiraPalavraEndereco := strings.Split(endereco, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if strings.EqualFold(tipo, primeiraPalavraEndereco) {
			endereco = tipo
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return endereco
	}

	return "Tipo Inválido"
}
