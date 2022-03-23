package debeziumSimulation

func (e *DebeziumSimulation) GetAllCreate() (data map[interface{}]FileLineFormat, err error) {
	data = e.create
	return
}
