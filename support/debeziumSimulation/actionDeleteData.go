package debeziumSimulation

import "github.com/helmutkemper/util"

func (e *DebeziumSimulation) actionDeleteData() {
	var err error
	var before interface{}

	_, before, err = e.GetDelete()
	if err != nil {
		util.TraceToLog()
		e.ErrChan <- err
		return
	}

	err = e.SendOnDeleteData(before)
	if err != nil {
		util.TraceToLog()
		e.ErrChan <- err
		return
	}

	return
}
