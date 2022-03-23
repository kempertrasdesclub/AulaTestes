package debeziumSimulation

import "github.com/helmutkemper/util"

func (e *DebeziumSimulation) actionUpdateData() {
	var err error

	var after interface{}
	var before interface{}
	var key interface{}
	var line FileLineFormat
	for key, line = range e.update {
		after = line.Data
		break
	}
	delete(e.delete, key)

	before = e.create[key].Data

	err = e.SendOnUpdateData(after, before)
	if err != nil {
		util.TraceToLog()
		e.ErrChan <- err
		return
	}

	return
}
