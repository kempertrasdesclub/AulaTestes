package messagingSystemNats

import (
	"errors"
	"github.com/nats-io/nats.go"
	"time"
	"toContainer/commonTypes"
)

// Connect (Português): Conecta ao servidor de fila Nats
//   Entrada:
//     connection: string de conexão nats
//     options: lista de objetos nats.Option, compatível com o framework
//              github.com/nats-io/nats.go, separados por vírgula.
//   Saída:
//     err: Objeto padrão de erro
//
//   Nota: - Use a New para preparar o objeto e conectar ao servidor de fila Nats
func (e *MessagingSystemNats) Connect(connection interface{}, options ...interface{}) (err error) {

	switch connection.(type) {
	case string:
	default:
		err = errors.New("connection must be a string")
		return
	}

	for _, optionValue := range options {
		switch optionValue.(type) {
		case nats.Option:
		default:
			err = errors.New("options must be a nats.Option object")
			return
		}
	}

	e.errorCounter = 0
	e.connectionString = connection.(string)
	e.publishList = make(map[string][]func(subject string, data []byte) (err error))

	e.options = make([]nats.Option, 0)
	for _, optionValue := range options {
		e.options = append(e.options, optionValue.(nats.Option))
	}

	e.connectToNats()
	e.report()

	e.ticker = time.NewTicker(commonTypes.KQueueTickerPing)
	return
}
