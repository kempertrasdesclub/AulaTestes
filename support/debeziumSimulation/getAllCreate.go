package debeziumSimulation

// GetAllCreate
//
// Retorna todos os dados criados;
//
//   Saída:
//     data: Lista de dados;
//     err: Objeto de erro padrão do go.
func (e *DebeziumSimulation) GetAllCreate() (data map[interface{}]FileLineFormat, err error) {
	data = e.create
	return
}
