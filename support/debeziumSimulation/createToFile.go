package debeziumSimulation

import (
	"encoding/json"
	"os"
)

// createToFile
//
// Converte um bloco de dados salvos em json e salva no arquivo.
//
//   Entrada:
//     file: Ponteiro do arquivo a ser salvo;
//     list: Ponteiro para o bloco de dados a ser salvo;
//     action: c - data created; u - data updated; d - data deleted.
//   Saída:
//     err: Objeto de erro padrão do go
func (e *DebeziumSimulation) createToFile(file *os.File, list *map[interface{}]FileLineFormat, action string) (err error) {
	var dataToJSon []byte
	for _, data := range *list {
		var toFile FileLineFormat
		toFile.Id = data.Id
		toFile.Action = action
		toFile.Data = data.Data

		dataToJSon, err = json.Marshal(&toFile)
		if err != nil {
			return
		}

		_, err = file.Write(dataToJSon)
		if err != nil {
			return
		}

		_, err = file.WriteString("\n")
		if err != nil {
			return
		}
	}

	return
}
