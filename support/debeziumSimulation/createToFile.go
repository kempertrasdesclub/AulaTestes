package debeziumSimulation

import (
	"encoding/json"
	"os"
)

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
