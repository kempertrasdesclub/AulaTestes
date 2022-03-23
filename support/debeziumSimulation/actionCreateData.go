package debeziumSimulation

import "github.com/helmutkemper/util"

func (e *DebeziumSimulation) actionCreateData() {
	var err error

	//log.Printf("actionCreateData()")

	var after interface{}
	_, after, err = e.GetCreate()
	if err != nil {
		util.TraceToLog()
		e.ErrChan <- err
		return
	}

	err = e.SendOnNewData(after)
	if err != nil {
		util.TraceToLog()
		e.ErrChan <- err
		return
	}

	return
}
