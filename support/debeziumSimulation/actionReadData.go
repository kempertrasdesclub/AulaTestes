package debeziumSimulation

import "github.com/helmutkemper/util"

// actionCreateData
//
// Cria um único novo dado e envia pelo sistema de mensagem.
func (e *DebeziumSimulation) actionReadData() {
	var err error

	var after interface{}
	_, after, err = e.getCreate()
	if err != nil {
		util.TraceToLog()
		e.ErrChan <- err
		return
	}

	err = e.sendOnReadData(after)
	if err != nil {
		util.TraceToLog()
		e.ErrChan <- err
		return
	}

	return
}
