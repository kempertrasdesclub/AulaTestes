package debeziumSimulation

import "test/support/interfaces"

func (e *DebeziumSimulation) SetData(realData interfaces.DataToSimulateInterface) {
	e.realDataPointer = realData
}
