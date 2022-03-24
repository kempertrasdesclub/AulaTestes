package debeziumSimulation

// GetTerminationChannel
//
// Retorna o channel usado para indicar fim da simulação.
//
//   Saída:
//     terminationChannel: canal struct{} indicando fim da simulação
func (e *DebeziumSimulation) GetTerminationChannel() (terminationChannel chan struct{}) {
	return e.TerminationChan
}
