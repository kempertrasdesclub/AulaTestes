package debeziumSimulation

import "github.com/helmutkemper/util"

func (e *DebeziumSimulation) actionUpdateData() {
	var err error

	var after interface{}
	var before interface{}

	_, before, after, err = e.GetUpdate()
	if err != nil {
		util.TraceToLog()
		e.ErrChan <- err
		return
	}

	err = e.SendOnUpdateData(after, before)
	if err != nil {
		util.TraceToLog()
		e.ErrChan <- err
		return
	}

	return
}
