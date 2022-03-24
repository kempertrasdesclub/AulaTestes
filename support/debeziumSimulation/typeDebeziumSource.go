package debeziumSimulation

// debeziumSource objeto usado para simular o funcionamento do Debezium e recebe dados falsos
type debeziumSource struct {
	Version   string      `json:"version"`
	Connector string      `json:"connector"`
	Name      string      `json:"name"`
	TsMs      int64       `json:"ts_ms"`
	Snapshot  bool        `json:"snapshot"`
	Db        string      `json:"db"`
	Sequence  interface{} `json:"sequence"`
	Table     string      `json:"table"`
	ServerId  int64       `json:"server_id"`
	Gtid      interface{} `json:"gtid"`
	File      string      `json:"file"`
	Pos       int64       `json:"pos"`
	Row       int64       `json:"row"`
	Thread    interface{} `json:"thread"`
	Query     interface{} `json:"query"`
}
