package debeziumSimulation

import (
	"errors"
)

// getUpdate
//
// Retorna o dado com novos valores após a atualização
//
//   Saída:
//     id: ID do banco de dados;
//     before: Dado antes do evento;
//     after: Dado depois do evento;
//     err: Objeto de erro padrão.
func (e *DebeziumSimulation) getUpdate() (id, before, after interface{}, err error) {
	if e.realDataPointer == nil {
		err = errors.New(KErrorUseSetDataFunctionFirst)
		return
	}

	if e.create == nil {
		err = errors.New(KErrorUseSetDataFunctionFirst)
		return
	}

	id, err = e.realDataPointer.GetID()
	if err != nil {
		return
	}

	before = e.realDataPointer.Get()

	err = e.realDataPointer.Update()
	if err != nil {
		return
	}

	after = e.realDataPointer.Get()

	if e.enableSaveData == true {
		e.update[id] = FileLineFormat{Id: id, Data: after}
	}

	return
}
