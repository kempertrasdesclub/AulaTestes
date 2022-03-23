package debeziumSimulation

import (
	"io/fs"
	"log"
	"os"
)

func (e *DebeziumSimulation) ToJSonFile(path string) (err error) {
	var file *os.File

	file, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, fs.ModePerm)
	if err != nil {
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("file.Close().err: %v", err)
		}
	}()

	err = e.createToFile(file, &e.create, "c")
	if err != nil {
		return
	}

	err = e.createToFile(file, &e.update, "u")
	if err != nil {
		return
	}

	err = e.createToFile(file, &e.delete, "d")
	if err != nil {
		return
	}

	return
}
