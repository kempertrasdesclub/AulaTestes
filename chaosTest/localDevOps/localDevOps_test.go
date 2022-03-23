package localDevOps

import (
	dockerBuilder "github.com/helmutkemper/iotmaker.docker.builder"
	dockerBuilderNetwork "github.com/helmutkemper/iotmaker.docker.builder.network"
	"github.com/helmutkemper/util"
	"log"
	"test/chaosTest/dataTest"
	"test/support/debeziumSimulation"
	"test/support/messagingSystemNats"
	"testing"
	"time"
)

func TestLocalDevOps(t *testing.T) {
	var err error
	var netDocker *dockerBuilderNetwork.ContainerBuilderNetwork

	// Remove os elementos docker do teste anterior
	dockerBuilder.SaGarbageCollector()

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
		var simulation = &dockerBuilder.ContainerBuilder{}
		err = dockerSimulationInstall(netDocker, simulation, i)
		if err != nil {
			util.TraceToLog()
			log.Printf("error: %v", err.Error())
			t.FailNow()
		}
	}

	var dTest = dataTest.DataTest{}
	var messageSystem = messagingSystemNats.MessagingSystemNats{}
	_, err = messageSystem.New("nats://0.0.0.0:4222")
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	err = messageSystem.Subscribe("stocksMessage", func(subject string, data []byte) (err error) {
		log.Printf("nats: %s", data)
		return
	})
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	var debezium = &debeziumSimulation.DebeziumSimulation{}
	debezium.EnableOnStartData(10)
	debezium.SetData(&dTest)
	debezium.SetMessagingSystem(&messageSystem)
	debezium.SetMessagingTopic("stocksMessage")
	debezium.SetTimers(
		5*time.Millisecond,
		50*time.Millisecond,
		70*time.Millisecond,
		100*time.Millisecond,
	)

	err = debezium.Init(false, "tradersclub", "simulation")
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	time.Sleep(10 * time.Minute)
}
