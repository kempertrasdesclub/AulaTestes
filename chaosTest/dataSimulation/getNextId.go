package dataSimulation

import "github.com/brianvoe/gofakeit/v6"

// getNextId: Retorna o próximo Id como em um banco de dados
func (e *RealDataSimulation) getNextId() (id string) {
	return gofakeit.UUID()
}
