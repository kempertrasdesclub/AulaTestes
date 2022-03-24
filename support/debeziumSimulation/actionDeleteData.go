package debeziumSimulation

import "github.com/helmutkemper/util"

// actionDeleteData
//
// Gera um dado apagado e envia pelo sistema de mensagem.
func (e *DebeziumSimulation) actionDeleteData() {
	var err error
	var before interface{}

	_, before, err = e.getDelete()
	if err != nil {
		util.TraceToLog()
		e.ErrChan <- err
		return
	}

	err = e.sendOnDeleteData(before)
	if err != nil {
		util.TraceToLog()
		e.ErrChan <- err
		return
	}

	return
}
