package debeziumSimulation

import (
	"encoding/json"
	"errors"
	"github.com/helmutkemper/util"
	"time"
)

func (e *DebeziumSimulation) SendOnPopulateData(after interface{}) (err error) {
	if e.messagingTopicOnStart == "" {
		util.TraceToLog()
		err = errors.New("messaging topic on start is not set")
		return
	}

	if e.messagingSystem == nil {
		util.TraceToLog()
		err = errors.New("messaging interface is not set")
		return
	}

	var dataToSend []byte

	e.Before = nil
	e.After = after
	e.Operation = "r"
	e.EventDate = time.Now().Unix()

	dataToSend, err = json.Marshal(e)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = e.messagingSystem.Publish(e.messagingTopicOnStart, dataToSend)
	if err != nil {
		util.TraceToLog()
	}

	return
}
