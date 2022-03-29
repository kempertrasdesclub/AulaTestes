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
