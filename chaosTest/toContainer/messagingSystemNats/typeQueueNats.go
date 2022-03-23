package messagingSystemNats

import (
	"github.com/nats-io/nats.go"
	"sync"
	"time"
	"toContainer/commonTypes"
)

type MessagingSystemNats struct {
	errorCounter     int
	connectionString string
	conn             *nats.Conn
	options          []nats.Option
	publishList      map[string][]func(subject string, data []byte) (err error)
	mutex            sync.RWMutex
	ticker           *time.Ticker
	tickerStop       chan bool
	reportFunc       func(status commonTypes.QueueStatus)
}
