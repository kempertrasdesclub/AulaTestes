package messagingSystemNats

import (
	"github.com/helmutkemper/util"
	"github.com/nats-io/nats.go"
	"log"
)

// connectToNats (português): conecta ao servidor de fila Nats
func (e *MessagingSystemNats) connectToNats() {
	var err error

	e.conn, err = nats.Connect(e.connectionString, e.options...)
	if err != nil {
		util.TraceToLog()
		e.errorCounter += 1
		log.Printf("nats connection error: %v", err.Error())
	} else {
		e.errorCounter = 0
		log.Printf("nats connection ok")
	}
}
