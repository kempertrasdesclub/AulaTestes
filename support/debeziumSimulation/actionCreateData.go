package debeziumSimulation

import "github.com/helmutkemper/util"

// actionCreateData
//
// Cria um único novo dado e envia pelo sistema de mensagem.
func (e *DebeziumSimulation) actionCreateData() {
	var err error

	var after interface{}
	_, after, err = e.getCreate()
	if err != nil {
		util.TraceToLog()
		e.ErrChan <- err
		return
	}

	err = e.sendOnNewData(after)
	if err != nil {
		util.TraceToLog()
		e.ErrChan <- err
		return
	}

	return
}
