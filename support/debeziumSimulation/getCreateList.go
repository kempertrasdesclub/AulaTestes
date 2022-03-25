package debeziumSimulation

import "errors"

// GetCreateList
//
// Retorna a lista de dados a serem criados
//
//   Entrada:
//     list: map[string]struct{Id string `json:"id"`, Name string `json:"name"`}
func (e *DebeziumSimulation) GetCreateList() (list interface{}, err error) {
	if e.realDataPointer == nil {
		err = errors.New(KErrorUseSetDataFunctionFirst)
		return
	}

	return e.create, err
}
