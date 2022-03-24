package debeziumSimulation

import (
	"math/rand"
	"time"
)

// getRandFloat64
//
// Retorna um número aleatório.
//
//   Saída:
//     number: número de ponto flutuante entre 0 e 1 aleatório
func (e *DebeziumSimulation) getRandFloat64() (number float64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	number = r.Float64()
	return
}
