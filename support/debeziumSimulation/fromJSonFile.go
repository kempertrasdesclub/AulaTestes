package debeziumSimulation

import (
	"bytes"
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
)

func (e *DebeziumSimulation) FromJSonFile(path string) (err error) {
	var file *os.File

	file, err = os.OpenFile(path, os.O_RDONLY, fs.ModePerm)
	if err != nil {
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("file.Close().err: %v", err)
		}
	}()

	var line = make([]byte, 0)
	var endOfFile bool
	for {
		var char = make([]byte, 1)
		for {
			_, err = file.Read(char)
			if errors.Is(err, io.EOF) == true {
				endOfFile = true
				break

			} else if err != nil {
				return
			}

			if bytes.Equal(char, []byte("\n")) {
				break
			}

			line = append(line, char...)
		}

		if endOfFile == true {
			return
		}

		err = e.processLine(line)
		if err != nil {
			return
		}

		line = make([]byte, 0)
	}
}
