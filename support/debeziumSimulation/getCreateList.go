package debeziumSimulation

import "errors"

// GetCreateList (português): Retorna a lista de dados a serem criados
//   Entrada:
//     list: map[string]struct{Id string `json:"id"`, Name string `json:"name"`}
func (e *DebeziumSimulation) GetCreateList() (list interface{}, err error) {
	if e.realDataPointer == nil {
		err = errors.New("use SetData() function first")
		return
	}

	return e.create, err
}
