package debeziumSimulation

import (
	"encoding/json"
	"errors"
	"github.com/helmutkemper/util"
	"log"
	"time"
)

// sendOnReadData
//
// Envia um dado criado pelo sistema de mensageria.
//
//   Entrada:
//     after: dado criado;
//
//   Saída:
//     err: Objeto de erro padrão do go.
func (e *DebeziumSimulation) sendOnReadData(after interface{}) (err error) {
	if e.messagingTopicOnStart == "" {
		util.TraceToLog()
		err = errors.New(KErrorMessagingTopicOnStartIsNotSet)
		return
	}

	if e.messagingSystem == nil {
		util.TraceToLog()
		err = errors.New(KErrorMessagingInterfaceIsNotSet)
		return
	}

	var dataToSend []byte

	e.Before = nil
	e.After = after
	e.Operation = KOperationRead
	e.EventDate = time.Now().Unix()

	dataToSend, err = json.Marshal(e)
	if err != nil {
		util.TraceToLog()
		log.Printf("json.Marshal(e).error: %v", err.Error())
		return
	}

	err = e.messagingSystem.Publish(e.messagingTopicOnCreate, dataToSend)
	if err != nil {
		util.TraceToLog()
	}

	return
}
