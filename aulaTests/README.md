# Tests

Esta pasta mostra uma forma simples de rodar os testes em formato linear e em formato paralelo de 
forma simples.

Testes em paralelo são uma excelente forma de encontrar segmentação de memória e travamentos.

### Falha de concepção

Este é um teste sem paralelismo rodando um código escrito para travar quando rodado em paralelo.

Ele funcionará normalmente, sem travar, mas, o erro está na forma como o teste foi concebido.

```go
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
```

### Forma correta

O teste abaixo é praticamente o mesmo teste, porém, rodando com paralelismo habilitado, possibilitando 
a detecção do problema.

```go
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
```

### Forma correta de usar map em go

O código abaixo mostra a forma correta de usar mapa em Golang e a forma correta de testar, com 
paralelismo.

Forma correta de usar map em Golang:
```go
package aulaTests

import "sync"

type CorrectMap struct {
	m    sync.Mutex
	data map[interface{}]interface{}
}

func (e *CorrectMap) Store(key, value interface{}) {
	e.m.Lock()
	defer e.m.Unlock()
	
	if e.data == nil {
		e.data = make(map[interface{}]interface{})
	}
	
	e.data[key] = value
}

func (e *CorrectMap) Load(key interface{}) (value interface{}) {
	e.m.Lock()
	defer e.m.Unlock()
	
	if e.data == nil {
		e.data = make(map[interface{}]interface{})
	}
	
	return e.data[key]
}
```

Forma correta de escrever o teste em Golang:

```go
package aulaTests

import (
	"github.com/brianvoe/gofakeit/v6"
	"testing"
)

// TestNormalMapParallelsCorrect usa o tipo CorrectMap{} com a forma correta de usar um mapa em
// paralelo.
//
// Este código funciona corretamente e o teste também está correto, com paralelismo.
func TestNormalMapParallelsCorrect(t *testing.T) {
	
	// Mapa feito para não violar a memória e permitir paralelismo.
	var memory = CorrectMap{}
	
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
```