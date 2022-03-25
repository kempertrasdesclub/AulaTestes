package main

import (
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/hashicorp/memberlist"
	"github.com/helmutkemper/util"
	"log"
	"net"
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
var cacheClient *memcache.Client

// dataFilePath: Environment var contendo o caminho do arquivo com dados a serem gravados na memória
// cache para teste.
//
// O teste consiste em lê um arquivo externo, contendo dados predefinidos, fazer o sincronismo entre
// as instâncias e salvar uma descarga da memória cache, em outro arquivo, para saber se os dados
// batem.
func main() {
	var err error
	var endOfDataStream = make(chan struct{})
	var updateMemberListTicker *time.Ticker

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
		ip, err = net.LookupIP("memcache_container_delete_after_test_0")
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

	cacheClient = &memcache.Client{}
	updateMemberListCache()

	updateMemberListTicker = time.NewTicker(KMemberListUpdateInterval)
	go func() {
		for {
			select {
			case <-updateMemberListTicker.C:
				updateMemberListCache()
			}
		}
	}()

	var messageSystem = messagingSystemNats.MessagingSystemNats{}
	_, err = messageSystem.New("nats://10.0.0.2:4222")
	if err != nil {
		util.TraceToLog()
		log.Printf("bug: messageSystem.New().error: %v", err.Error())
		return
	}

	err = messageSystem.Subscribe(
		"stocksMessage",
		func(subject string, data []byte) (err error) {
			var debezium Debezium
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
				err = cacheClient.Set(
					&memcache.Item{
						Key:   debezium.After.Id,
						Value: data,
					},
				)
				if err != nil {
					util.TraceToLog()
					log.Printf("bug: cacheClient.Set().error: %v", err.Error())
					return
				}
			case "r":
				err = cacheClient.Set(
					&memcache.Item{
						Key:   debezium.After.Id,
						Value: data,
					},
				)
				if err != nil {
					util.TraceToLog()
					log.Printf("bug: cacheClient.Set().error: %v", err.Error())
					return
				}
			case "u":
				err = cacheClient.Set(
					&memcache.Item{
						Key:   debezium.After.Id,
						Value: data,
					},
				)
				if err != nil {
					util.TraceToLog()
					log.Printf("bug: cacheClient.Set().error: %v", err.Error())
					return
				}
			case "d":
				err = cacheClient.Set(
					&memcache.Item{
						Key:   debezium.After.Id,
						Value: data,
					},
				)
				if err != nil {
					util.TraceToLog()
					log.Printf("bug: cacheClient.Set().error: %v", err.Error())
					return
				}
			case "z":
				endOfDataStream <- struct{}{}
			}
			return
		},
	)
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	<-endOfDataStream

}

func updateMemberListCache() {
	var listOfCacheServers = make([]string, 0)
	var members []*memberlist.Node
	members = list.Members()
	for _, nodePointer := range members {
		listOfCacheServers = append(listOfCacheServers, nodePointer.Addr.String()+KCachePort)
	}

	cacheClient = memcache.New(listOfCacheServers...)
}
