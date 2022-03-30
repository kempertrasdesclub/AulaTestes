package messagingSystemNats

import (
	"errors"
	"github.com/helmutkemper/util"
)

// New (Português): fábrica do objeto
//   Entrada:
//     connection: elementos de conexão do nats
//       Elemento 0: string de conexão
//       Elemento N+1: lista de objetos nats.Options separados por vírgula
//   Saída:
//     object: Ponteiro de MessagingSystemNats pronto para uso
//     err: Objeto de erro padrão
//
// Exemplo:
//
//     New("nats://10.0.0.2:4222", nats.MaxReconnects(-1), nats.PingInterval(20*time.Second))
//
//     Nota: - Go não recomenda função compatível com sobrecarga, mas, a construção dessa
//             função facilita a criação de módulos dinâmicos.
func (e *MessagingSystemNats) New(connection ...interface{}) (object interface{}, err error) {
	if connection == nil {
		util.TraceToLog()
		err = errors.New("connection first element must be a string connection")
		return
	}

	switch connection[0].(type) {
	case string:
	default:
		util.TraceToLog()
		err = errors.New("connection first element must be a string connection")
		return
	}

	var options []interface{}
	for k, opt := range connection {
		if k == 0 {
			continue
		}

		options = append(options, opt)
	}

	err = e.Connect(connection[0].(string), options...)
	return e, err
}
