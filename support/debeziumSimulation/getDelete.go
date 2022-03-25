package debeziumSimulation

import "errors"

// getDelete
//
// Retorna o dado antes da atualização.
//
//   Saída:
//     id: ID do dado no banco de dados.
//     before: Dado antes do evento.
//     err: Objeto de erro padrão;
func (e *DebeziumSimulation) getDelete() (id, before interface{}, err error) {
	if e.realDataPointer == nil {
		err = errors.New(KErrorUseSetDataFunctionFirst)
		return
	}

	if e.create == nil {
		return
	}

	id, err = e.realDataPointer.GetID()
	if err != nil {
		return
	}

	before = e.realDataPointer.Get()

	if e.enableSaveData == true {
		e.delete[id] = FileLineFormat{Id: id, Data: before}
	}

	return
}
