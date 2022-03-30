package messagingSystemNats

import "toContainer/commonTypes"

// SetReport (Português): Invoca uma função periodicamente para informar o status da
// conexão.
func (e *MessagingSystemNats) SetReport(function func(status commonTypes.QueueStatus)) {
	e.reportFunc = function
}
