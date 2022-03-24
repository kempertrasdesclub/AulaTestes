package debeziumSimulation

import "github.com/helmutkemper/util"

// actionUpdateData
//
// Cria um dado atualizado e envia pelo sistema de mensagem.
func (e *DebeziumSimulation) actionUpdateData() {
	var err error

	var after interface{}
	var before interface{}

	_, before, after, err = e.getUpdate()
	if err != nil {
		util.TraceToLog()
		e.ErrChan <- err
		return
	}

	err = e.sendOnUpdateData(after, before)
	if err != nil {
		util.TraceToLog()
		e.ErrChan <- err
		return
	}

	return
}
