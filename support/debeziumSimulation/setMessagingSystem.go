package debeziumSimulation

import "test/support/interfaces"

func (e *DebeziumSimulation) SetMessagingSystem(messagingSystem interfaces.MessagingSystemInterface) {
	e.messagingSystem = messagingSystem
}
