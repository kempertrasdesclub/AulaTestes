package dataSimulation

import "github.com/brianvoe/gofakeit/v6"

// Populate (português): Preenche o dado durante os eventos
func (e *RealDataSimulation) Populate() (err error) {
	e.Id = e.getNextId()
	e.Name = gofakeit.Name()
	return
}
