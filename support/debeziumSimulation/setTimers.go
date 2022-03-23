package debeziumSimulation

import "time"

// SetTimers (português): Define os intervalos de tempo para cada ação sofrida pelo dado
//   Entrada:
//     start: Tempo de atraso para início dos eventos.
//     create: Intervalo entre os eventos de criação de dados.
//     update: Intervalos entre os eventos de atualização.
//     delete: Intervalo entre os eventos onde o dado é apagado.
func (e *DebeziumSimulation) SetTimers(start, create, update, delete time.Duration) {
	e.sendOnStartDelay = start
	e.sendOnCreateDelay = create
	e.sendOnUpdateDelay = update
	e.sendOnDeleteDelay = delete
}
