package debeziumSimulation

import (
	"encoding/json"
	"errors"
	"github.com/helmutkemper/util"
	"time"
)

// sendOnDeleteData
//
// Envia o dado apagado pelo sistema de mensageria.
//
//   Entrada:
//     before: dado antes de ser apagado;
//
//   Saída:
//     err: Objeto de erro padrão do go.
func (e *DebeziumSimulation) sendOnDeleteData(before interface{}) (err error) {
	if e.messagingTopicOnDelete == "" {
		util.TraceToLog()
		err = errors.New("messaging topic on delete is not set")
		return
	}

	if e.messagingSystem == nil {
		util.TraceToLog()
		err = errors.New("messaging interface is not set")
		return
	}

	var dataToSend []byte

	e.Before = before
	e.After = nil
	e.Operation = "d"
	e.EventDate = time.Now().Unix()

	dataToSend, err = json.Marshal(e)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = e.messagingSystem.Publish(e.messagingTopicOnDelete, dataToSend)
	if err != nil {
		util.TraceToLog()
	}

	return
}
