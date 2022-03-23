package debeziumSimulation

func (e *DebeziumSimulation) GetAllDelete() (data map[interface{}]FileLineFormat, err error) {
	data = e.delete
	return
}
