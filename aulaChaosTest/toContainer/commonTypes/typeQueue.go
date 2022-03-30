package commonTypes

import (
	"time"
)

type QueueStatus int

const (
	KQueueDisconnected QueueStatus = iota
	KQueueConnected
	KQueueClosed
	KQueueReconnecting
	KQueueConnecting
)

const (
	KQueueTickerPing = 5 * time.Second
)
