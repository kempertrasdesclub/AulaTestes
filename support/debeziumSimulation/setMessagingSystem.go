package debeziumSimulation

import (
	"github.com/kempertrasdesclub/AulaTestes/support/interfaces"
)

// SetMessagingSystem
//
// Recebe o ponteiro do objeto de controle do sistema de mensageria.
//
//   Entrada:
//     messagingSystem: ponteiro do objeto de controle do sistema de mensageria compatível com a
//       interfaces.MessagingSystemInterface.
func (e *DebeziumSimulation) SetMessagingSystem(messagingSystem interfaces.MessagingSystemInterface) {
	e.messagingSystem = messagingSystem
}
