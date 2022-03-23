package debeziumSimulation

import "errors"

// GetUpdateList (português): Retorna a lista de dados a serem atualizados
//   Saída:
//     list: map[string]struct{Id string `json:"id"`, Name string `json:"name"`}
func (e *DebeziumSimulation) GetUpdateList() (list interface{}, err error) {
	if e.realDataPointer == nil {
		err = errors.New("use SetData() function first")
		return
	}

	return e.update, err
}
