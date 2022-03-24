package debeziumSimulation

// SetMessagingTopic
//
// Define o tópico do sistema de mensageria.
//
//   Entrada:
//     topic: texto identificador do tópico do sistema de mensageria.
//
//   Nota:
//     * Use as funções SetMessagingTopicOnCreate(), SetMessagingTopicOnDelete(),
//       SetMessagingTopicOnStart() e SetMessagingTopicOnUpdate() para definir tópicos específicos.
func (e *DebeziumSimulation) SetMessagingTopic(topic string) {
	e.messagingTopicOnStart = topic
	e.messagingTopicOnCreate = topic
	e.messagingTopicOnUpdate = topic
	e.messagingTopicOnDelete = topic
}
