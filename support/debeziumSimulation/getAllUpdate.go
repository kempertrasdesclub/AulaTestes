package debeziumSimulation

func (e *DebeziumSimulation) GetAllUpdate() (data map[interface{}]FileLineFormat, err error) {
	data = e.update
	return
}
