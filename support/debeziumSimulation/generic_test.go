package debeziumSimulation

import (
	"github.com/helmutkemper/util"
	"github.com/kempertrasdesclub/AulaTestes/chaosTest/dataTest"
	"github.com/kempertrasdesclub/AulaTestes/support/messagingSystemNats"
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestDebeziumSimulation_SetMessagingSystem(t *testing.T) {
	var err error
	var dTest = dataTest.DataTest{}
	var messageSystem = messagingSystemNats.MessagingSystemNats{}
	_, err = messageSystem.New("nats://localhost:4222")
	if err != nil {
		util.TraceToLog()
		t.FailNow()
	}

	var debezium = &DebeziumSimulation{}
	debezium.SetData(&dTest)
	debezium.SetMessagingSystem(&messageSystem)
	debezium.SetMessagingTopicOnStart("stocksMessage")
	debezium.SetMessagingTopicOnCreate("stocksMessage")
	debezium.SetMessagingTopicOnUpdate("stocksMessage")
	debezium.SetMessagingTopicOnDelete("stocksMessage")
	debezium.SetTimers(
		0,
		50*time.Millisecond,
		1*time.Millisecond,
		130*time.Millisecond,
		300*time.Millisecond,
	)

	err = debezium.Init(true, "db", "table")
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}
}

func TestFile(t *testing.T) {
	var err error
	var dTest = dataTest.DataTest{}

	var debeziumWriter = &DebeziumSimulation{}
	debeziumWriter.SetData(&dTest)

	err = debeziumWriter.ToJSonFile("./data.test.json.txt")
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	var debeziumReader = &DebeziumSimulation{}
	debeziumReader.SetData(&dTest)
	err = debeziumReader.FromJSonFile("./data.test.json.txt")
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	var dataWriter = make(map[interface{}]FileLineFormat)
	dataWriter, err = debeziumWriter.GetAllCreate()
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	var dataReader = make(map[interface{}]FileLineFormat)
	dataReader, err = debeziumReader.GetAllCreate()
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	if len(dataWriter) != len(dataReader) {
		t.FailNow()
	}

	for k := range dataWriter {
		if reflect.DeepEqual(dataReader[k].Id, dataWriter[k].Id) == false {
			util.TraceToLog()
			log.Printf("error: %v", err.Error())
			t.FailNow()
		}
	}

	dataWriter, err = debeziumWriter.GetAllUpdate()
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	dataReader, err = debeziumReader.GetAllUpdate()
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	if len(dataWriter) != len(dataReader) {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	for k := range dataWriter {
		if reflect.DeepEqual(dataReader[k].Id, dataWriter[k].Id) == false {
			util.TraceToLog()
			log.Printf("error: %v", err.Error())
			t.FailNow()
		}
	}

	dataWriter, err = debeziumWriter.GetAllDelete()
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	dataReader, err = debeziumReader.GetAllDelete()
	if err != nil {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	if len(dataWriter) != len(dataReader) {
		util.TraceToLog()
		log.Printf("error: %v", err.Error())
		t.FailNow()
	}

	for k := range dataWriter {
		if reflect.DeepEqual(dataReader[k].Id, dataWriter[k].Id) == false {
			util.TraceToLog()
			log.Printf("error: %v", err.Error())
			t.FailNow()
		}
	}

	_ = os.Remove("./data.test.json.txt")
}
