package debeziumSimulation

func (e *DebeziumSimulation) EnableOnStartData() {
	e.sendOnPopulateData = true
}
