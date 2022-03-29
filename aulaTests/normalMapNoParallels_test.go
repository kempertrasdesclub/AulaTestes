package aulaTests

import (
	"github.com/brianvoe/gofakeit/v6"
	"testing"
)

// TestNormalMapNoParallels é um teste aparentemente normal, mas, não está rodando em paralelo,
// fazendo o teste passar, porém o mesmo código falhará em produção.
func TestNormalMapNoParallels(t *testing.T) {

	// mapa não pode ter escrita e leitura em paralelo. Isto trava o código.
	var memory = make(map[interface{}]interface{})

	for i := 0; i != 1000; i += 1 {
		// Sem o flag de paralelismo, Run() roda apenas um teste de cada vez, dando uma falsa sensação de
		// segurança.
		t.Run(
			"escrita",
			func(t *testing.T) {
				id := gofakeit.UUID()
				memory[id] = id
			},
		)
		t.Run(
			"leitura",
			func(t *testing.T) {
				id := gofakeit.UUID()
				_ = memory[id]
			},
		)
	}
}
