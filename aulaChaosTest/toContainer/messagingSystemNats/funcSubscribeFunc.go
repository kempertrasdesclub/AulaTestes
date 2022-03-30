package messagingSystemNats

import (
	"github.com/helmutkemper/util"
	"github.com/nats-io/nats.go"
	"log"
)

// subscribeFunc (português): função compatível com o framework github.com/nats-io/nats.go
// encarregada de chamar as funções de evento do usuário.
//
//   Nota: - Veja a função Subscribe
func (e *MessagingSystemNats) subscribeFunc(msg *nats.Msg) {
	var err error
	var subject = msg.Subject

	e.mutex.Lock()
	defer e.mutex.Unlock()

	for _, funcItem := range e.publishList[subject] {
		err = funcItem(subject, msg.Data)
		if err != nil {
			util.TraceToLog()
			log.Printf("%v.funct().error: %v", subject, err.Error())
		}
	}
}
