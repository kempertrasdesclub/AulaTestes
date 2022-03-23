package debeziumSimulation

import (
	"math/rand"
	"time"
)

func (e *DebeziumSimulation) getRandFloat64() (number float64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	number = r.Float64()
	return
}
