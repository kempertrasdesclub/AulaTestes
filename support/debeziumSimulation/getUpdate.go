package debeziumSimulation

import (
	"errors"
)

// GetUpdate (português): Retorna o dado com novos valores após a atualização
//   Saída:
//     id: ID do banco de dados
//     before: Dado antes do evento
//     after: Dado depois do evento
//     err: Objeto de erro padrão
func (e *DebeziumSimulation) GetUpdate() (id, before, after interface{}, err error) {
	if e.realDataPointer == nil {
		err = errors.New("use SetData() function first")
		return
	}

	if e.create == nil {
		err = errors.New("use Populate() function first")
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

	e.update[id.(string)] = FileLineFormat{Id: id.(string), Data: after}

	return
}
