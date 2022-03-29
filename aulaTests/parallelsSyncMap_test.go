package aulaTests

import (
	"github.com/brianvoe/gofakeit/v6"
	"sync"
	"testing"
)

func TestParallelsSyncMap(t *testing.T) {
	// Esta é implementação de map[interface{}]interface{} pronta para uso e segura para ser usada em
	// threads.
	var memory = new(sync.Map)

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
				memory.Store(id, id)
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
				memory.Load(id)
			},
		)
	}
}
