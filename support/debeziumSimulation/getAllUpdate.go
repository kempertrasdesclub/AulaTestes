package debeziumSimulation

// GetAllUpdate
//
// Retorna todos os dados atualizados;
//
//   Saída:
//     data: Lista de dados;
//     err: Objeto de erro padrão do go.
func (e *DebeziumSimulation) GetAllUpdate() (data map[interface{}]FileLineFormat, err error) {
	data = e.update
	return
}
