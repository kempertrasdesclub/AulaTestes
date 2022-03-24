package debeziumSimulation

import (
	"encoding/json"
)

// processLine
//
// Processa a uma única linha do arquivo de dados e preenche a memória.
//
//   Entrada:
//     line: Linha do arquivo de dados na forma de byte.
//
//   Saída:
//     err: Objeto padrão de erro do go.
func (e *DebeziumSimulation) processLine(line []byte) (err error) {
	if e.create == nil {
		e.create = make(map[interface{}]FileLineFormat)
		e.update = make(map[interface{}]FileLineFormat)
		e.delete = make(map[interface{}]FileLineFormat)
	}

	var fromFile FileLineFormat
	err = json.Unmarshal(line, &fromFile)
	if err != nil {
		return
	}

	switch fromFile.Action {
	case "c":
		e.create[fromFile.Id] = FileLineFormat{Id: fromFile.Id, Data: fromFile.Data}
	case "u":
		e.update[fromFile.Id] = FileLineFormat{Id: fromFile.Id, Data: fromFile.Data}
	case "d":
		e.delete[fromFile.Id] = FileLineFormat{Id: fromFile.Id, Data: fromFile.Data}
	}

	return
}
