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
//       SetMessagingTopicOnStart(), SetMessagingTopicOnUpdate() e SetMessagingTopicOnTerminate() para
//       definir tópicos específicos.
//     * Use a função SetMessagingTopic() para definir todos os tópicos simultaneamente, e em seguida
//       use as demais funções para definir um tópico específico.
func (e *DebeziumSimulation) SetMessagingTopic(topic string) {
	e.messagingTopicOnStart = topic
	e.messagingTopicOnCreate = topic
	e.messagingTopicOnUpdate = topic
	e.messagingTopicOnDelete = topic
	e.messagingTopicOnTerminate = topic
}
