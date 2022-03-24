package debeziumSimulation

import "time"

// SetTimers (português): Define os intervalos de tempo para cada ação sofrida pelo dado.
//
//   Entrada:
//     startDelay: Tempo de atraso para início dos eventos;
//     create: Intervalo entre os eventos de criação de dados;
//     update: Intervalos entre os eventos de atualização;
//     delete: Intervalo entre os eventos onde o dado é apagado;
//     termination: Intervalo para terminar o processo de teste.
func (e *DebeziumSimulation) SetTimers(startDelay, create, update, delete, termination time.Duration) {
	e.sendOnStartDelay = startDelay
	e.sendOnCreateDelay = create
	e.sendOnUpdateDelay = update
	e.sendOnDeleteDelay = delete
	e.sendTestProcessTerminationDelay = termination
}
