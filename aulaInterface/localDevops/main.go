package main

import (
	builder "github.com/helmutkemper/iotmaker.docker.builder"
	"time"
)

func main() {
	var err error
	var mongoDocker = &builder.ContainerBuilder{}
	mongoDocker.SetImageName("mongo:latest")
	mongoDocker.SetContainerName("container_delete_mongo_after_test")
	mongoDocker.AddPortToExpose("27017")
	mongoDocker.SetEnvironmentVar(
		[]string{
			"--host 0.0.0.0",
		},
	)
	//err = mongoDocker.AddFileOrFolderToLinkBetweenConputerHostAndContainer("./test/data", "/data")
	//if err != nil {
	//	panic(err)
	//}

	mongoDocker.SetWaitStringWithTimeout(`"msg":"Waiting for connections","attr":{"port":27017`, 20*time.Second)
	err = mongoDocker.Init()
	if err != nil {
		panic(err)
	}

	err = mongoDocker.ContainerBuildAndStartFromImage()
	if err != nil {
		panic(err)
	}
}
