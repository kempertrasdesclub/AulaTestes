# Debezium CDC Simulation

Este código tem a finalidade de simular o CDC com Debezium.

Os passos para usar este módulo são:

Gere um `struct{}` com todos os dados que você gostaria de simular e construa as funções `Populate()`,
`Update()`, `Get()` e `GetID()` de modo a serem compatíveis com a interface abaixo.

```go
package interfaces

type DataToSimulateInterface interface {
	Populate() (err error)
	Update() (err error)
	Get() (data interface{})
	GetID() (ID interface{}, err error)
}
```

Veja um exemplo simples de dado a ser simulado.

```go
package dataTest

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
)

type DataTest struct {
	Id   string
	Name string
}

func (e *DataTest) GetID() (ID interface{}, err error) {
	ID = e.Id
	return
}

func (e *DataTest) getNextId() (id string) {
	return gofakeit.UUID()
}

func (e *DataTest) Get() (data interface{}) {
	return *e
}

func (e *DataTest) Populate() (err error) {
	e.Id = e.getNextId()
	e.Name = gofakeit.Name()
	return
}

func (e *DataTest) Update() (err error) {
	e.Name = gofakeit.Name()
	return
}
```

Inicialize o sistema de mensageria.

```go
package localDevOps

import (
	"log"
	"testing"
	"time"
)

func TestLocalDevOps(t *testing.T) {
	var err error
	
	var messageSystem = messagingSystemNats.MessagingSystemNats{}
	_, err = messageSystem.New("nats://0.0.0.0:4222")
	if err != nil {
		log.Printf("messageSystem.New().error: %v", err.Error())
		t.FailNow()
	}
}
```

Em seguida, inicialize a simulação do debezium.

```go
package localDevOps

import (
	"log"
	"testing"
	"time"
)

func TestLocalDevOps(t *testing.T) {
	var err error
	
	var messageSystem = messagingSystemNats.MessagingSystemNats{}
	_, err = messageSystem.New("nats://0.0.0.0:4222")
	if err != nil {
		log.Printf("messageSystem.New().error: %v", err.Error())
		t.FailNow()
	}
	
	var dataSimulation = dataTest.DataTest{}
	var debezium = &debeziumSimulation.DebeziumSimulation{}
	debezium.EnableOnStartData(10)
	debezium.SetData(&dataSimulation)
	debezium.SetMessagingSystem(&messageSystem)
	debezium.SetMessagingTopic("dataSimulation")
	debezium.SetTimers(
		50*time.Millisecond,
		500*time.Millisecond,
		700*time.Millisecond,
		1000*time.Millisecond,
		60*time.Minute,
	)
	
	err = debezium.Init(false, "tradersclub", "simulation")
	if err != nil {
		log.Printf("messageSystem.New().error: %v", err.Error())
		t.FailNow()
	}
	
	ch := debezium.GetTerminationChannel()
	<-ch
	
	log.Print("fim da simulação")
}
```

## Envio de mensagens

Caso você não esteja familiarizado com debezium, ele envia mensagens quando o dado contido no banco de 
dados muda.

Os principais pontos da mensagem são:

  * **source:** dados do banco de dados, onde você deve ficar de olho no campos **bd** e **table** com os 
	respectivos nomes do banco de dados e da table;
  * **before:** dado antes da alteração;
  * **after:** dado depois da alteração;
  * **op:** operação realizada, pode ser `c` create; `r` read; `u` update; `d` delete e `z` fim da simulação.

Veja exemplos para o dado criado acima.

Crate:
```json
{
  "source": {
    "version": "1.5.0.Final",
    "connector": "mysql",
    "name": "mysql",
    "ts_ms": 1622826301925,
    "snapshot": true,
    "db": "tradersclub",
    "sequence": null,
    "table": "simulation",
    "server_id": 0,
    "gtid": null,
    "file": "mysql-bin.000008",
    "pos": 82625,
    "row": 0,
    "thread": null,
    "query": null
  },
  "before": null,
  "after": {
    "Id": "93cec220-8278-49c2-86bd-994641d239b8",
    "Name": "Adelle Hamill"
  },
  "op": "c",
  "ts_ms": 1648133772,
  "transaction": null
}
```

Update:
```json
{
  "source": {
    "version": "1.5.0.Final",
    "connector": "mysql",
    "name": "mysql",
    "ts_ms": 1622826301925,
    "snapshot": true,
    "db": "tradersclub",
    "sequence": null,
    "table": "simulation",
    "server_id": 0,
    "gtid": null,
    "file": "mysql-bin.000008",
    "pos": 82625,
    "row": 0,
    "thread": null,
    "query": null
  },
  "before": {
    "Id": "e3112e32-c819-4097-9ef6-43a31b2f6a5d",
    "Name": "Obie Windler"
  },
  "after": {
    "Id": "e3112e32-c819-4097-9ef6-43a31b2f6a5d",
    "Name": "Karolann Leannon"
  },
  "op": "u",
  "ts_ms": 1648133772,
  "transaction": null
}
```

Delete:
```json
{
  "source": {
    "version": "1.5.0.Final",
    "connector": "mysql",
    "name": "mysql",
    "ts_ms": 1622826301925,
    "snapshot": true,
    "db": "tradersclub",
    "sequence": null,
    "table": "simulation",
    "server_id": 0,
    "gtid": null,
    "file": "mysql-bin.000008",
    "pos": 82625,
    "row": 0,
    "thread": null,
    "query": null
  },
  "before": {
    "Id": "4f3d5501-aa95-4756-a5c0-dfa8807f5361",
    "Name": "Leonie Kerluke"
  },
  "after": null,
  "op": "d",
  "ts_ms": 1648133774,
  "transaction": null
}
```

Fim de simulação:

```json
{
  "source": {
    "version": "1.5.0.Final",
    "connector": "mysql",
    "name": "mysql",
    "ts_ms": 1622826301925,
    "snapshot": true,
    "db": "tradersclub",
    "sequence": null,
    "table": "simulation",
    "server_id": 0,
    "gtid": null,
    "file": "mysql-bin.000008",
    "pos": 82625,
    "row": 0,
    "thread": null,
    "query": null
  },
  "before": null,
  "after": null,
  "op": "z",
  "ts_ms": 1648135688,
  "transaction": null
}
```

## Sistema de mensageria

Caso você necessite refazer o sistema mensageria para algum outro diferente do [nats](https://nats.io/),
escreva seu código de modo a compatibilizar com a interface abaixo.

```go
package interfaces

import "test/support/commonTypes"

type MessagingSystemInterface interface {
	
	// Subscribe
	//
	// Adiciona uma função para ser invocada quando um novo evento ocorre no canal.
	//   Entrada:
	//     subject: nome do canal.
	//     function: Ponteiro de função a ser executada quando houver um evento no canal;
	//   Saída:
	//     err: Objeto de erro padrão do go.
	//
	//   Nota: 
	//     * O segundo parâmetro, de nome function, é o ponteiro de uma função invocada pelo código 
	//       quando ocorre um evento no canal, de nome contido em subject, e ela receberá o nome do 
	//       canal e o dado transmitido.
	//         Entrada:
	//           subject: nome do canal
	//           data: array de byte com a informação transmitida no canal (geralmente json)
	//         Saída:
	//           err: Objeto de erro padrão do go.
	Subscribe(subject string, function func(subject string, data []byte) (err error)) (err error)
	
	// SetReport
	//
	// Invoca uma função periodicamente para informar o status da conexão.
	SetReport(function func(status commonTypes.QueueStatus))
	
	// Publish
	//
	// Publica uma informação no canal específico da fila.
	//   Entrada:
	//     subject: Nome do canal;
	//     data: Array de bytes contendo a informação a ser compartilhada.
	//   Saída:
	//     err: Objeto de erro padrão do go.
	Publish(subject string, data []byte) (err error)
}
```
