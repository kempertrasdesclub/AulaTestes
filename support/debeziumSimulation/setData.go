package debeziumSimulation

import (
	"github.com/kempertrasdesclub/AulaTestes/support/interfaces"
)

// SetData
//
// Recebe o ponteiro para o dado a ser usado na simulação.
//
//   Entrada:
//     realData: objeto compatível com a interfaces.DataToSimulateInterface.
func (e *DebeziumSimulation) SetData(realData interfaces.DataToSimulateInterface) {
	e.realDataPointer = realData
}
