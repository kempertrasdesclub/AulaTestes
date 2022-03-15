package main

import (
	"github.com/helmutkemper/util"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"test/aulaInterface/businessRules/system/datasource"
	"test/aulaInterface/gin/server"
)

func main() {
	var err error

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err = datasource.Linker.Init(datasource.KSQLite)
	if err != nil {
		panic(err)
	}

	err = server.ConfigAndStart()
	if err != nil {
		util.TraceToLog()
		log.Panic(err)
	}
}
