package main

import (
	"github.com/helmutkemper/util"
)

// Install (Português): Instala o menu e popula os primeiros dados.
func (e *SQLiteUser) Install() (err error) {
	var installed = false

	installed, err = e.verifyInstallUser()
	if err != nil {
		util.TraceToLog()
		return
	}

	if installed == false {
		err = e.createTableUser()
		if err != nil {
			util.TraceToLog()
			return
		}

		err = e.populateInitialUser()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	return
}
