package debeziumSimulation

import (
	"errors"
	"reflect"
)

// CompareJSonFile
//
// Compara um arquivo salvo com a função ToJSonFile() com o conteúdo da memória de modo a permitir a
// comparação entre os testes, principalmente os testes de caos.
//
//   Entrada:
//     path: caminho do arquivo a ser lido;
//   Saída:
//     err: objeto de erro padrão do go.
func (e *DebeziumSimulation) CompareJSonFile(path string) (err error) {

	var loadedData = DebeziumSimulation{}
	err = loadedData.FromJSonFile(path)
	if err != nil {
		return
	}

	var loadedCreate map[interface{}]FileLineFormat
	loadedCreate, err = loadedData.GetAllCreate()

	if len(e.create) != len(loadedCreate) {
		err = errors.New("the amount of data created does not match")
		return
	}

	for k := range e.create {
		if reflect.DeepEqual(loadedCreate[k].Id, e.create[k].Id) == false {
			err = errors.New("data id does no match")
			return
		}
	}

	var loadedUpdate map[interface{}]FileLineFormat
	loadedUpdate, err = loadedData.GetAllUpdate()

	if len(e.update) != len(loadedUpdate) {
		err = errors.New("the amount of data updated does not match")
		return
	}

	for k := range e.update {
		if reflect.DeepEqual(loadedUpdate[k].Id, e.update[k].Id) == false {
			err = errors.New("data id does no match")
			return
		}
	}

	var loadedDelete map[interface{}]FileLineFormat
	loadedDelete, err = loadedData.GetAllDelete()

	if len(e.delete) != len(loadedDelete) {
		err = errors.New("the amount of data deleted does not match")
		return
	}

	for k := range e.delete {
		if reflect.DeepEqual(loadedDelete[k].Id, e.delete[k].Id) == false {
			err = errors.New("data id does no match")
			return
		}
	}

	return
}
