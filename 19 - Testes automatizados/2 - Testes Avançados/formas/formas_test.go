package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		retangulo := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := retangulo.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da esperada %f\n", areaRecebida, areaEsperada)
		}
	})
	t.Run("Circulo", func(t *testing.T) {
		circulo := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circulo.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da esperada %f\n", areaRecebida, areaEsperada)
		}
	})
}
