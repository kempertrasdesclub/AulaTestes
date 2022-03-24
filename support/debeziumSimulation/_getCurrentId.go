package debeziumSimulation

import "errors"

// GetCurrentId
//
// Retorna o ID atual do banco de dados.
//
//   Saída:
//     id: ID atual do banco de dados
func (e *DebeziumSimulation) GetCurrentId() (id interface{}, err error) {
	if e.realDataPointer == nil {
		err = errors.New("use SetData() function first")
		return
	}
	
	if e.create == nil {
		return
	}
	
	return e.id, err
}
