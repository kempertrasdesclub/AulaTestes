package debeziumSimulation

import (
	"encoding/json"
	"errors"
	"github.com/helmutkemper/util"
	"log"
	"time"
)

func (e *DebeziumSimulation) sendOnSimulationEnd() (err error) {
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
	e.After = nil
	e.Operation = "z"
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
