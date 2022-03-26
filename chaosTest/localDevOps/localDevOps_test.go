package localDevOps

import (
	"errors"
	dockerBuilder "github.com/helmutkemper/iotmaker.docker.builder"
	dockerBuilderNetwork "github.com/helmutkemper/iotmaker.docker.builder.network"
	"github.com/helmutkemper/util"
	"github.com/kempertrasdesclub/AulaTestes/chaosTest/dataTest"
	"github.com/kempertrasdesclub/AulaTestes/support/debeziumSimulation"
	"github.com/kempertrasdesclub/AulaTestes/support/messagingSystemNats"
	"io/fs"
	"log"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestLocalDevOps(t *testing.T) {
	var err error
	var netDocker *dockerBuilderNetwork.ContainerBuilderNetwork

	// Remove os elementos docker do teste anterior
	dockerBuilder.SaGarbageCollector()

	// Remove o arquivo de dump de memória.
	_ = os.Remove("data.file.json")

	// Remove os elementos docker ao final do teste
	//defer dockerBuilder.SaGarbageCollector()

	err = dockerBuildImageCacheBase()
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	err = dockerBuildImageCache()
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	// Cria uma rede docker
	// Como o gateway é 10.0.0.1, o primeiro endereço será 10.0.0.2
	netDocker, err = dockerTestNetworkCreate()
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	err = dockerNatsDownloadAndInstall(netDocker)
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	for i := int64(0); i != 2; i += 1 {
		var suffix = strconv.FormatInt(i, 10)

		var memoryPath = "./localDevOps/memory/container_" + suffix
		_ = os.MkdirAll(memoryPath, fs.ModePerm)

		var fileInfo fs.FileInfo
		fileInfo, err = os.Stat(memoryPath)
		if err != nil {
			util.TraceToLog()
			log.Printf("error: %v", err.Error())
			t.FailNow()
		}

		if fileInfo.IsDir() == false {
			err = errors.New("directory to archive memory data not created")
			util.TraceToLog()
			log.Printf("error: %v", err.Error())
			t.FailNow()
		}

		var simulation = &dockerBuilder.ContainerBuilder{}
		err = dockerSimulationInstall(netDocker, simulation, i, memoryPath)
		if err != nil {
			util.TraceToLog()
			log.Printf("error: %v", err.Error())
			t.FailNow()
		}
	}

	var dataSimulation = dataTest.DataTest{}
	var messageSystem = messagingSystemNats.MessagingSystemNats{}
	_, err = messageSystem.New("nats://10.0.0.2:4222")
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	err = messageSystem.Subscribe("stocksMessage", func(subject string, data []byte) (err error) {
		log.Printf("nats: %s", data)
		return
	})
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	var debezium = &debeziumSimulation.DebeziumSimulation{}
	debezium.EnableOnStartData(1)
	debezium.SetData(&dataSimulation)
	debezium.SetMessagingSystem(&messageSystem)
	debezium.SetMessagingTopic("stocksMessage")
	debezium.SetTimers(
		50*time.Millisecond,
		500*time.Millisecond,
		700*time.Millisecond,
		1000*time.Millisecond,
		5*time.Second,
	)

	err = debezium.Init(true, "tradersclub", "simulation")
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	ch := debezium.GetTerminationChannel()
	<-ch

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(10 * time.Second)
		wg.Done()
	}()
	wg.Wait()

	for i := int64(0); i != 2; i += 1 {
		var suffix = strconv.FormatInt(i, 10)
		var memoryPath = "./localDevOps/memory/container_" + suffix

		for {
			_, err = os.Stat(memoryPath + "/data.file.json")
			if err == nil {
				break
			}

			time.Sleep(1 * time.Second)
		}

		err = debezium.CompareJSonFile(memoryPath + "/data.file.json")
		if err != nil {
			util.TraceToLog()
			log.Printf("error: %v", err.Error())
			t.FailNow()
		}
	}

	log.Print("fim!")
}
