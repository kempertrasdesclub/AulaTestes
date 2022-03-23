package interfaces

import (
	"test/support/commonTypes"
	"time"
)

type ParserReceiverInterface interface {
	Receiver(
		data interface{},
	) (
		event commonTypes.ReceiveEvent,
		eventDate time.Time,
		keyToCache,
		dataToCache interface{},
		err error,
	)
}
