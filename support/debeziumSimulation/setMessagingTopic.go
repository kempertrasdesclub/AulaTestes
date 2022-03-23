package debeziumSimulation

func (e *DebeziumSimulation) SetMessagingTopic(topic string) {
	e.messagingTopicOnStart = topic
	e.messagingTopicOnCreate = topic
	e.messagingTopicOnUpdate = topic
	e.messagingTopicOnDelete = topic
}
