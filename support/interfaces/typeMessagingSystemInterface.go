package interfaces

import "test/support/commonTypes"

type MessagingSystemInterface interface {

	// Subscribe (Português): Adiciona uma função para ser invocada quando um novo evento
	// ocorre no canal.
	//   Entrada:
	//     subject: nome do canal
	//     function: Ponteiro de função a ser executada quando houver um evento no canal
	//   Saída:
	//     err: Objeto de erro padrão
	//
	//     Nota: - O segundo parâmetro, de nome function, é o ponteiro de uma função invocada
	//             pelo código quando ocorre um evento no canal, de nome contido em subject,
	//             e ela vai receber o nome do canal e o dado transmitido.
	//             Entrada:
	//               subject: nome do canal
	//               data: array de byte com a informação transmitida no canal (geralmente
	//               json)
	//             Saída:
	//               err: Objeto padrão de erro
	Subscribe(subject string, function func(subject string, data []byte) (err error)) (err error)

	// SetReport (Português): Invoca uma função periodicamente para informar o status da
	// conexão.
	SetReport(function func(status commonTypes.QueueStatus))

	// Publish (português): Publica uma informação no canal específico da fila
	//   Entrada:
	//     subject: Nome do canal
	//     data: Array de bytes contendo a informação a ser compartilhada
	//   Saída:
	//     err: Objeto de erro padrão
	Publish(subject string, data []byte) (err error)
}
