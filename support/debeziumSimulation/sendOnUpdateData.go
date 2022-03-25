package debeziumSimulation

import (
	"encoding/json"
	"errors"
	"github.com/helmutkemper/util"
	"time"
)

// sendOnUpdateData
//
// Envia o dado atualizado pelo sistema de mensageria.
//
//   Entrada:
//     after: dado após ser atualizado;
//     before: dado antes de ser atualizado;
//
//   Saída:
//     err: Objeto de erro padrão do go.
func (e *DebeziumSimulation) sendOnUpdateData(after, before interface{}) (err error) {
	if e.messagingTopicOnUpdate == "" {
		util.TraceToLog()
		err = errors.New(KErrorMessagingTopicOnUpdateIsNotSet)
		return
	}

	if e.messagingSystem == nil {
		util.TraceToLog()
		err = errors.New(KErrorMessagingTopicOnStartIsNotSet)
		return
	}

	var dataToSend []byte

	e.Before = before
	e.After = after
	e.Operation = KOperationUpdate
	e.EventDate = time.Now().Unix()

	dataToSend, err = json.Marshal(e)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = e.messagingSystem.Publish(e.messagingTopicOnUpdate, dataToSend)
	if err != nil {
		util.TraceToLog()
	}

	return
}
