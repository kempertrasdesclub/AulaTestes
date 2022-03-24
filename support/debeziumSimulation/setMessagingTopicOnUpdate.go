package debeziumSimulation

// SetMessagingTopicOnUpdate
//
// Define o tópico do sistema de mensageria quando o dado é atualizado.
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
func (e *DebeziumSimulation) SetMessagingTopicOnUpdate(topic string) {
	e.messagingTopicOnUpdate = topic
}
