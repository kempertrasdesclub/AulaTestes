package main

import (
	"fmt"
	builder "github.com/helmutkemper/iotmaker.docker.builder"
	"github.com/helmutkemper/util"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/signal"
	"syscall"
	"test/aulaInterface/businessRules/system/datasource"
	"test/aulaInterface/gin/server"
	"time"
)

func main() {
	var err error

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		log.Printf("Removendo elementos docker")
		builder.SaGarbageCollector()
		var resposta string
		log.Printf("Remover a imagem mongo:latest? [y/n]")
		_, _ = fmt.Scan(&resposta)

		if resposta == "y" || resposta == "Y" {
			builder.SaGarbageCollector("mongo:latest")
		}

		os.Exit(0)
	}()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	builder.SaGarbageCollector()
	err = installEphemeralMongodb()
	if err != nil {
		util.TraceToLog()
		log.Panic(err)
	}

	err = datasource.Linker.Init(datasource.KMongoDB)
	if err != nil {
		panic(err)
	}

	err = server.ConfigAndStart()
	if err != nil {
		util.TraceToLog()
		log.Panic(err)
	}
}

func installEphemeralMongodb() (err error) {
	var mongoDocker = &builder.ContainerBuilder{}
	mongoDocker.SetImageName("mongo:latest")
	mongoDocker.SetContainerName("container_delete_mongo_after_test")
	mongoDocker.AddPortToExpose("27017")
	mongoDocker.SetEnvironmentVar(
		[]string{
			"--host 0.0.0.0",
		},
	)

	mongoDocker.SetWaitStringWithTimeout(`"msg":"Waiting for connections","attr":{"port":27017`, 20*time.Second)
	err = mongoDocker.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	err = mongoDocker.ContainerBuildAndStartFromImage()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
