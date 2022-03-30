package messagingSystemNats

import (
	"github.com/helmutkemper/util"
)

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
func (e *MessagingSystemNats) Subscribe(subject string, function func(subject string, data []byte) (err error)) (err error) {
	var found bool
	var list []func(subject string, data []byte) (err error)

	e.mutex.Lock()

	list, found = e.publishList[subject]
	if found == false {
		list = make([]func(subject string, data []byte) (err error), 0)
	}

	list = append(list, function)
	list = append(list, e.Log)
	e.publishList[subject] = list
	e.mutex.Unlock()

	_, err = e.conn.Subscribe(subject, e.subscribeFunc)
	if err != nil {
		util.TraceToLog()
	}

	return
}
