package messagingSystemNats

import "github.com/helmutkemper/util"

// Publish (português): Publica uma informação no canal específico da fila
//   Entrada:
//     subject: Nome do canal
//     data: Array de bytes contendo a informação a ser compartilhada
//   Saída:
//     err: Objeto de erro padrão
func (e *MessagingSystemNats) Publish(subject string, data []byte) (err error) {
	err = e.conn.Publish(subject, data)
	if err != nil {
		util.TraceToLog()
	}
	return
}
