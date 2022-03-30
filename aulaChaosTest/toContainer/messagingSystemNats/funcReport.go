package messagingSystemNats

import (
	"github.com/nats-io/nats.go"
	"toContainer/commonTypes"
)

// report (português): Traduz o estado da conexão no Nats.
func (e *MessagingSystemNats) report() {
	if e.reportFunc == nil {
		return
	}

	var status commonTypes.QueueStatus
	var natsStatus = e.conn.Status()
	switch natsStatus {
	case nats.DISCONNECTED:
		status = commonTypes.KQueueDisconnected

	case nats.CONNECTED:
		status = commonTypes.KQueueConnected

	case nats.CLOSED:
		status = commonTypes.KQueueClosed

	case nats.RECONNECTING:
		status = commonTypes.KQueueReconnecting

	case nats.CONNECTING:
		status = commonTypes.KQueueConnecting
	}

	e.reportFunc(status)
}
