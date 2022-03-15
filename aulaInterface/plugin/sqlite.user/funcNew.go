package main

import (
	"github.com/helmutkemper/util"
	"test/aulaInterface/constants"
)

func (e *SQLiteUser) New() (referenceInitialized interface{}, err error) {
	err = e.Connect(constants.KSQLiteConnectionString)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = e.Install()
	if err != nil {
		util.TraceToLog()
		return
	}

	return e, nil
}
