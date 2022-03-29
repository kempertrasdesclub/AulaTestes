package aulaTests

import (
	"github.com/brianvoe/gofakeit/v6"
	"testing"
)

// TestNormalMapParallelsError este teste é o mesmo teste de TestNormalMapNoParallels(), porém,
// escrito corretamente e rodando em paralelo.
func TestNormalMapParallelsError(t *testing.T) {

	// mapa não pode ter escrita e leitura em paralelo. Isto trava o código.
	var memory = make(map[interface{}]interface{})

	for i := 0; i != 1000; i += 1 {
		t.Run(
			"escrita",
			func(t *testing.T) {

				// Com o flag de paralelismo, Run() roda todos os testes contendo o flag t.Parallel() em
				// paralelo, e isto possibilita encontrar erros de violação de memória causados pelo
				// paralelismo.
				//
				// Este flag deve está contido em todas as funções feitas para rodar em paralelo.
				t.Parallel()
				id := gofakeit.UUID()
				memory[id] = id
			},
		)
		t.Run(
			"leitura",
			func(t *testing.T) {

				// Com o flag de paralelismo, Run() roda todos os testes contendo o flag t.Parallel() em
				// paralelo, e isto possibilita encontrar erros de violação de memória causados pelo
				// paralelismo.
				//
				// Este flag deve está contido em todas as funções feitas para rodar em paralelo.
				t.Parallel()
				id := gofakeit.UUID()
				_ = memory[id]
			},
		)
	}
}
