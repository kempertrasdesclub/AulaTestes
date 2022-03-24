package debeziumSimulation

// GetAllDelete
//
// Retorna todos os dados apagados;
//
//   Saída:
//     data: Lista de dados;
//     err: Objeto de erro padrão do go.
func (e *DebeziumSimulation) GetAllDelete() (data map[interface{}]FileLineFormat, err error) {
	data = e.delete
	return
}
