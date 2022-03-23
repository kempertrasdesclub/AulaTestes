package debeziumSimulation

// EnableOnStartData
//
// Habilita uma descarga de dados iniciais no sistema de mensagem, simulando recebimento de dados
// prévios do banco de dados.
//
//   Entrada:
//     length: quantidade de dados a serem enviados.
func (e *DebeziumSimulation) EnableOnStartData(length int) {
	e.sendOnPopulateData = length
}
