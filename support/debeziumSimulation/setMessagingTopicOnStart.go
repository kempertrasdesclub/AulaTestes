package debeziumSimulation

// SetMessagingTopicOnStart
//
// Define o tópico do sistema de mensageria quando ocorre a carga inicial de dados.
//
//   Entrada:
//     topic: texto identificador do tópico do sistema de mensageria.
//
//   Nota:
//     * Use as funções SetMessagingTopicOnCreate(), SetMessagingTopicOnDelete(),
//       SetMessagingTopicOnStart() e SetMessagingTopicOnUpdate() para definir tópicos específicos.
func (e *DebeziumSimulation) SetMessagingTopicOnStart(topic string) {
	e.messagingTopicOnStart = topic
}
