# Aula Benchmark

Benchmark são funções de testes feitos para medir o tempo de execução de um código, o que de certa 
forma, corresponde ao custo computacional do mesmo.

Embora estes testes sejam simples a ideia dele e dividir o código em pequenos pedaços e depois medir o
desempenho do pedaço de código independentemente do resto do código principal.

```go
package aulaBenchmark

import (
  "testing"
)

func BenchmarkNomeDoSeuTeste(b *testing.B) {
  
  // Habilita o relatório do consumo de memória.
  b.ReportAllocs()
  
  // Define a última coisa a ser feita, ao final dos testes.
  b.Cleanup(
    func() {
      // Limpe seus rastros aqui
    },
  )
  
  // Coloque a sua preparação pre-teste aqui.
  
  // Reinicia o tempo medido e o consumo de memória. Use depois das suas preparações pre-teste.
  b.ResetTimer()
  
  // Seus testes devem ficar dentro desse laço. Ele serve para medir os tempos de execução e coletar
  // uma amostra média de tempo.
  for i := 0; i < b.N; i++ {
    
    // Caso necessite, as funções b.StartTimer() e b.StopTimer() podem te ajudar a controlar melhor o
    // tempo de execução
    
    b.RunParallel(
      func(pb *testing.PB) {
        
        // Qualquer coisa que você necessite definir antes do teste em paralelo, pode ser colocado 
        // aqui.
        
        for pb.Next() {
          
          // O código a ser testado com paralelismo, fica qui.
          
        }
      },
    )
    
  }
}
```

O arquivo `tcserver_test.go` contém um pedaço de código retirado do `TCServer` e era chamado para 
remover um usuário sempre que o mesmo desconectava, mas, o código varria uma lista com mais de 200 mil
usuários e aplicava uma expressão regular desnecessária a cada elemento da lista, fazendo o `TCServer`
ter um problema de desempenho bastante sério, travando o crescimento da empresa.

Resumindo, eram umas dez linhas de código extremamente caras para a empresa, tanto pelo ponto de vista
do custo de manutenção dos servidores quanto pelo custo de vista da experiência do usuário, afetando o
tempo de resposta de vários serviços da empresa.

Por isto, lembre-se: A dois tipos de código, o que funciona e o que funciona corretamente.

Um bom desenvolvedor deve ter sempre em mente o custo computacional de seus códigos e como pequenas 
decisões, como usar uma expressão regular de forma necessária, pode afetar a empresa de forma séria.

Veja um exemplo real abaixo.

```go
package aulaBenchmark

// Este código mostra um pedaço de código real removido do TCServer.
// Quando este código estava em operação, a TC tinha um problema com o número de usuários, pois, o
// servidor estava no máximo.
//
// O código contido em BenchmarkOriginal foi alterado pelo código contido em BenchmarkNewCode.
//
// Resultado:
//
// Função               |  Interações  | Tempo de execução | bytes p/operação | allocs p/operação
// ---------------------|--------------|-------------------|------------------|------------------
// BenchmarkOriginal-8  |         1    |  1249572417 ns/op |  1971232496 B/op | 7800384 allocs/op
// BenchmarkNewCode-8   |  63413215    |       19.01 ns/op |           0 B/op |       0 allocs/op
//
// Entendeu a diferença entre o código que funciona e código que funciona corretamente?
//
// Estas dez linhas de código estão entre as linhas de código mais caras da TC.
//
// Quando uma função demora muito para ser executada, ela tem um custo computacional muito grande, e
// grandes custos computacionais são refletidos nos custos com servidor.
// Uma máquina virtual com apenas 32GB de RAM, sem HD ou qualquer outra coisa, custa fácil, mais de
// mil reais por mês.
//
// Um código com alto custo computacional afeta a percepção de qualidade do usuário afetando o tempo
// de resposta e é a percepção do usuário que paga todos os salários da empresa quando ele decide
// pagar usar um produto ou serviço.
//
// Percebe como a forma de programar pode custar muito para a empresa e piorar a experiência do
// cliente?

import (
  "github.com/brianvoe/gofakeit/v6"
  "regexp"
  "sync"
  "testing"
)

type Status struct {
  UserId         string `json:"user_id"`
  Status         string `json:"status"`
  Manual         bool   `json:"manual"`
  LastActivityAt int64  `json:"last_activity_at"`
  ActiveChannel  string `json:"-" db:"-"`
}

var id string
var key string
var std *Status

var status map[string]*Status

// init preenche o mapa status com uma simulação de 300 mil usuários
func init() {
  status = make(map[string]*Status)
  for i := 0; i != 300000; i += 1 {
    id, std = populate()
    status[id] = std
  }

  // determina key como sendo o último id criado
  key = id
}

// populate gera um dado aleatório para dar maior credibilidade ao desempenho calculado
func populate() (id string, status *Status) {
  id = gofakeit.UUID()
  status = &Status{
    UserId:         id,
    Status:         gofakeit.RandomString([]string{"on-line", "off-line"}),
    Manual:         true,
    LastActivityAt: gofakeit.Date().UnixNano(),
  }

  return
}

// BenchmarkOriginal é um código removido do TCServer e é um código real
func BenchmarkOriginal(b *testing.B) {
  var syncStatus = sync.RWMutex{}

  b.ReportAllocs()
  for i := 0; i < b.N; i++ {
    syncStatus.Lock()
    for i := range status {
      if key == "all" {
        delete(status, i)
      } else {
        matches, _ := regexp.MatchString(key, i)
        if matches {
          delete(status, i)
        }
      }
    }
    syncStatus.Unlock()
  }
}

// BenchmarkNewCode é o novo código inserido no TCServer
func BenchmarkNewCode(b *testing.B) {
  var syncStatus = sync.RWMutex{}

  b.ReportAllocs()
  for i := 0; i < b.N; i++ {
    syncStatus.Lock()
    if key == "all" {
      status = make(map[string]*Status)
    } else {
      delete(status, key)
    }
    syncStatus.Unlock()
  }
}
```