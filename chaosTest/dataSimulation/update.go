package dataSimulation

import "github.com/brianvoe/gofakeit/v6"

// Update (português): Altera algum dado para simular o usuário
func (e *RealDataSimulation) Update() (err error) {
	e.Name = gofakeit.Name()
	return
}
