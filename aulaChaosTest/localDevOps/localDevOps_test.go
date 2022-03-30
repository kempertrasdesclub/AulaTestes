package localDevOps

import (
	"bytes"
	"encoding/json"
	"errors"
	dockerBuilder "github.com/helmutkemper/iotmaker.docker.builder"
	dockerBuilderNetwork "github.com/helmutkemper/iotmaker.docker.builder.network"
	"github.com/helmutkemper/util"
	"github.com/kempertrasdesclub/AulaTestes/aulaChaosTest/dataTest"
	"github.com/kempertrasdesclub/AulaTestes/support/debeziumSimulation"
	"github.com/kempertrasdesclub/AulaTestes/support/messagingSystemNats"
	"io/fs"
	"io/ioutil"
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

	var enableChaos = os.Getenv("CHAOS_TEST") == "1"
	var showLog = os.Getenv("LOG") == "1"

	t.Cleanup(
		func() {
			dockerBuilder.SaGarbageCollector()
		},
	)

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

	var simulation = make([]*dockerBuilder.ContainerBuilder, 2)
	for i := int64(0); i != 2; i += 1 {
		simulation[i] = &dockerBuilder.ContainerBuilder{}

		var suffix = strconv.FormatInt(i, 10)

		var memoryPath = "../localDevOps/memory/container_" + suffix
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

		err = dockerSimulationInstall(netDocker, simulation[i], i, memoryPath, enableChaos)
		if err != nil {
			util.TraceToLog()
			log.Printf("error: %v", err.Error())
			t.FailNow()
		}
	}

	for i := int64(0); i != 2; i += 1 {
		err = simulation[i].ContainerStartAfterBuild()
		if err != nil {
			util.TraceToLog()
			log.Printf("Error: %v", err.Error())
			return
		}

		// English: Starts container monitoring at two second intervals. This functionality monitors the container's standard output and generates the log defined by the SetCsvLogPath() function.
		//
		// Português: Inicializa o monitoramento do container com intervalos de dois segundos. Esta funcionalidade monitora a saída padrão do container e gera o log definido pela função SetCsvLogPath().
		// StartMonitor() é usado durante o teste de caos e na geração do log de desempenho do container.
		// [optional/opcional]
		simulation[i].StartMonitor()
	}

	var dataSimulation = dataTest.DataTest{}
	var messageSystem = messagingSystemNats.MessagingSystemNats{}
	_, err = messageSystem.New("nats://0.0.0.0:4222")
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	err = messageSystem.Subscribe("stocksMessage", func(subject string, data []byte) (err error) {
		if showLog == true {
			log.Printf("nats: %s", data)
		}

		return
	})
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	var debezium = &debeziumSimulation.DebeziumSimulation{}
	debezium.SetData(&dataSimulation)
	debezium.SetMessagingSystem(&messageSystem)
	debezium.SetMessagingTopic("stocksMessage")
	debezium.SetTimers(
		50*time.Millisecond,
		500*time.Millisecond,
		700*time.Millisecond,
		1000*time.Millisecond,
		2*60*time.Second,
	)

	err = debezium.Init(true, "tradersclub", "simulation")
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	ch := debezium.GetTerminationChannel()
	<-ch

	for i := int64(0); i != 2; i += 1 {
		// English: For container monitoring. Note: This function should be used to avoid trying to read a container that no longer exists, erased by the SaGarbageCollector() function.
		//
		// Português: Para o monitoramento do container. Nota: Esta função deve ser usada para evitar tentativa de leitura em um container que não existe mais, apagado pela função SaGarbageCollector().
		// [optional/opcional]
		_ = simulation[i].StopMonitor()
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(10 * time.Second)
		wg.Done()
	}()
	wg.Wait()

	for i := int64(0); i != 2; i += 1 {
		var suffix = strconv.FormatInt(i, 10)
		var memoryPath = "../localDevOps/memory/container_" + suffix

		for {
			_, err = os.Stat(memoryPath + "/data.file.json")
			if err == nil {
				break
			}

			time.Sleep(1 * time.Second)
		}

		var dataFileJson []byte
		var dataFileJsonLine [][]byte
		var dataFileJSonParsed = make(map[interface{}]debeziumSimulation.FileLineFormat)
		dataFileJson, err = ioutil.ReadFile(memoryPath + "/data.file.json")
		if err != nil {
			util.TraceToLog()
			log.Printf("error: %v", err.Error())
			t.FailNow()
		}

		dataFileJsonLine = bytes.Split(dataFileJson, []byte("\n"))
		for _, data := range dataFileJsonLine {
			if len(data) == 0 {
				continue
			}

			var line debeziumSimulation.FileLineFormat
			err = json.Unmarshal(data, &line)
			if err != nil {
				util.TraceToLog()
				log.Printf("error: %v", err.Error())
				t.FailNow()
			}

			dataFileJSonParsed[line.Id] = line
		}

		var debeziumData map[interface{}]debeziumSimulation.FileLineFormat
		debeziumData, err = debezium.GetAllCreate()
		if err != nil {
			util.TraceToLog()
			log.Printf("error: %v", err.Error())
			t.FailNow()
		}

		for k := range debeziumData {
			var found bool
			_, found = dataFileJSonParsed[k]
			if found == false {
				log.Printf("Id create fail: %v", debeziumData[k].Id)
				t.FailNow()
			}
		}

		debeziumData, err = debezium.GetAllUpdate()
		if err != nil {
			util.TraceToLog()
			log.Printf("error: %v", err.Error())
			t.FailNow()
		}

		for k := range debeziumData {
			var found bool
			_, found = dataFileJSonParsed[k]
			if found == false {
				log.Printf("Id update fail: %v", debeziumData[k].Id)
				t.FailNow()
			}
		}

		debeziumData, err = debezium.GetAllDelete()
		if err != nil {
			util.TraceToLog()
			log.Printf("error: %v", err.Error())
			t.FailNow()
		}

		for k := range debeziumData {
			var found bool
			_, found = dataFileJSonParsed[k]
			if found == false {
				log.Printf("Id delete fail: %v", debeziumData[k].Id)
				t.FailNow()
			}
		}

	}

	log.Print("fim!")
}

func mergeChannels(cs ...<-chan dockerBuilder.Event) <-chan dockerBuilder.Event {
	out := make(chan dockerBuilder.Event)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan dockerBuilder.Event) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
