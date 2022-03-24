package debeziumSimulation

import (
	"errors"
)

// getCreate
//
// Retorna o dado preenchido para ser enviado ao Debezium.
//
//   Saída:
//     id: ID do banco de dados;
//     after: Dado após o evento;
//     err: Objeto de erro padrão.
func (e *DebeziumSimulation) getCreate() (id, after interface{}, err error) {
	if e.realDataPointer == nil {
		err = errors.New("use SetData() function first")
		return
	}

	if e.create == nil {
		e.create = make(map[interface{}]FileLineFormat)
		e.update = make(map[interface{}]FileLineFormat)
		e.delete = make(map[interface{}]FileLineFormat)
	}

	err = e.realDataPointer.Populate()
	if err != nil {
		return
	}

	id, err = e.realDataPointer.GetID()
	if err != nil {
		return
	}

	after = e.realDataPointer.Get()

	if e.enableSaveData == true {
		e.create[id] = FileLineFormat{Id: id, Data: after}
	}

	return
}
