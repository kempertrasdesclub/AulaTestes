package debeziumSimulation

// SetMessagingTopicOnDelete
//
// Define o tópico do sistema de mensageria quando o dado é apagado.
//
//   Entrada:
//     topic: texto identificador do tópico do sistema de mensageria.
//
//   Nota:
//     * Use as funções SetMessagingTopicOnCreate(), SetMessagingTopicOnDelete(),
//       SetMessagingTopicOnStart() e SetMessagingTopicOnUpdate() para definir tópicos específicos.
func (e *DebeziumSimulation) SetMessagingTopicOnDelete(topic string) {
	e.messagingTopicOnDelete = topic
}
