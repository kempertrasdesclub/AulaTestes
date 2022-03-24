package debeziumSimulation

// SetMessagingTopicOnCreate
//
// Define o tópico do sistema de mensageria quando um novo dado é criado.
//
//   Entrada:
//     topic: texto identificador do tópico do sistema de mensageria.
//
//   Nota:
//     * Use as funções SetMessagingTopicOnCreate(), SetMessagingTopicOnDelete(),
//       SetMessagingTopicOnStart() e SetMessagingTopicOnUpdate() para definir tópicos específicos.
func (e *DebeziumSimulation) SetMessagingTopicOnCreate(topic string) {
	e.messagingTopicOnCreate = topic
}
