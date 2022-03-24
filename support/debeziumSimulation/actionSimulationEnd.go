package debeziumSimulation

import "github.com/helmutkemper/util"

func (e *DebeziumSimulation) actionSimulationEnd() {
	var err error

	err = e.sendOnSimulationEnd()
	if err != nil {
		util.TraceToLog()
		e.ErrChan <- err
		return
	}

	return
}
