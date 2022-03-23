package debeziumSimulation

import "github.com/helmutkemper/util"

func (e *DebeziumSimulation) actionCreateData() {
	var err error

	//log.Printf("actionCreateData()")

	var after interface{}
	var line FileLineFormat
	for _, line = range e.create {
		after = line.Data

		err = e.SendOnNewData(after)
		if err != nil {
			util.TraceToLog()
			e.ErrChan <- err
			return
		}
	}

	return
}
