package debeziumSimulation

import "errors"

// GetDeleteList (português): Retorna a lista de dados a serem apagados
//   Entrada:
//     list: map[string]struct{Id string `json:"id"`, Name string `json:"name"`}
func (e *DebeziumSimulation) GetDeleteList() (list interface{}, err error) {
	if e.realDataPointer == nil {
		err = errors.New("use SetData() function first")
		return
	}

	return e.delete, err
}
