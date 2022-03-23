package debeziumSimulation

import "github.com/helmutkemper/util"

func (e *DebeziumSimulation) actionDeleteData() {
	var err error

	var before interface{}
	var key interface{}
	var line FileLineFormat
	for key, line = range e.delete {
		before = line.Data
		break
	}
	delete(e.delete, key)

	err = e.SendOnDeleteData(before)
	if err != nil {
		util.TraceToLog()
		e.ErrChan <- err
		return
	}

	return
}
