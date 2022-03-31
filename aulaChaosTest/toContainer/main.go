package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/memberlist"
	"github.com/helmutkemper/util"
	"io/fs"
	"log"
	"net"
	"os"
	"sync"
	"time"
	"toContainer/messagingSystemNats"
)

const (
	KMemberListUpdateInterval = 5 * time.Second
	KCachePort                = ":11211"
)

type DebeziumSource struct {
	Db    string `json:"db"`
	Table string `json:"table"`
}

type DebeziumData struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

type Debezium struct {
	Source DebeziumSource `json:"source"`
	Before DebeziumData   `json:"before"`
	After  DebeziumData   `json:"after"`
	Op     string         `json:"op"`
}

var list *memberlist.Memberlist
var endOfDataStream = make(chan struct{})
var memory *sync.Map
var fileToSaveData *os.File

// dataFilePath: Environment var contendo o caminho do arquivo com dados a serem gravados na memória
// cache para teste.
//
// O teste consiste em lê um arquivo externo, contendo dados predefinidos, fazer o sincronismo entre
// as instâncias e salvar uma descarga da memória cache, em outro arquivo, para saber se os dados
// batem.
func main() {
	var err error

	memory = new(sync.Map)

	err = os.MkdirAll("/memory_container", fs.ModePerm)
	if err != nil {
		util.TraceToLog()
		log.Printf("bug: os.MkdirAll().error: %v", err.Error())
		return
	}

	list, err = memberlist.Create(memberlist.DefaultLocalConfig())
	if err != nil {
		util.TraceToLog()
		log.Printf("bug: memberlist.Create().error: %v", err.Error())
		return
	}

	// Join an existing cluster by specifying at least one known member.

	var ip = make([]net.IP, 0)
	var infinityLoop = 100
	for {
		ip, err = net.LookupIP("container_delete_after_test_0")
		if err != nil {
			infinityLoop -= 1
			time.Sleep(2 * time.Second)

			if infinityLoop <= 0 {
				log.Printf("bug: Infinity loop break")
				break
			}

			continue
		}
		log.Printf("IP: %v", ip[0].String())
		break
	}
	_, err = list.Join([]string{ip[0].String()})
	if err != nil {
		util.TraceToLog()
		log.Printf("bug: list.Join().error: %v", err.Error())
		return
	}

	// Ask for members of the cluster
	for _, member := range list.Members() {
		fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
	}

	_ = os.Remove("/memory_container/data.file.json")
	fileToSaveData, err = os.OpenFile("/memory_container/data.file.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, fs.ModePerm)
	if err != nil {
		util.TraceToLog()
		log.Printf("bug: os.OpenFile().error: %v", err.Error())
		return
	}

	defer func() {
		err = fileToSaveData.Close()
		if err != nil {
			util.TraceToLog()
			log.Printf("bug: file.Close().error: %v", err.Error())
			return
		}
	}()

	var messageSystem = &messagingSystemNats.MessagingSystemNats{}
	for {
		_, err = messageSystem.New("nats://10.0.0.2:4222")
		if err != nil {
			util.TraceToLog()
			log.Printf("bug: messageSystem.New().error: %v", err.Error())
		}

		err = messageSystem.Subscribe("stocksMessage", natsEventFunc)
		if err != nil {
			util.TraceToLog()
			log.Printf("bug: messageSystem.Subscribe().error: %v", err.Error())
		}

		if err == nil {
			break
		}
	}

	log.Print("chaos enable")
	go func() {
		time.Sleep(20 * time.Second)
		log.Print("you can restart now")
	}()

	<-endOfDataStream
	log.Printf("endOfDataStream")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(20 * time.Second)
	}()

}

type FileLineFormat struct {
	Action string
	Id     interface{}
	Data   DebeziumData
}

func writeToFile(key interface{}, value DebeziumData, action string) (err error) {
	var dataToJSon []byte
	var toFile FileLineFormat
	toFile.Id = key
	toFile.Action = action
	toFile.Data = value

	dataToJSon, err = json.Marshal(&toFile)
	if err != nil {
		util.TraceToLog()
		log.Printf("bug: json.Marshal().error: %v", err.Error())
		return
	}

	_, err = fileToSaveData.Write(dataToJSon)
	if err != nil {
		util.TraceToLog()
		log.Printf("bug: file.Write().error: %v", err.Error())
		return
	}

	_, err = fileToSaveData.WriteString("\n")
	if err != nil {
		util.TraceToLog()
		log.Printf("bug: file.WriteString().error: %v", err.Error())
		return
	}

	return
}

func natsEventFunc(subject string, data []byte) (err error) {
	var debezium Debezium

	if subject != "stocksMessage" {
		err = errors.New("subject topic error")
		return
	}

	err = json.Unmarshal(data, &debezium)
	if err != nil {
		util.TraceToLog()
		if err != nil {
			util.TraceToLog()
			log.Printf("bug: json.Unmarshal().error: %v", err.Error())
			return
		}
	}

	switch debezium.Op {
	case "c":
		log.Printf("create: %s", data)
		memory.Store(debezium.After.Id, data)

		err = writeToFile(debezium.After.Id, debezium.After, debezium.Op)
		if err != nil {
			util.TraceToLog()
			log.Printf("bug: writeToFile().error: %v", err.Error())
			return
		}

	case "r":
		log.Printf("read: %s", data)
		memory.Store(debezium.After.Id, data)

		err = writeToFile(debezium.After.Id, debezium.After, debezium.Op)
		if err != nil {
			util.TraceToLog()
			log.Printf("bug: writeToFile().error: %v", err.Error())
			return
		}

	case "u":
		log.Printf("update: %s", data)
		memory.Store(debezium.After.Id, data)

		err = writeToFile(debezium.After.Id, debezium.After, debezium.Op)
		if err != nil {
			util.TraceToLog()
			log.Printf("bug: writeToFile().error: %v", err.Error())
			return
		}

	case "d":
		log.Printf("delet: %s", data)
		memory.Delete(debezium.Before.Id)

		err = writeToFile(debezium.Before.Id, debezium.Before, debezium.Op)
		if err != nil {
			util.TraceToLog()
			log.Printf("bug: writeToFile().error: %v", err.Error())
			return
		}

	case "z":
		endOfDataStream <- struct{}{}
	}
	return
}
