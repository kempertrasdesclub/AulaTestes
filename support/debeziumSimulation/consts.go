package debeziumSimulation

const (
	KOperationCreate    = "c"
	KOperationRead      = "r"
	KOperationUpdate    = "u"
	KOperationDelete    = "d"
	KOperationTerminate = "z"

	KErrorMessagingTopicOnDeleteIsNotSet    = "messaging topic on delete is not set"
	KErrorMessagingTopicOnCreateIsNotSet    = "messaging topic on create is not set"
	KErrorMessagingTopicOnUpdateIsNotSet    = "messaging topic on update is not set"
	KErrorMessagingTopicOnStartIsNotSet     = "messaging topic on start is not set"
	KErrorMessagingTopicOnTerminateIsNotSet = "messaging topic on terminate is not set"
	KErrorUseSetDataFunctionFirst           = "use SetData() function first"
	KErrorMessagingInterfaceIsNotSet        = "messaging interface is not set"
)
