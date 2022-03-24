package debeziumSimulation

import (
	"test/support/interfaces"
	"time"
)

// DebeziumSimulation
//
// Objeto usado como memória na hora de criar dados alterados pelo usuário.
//
//   Nota:
//     * O dado representando a tabela do banco de dados fica contido em RealDataToSimulate.
type DebeziumSimulation struct {
	realDataPointer interfaces.DataToSimulateInterface

	// id: último id gerado, onde id é compatível com RealDataToSimulate.Id
	id interface{}

	// create: lista de todos os dados criados
	create map[interface{}]FileLineFormat

	// update: lista de todos os dados atualizados
	update map[interface{}]FileLineFormat

	// delete: lista de todos os dados apagados
	delete map[interface{}]FileLineFormat

	enableSaveData bool

	sendOnPopulateData     int
	messagingSystem        interfaces.MessagingSystemInterface
	messagingTopicOnStart  string
	messagingTopicOnCreate string
	messagingTopicOnUpdate string
	messagingTopicOnDelete string

	sendOnCreateTicker *time.Ticker
	sendOnUpdateTicker *time.Ticker
	sendOnDeleteTicker *time.Ticker

	sendOnStartDelay  time.Duration
	sendOnCreateDelay time.Duration
	sendOnUpdateDelay time.Duration
	sendOnDeleteDelay time.Duration

	sendTestProcessTerminationTimer *time.Timer
	sendTestProcessTerminationDelay time.Duration

	ErrChan         chan error    `json:"-"`
	TerminationChan chan struct{} `json:"-"`

	Source      debeziumSource `json:"source"`
	Before      interface{}    `json:"before"`
	After       interface{}    `json:"after"`
	Operation   string         `json:"op"`
	EventDate   int64          `json:"ts_ms"`
	Transaction interface{}    `json:"transaction"`
}
