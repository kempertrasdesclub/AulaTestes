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
//       SetMessagingTopicOnStart() e SetMessagingTopicOnUpdate() para definir tópicos específicos.
func (e *DebeziumSimulation) SetMessagingTopicOnUpdate(topic string) {
	e.messagingTopicOnUpdate = topic
}
