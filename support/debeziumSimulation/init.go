package debeziumSimulation

import (
	"log"
	"time"
)

// Init
//
// Inicializa a simulação
//
//   Entrada:
//     enableSaveData: habilita o salvamento de dados para uso posterior;
//     dbName: nome do banco de dados usado na simulação;
//     tableName: nome da tabela usada na simulação.
//
//   Saída:
//     err: Objeto padrão de erro do go.
func (e *DebeziumSimulation) Init(enableSaveData bool, dbName, tableName string) (err error) {
	e.TerminationChan = make(chan struct{})
	e.ErrChan = make(chan error)

	// Dados fantasia para MySQL
	e.Source.Version = "1.5.0.Final"
	e.Source.Connector = "mysql"
	e.Source.Name = "mysql"
	e.Source.TsMs = 1622826301925
	e.Source.Snapshot = true
	e.Source.Db = dbName
	e.Source.Sequence = nil
	e.Source.Table = tableName
	e.Source.ServerId = 0
	e.Source.Gtid = nil
	e.Source.File = "mysql-bin.000008"
	e.Source.Pos = 82625
	e.Source.Row = 0
	e.Source.Thread = nil
	e.Source.Query = nil

	e.enableSaveData = enableSaveData

	log.Print("início: populando dados")
	start := time.Now()

	for i := 0; i != e.sendOnPopulateData; i += 1 {
		e.actionReadData()

		if e.sendOnStartDelay != 0 {
			time.Sleep(e.sendOnStartDelay)
		}
	}

	log.Printf("fim: a cache foi populada sem erros")
	log.Printf("tempo total: %v", time.Since(start))
	log.Printf("inicio: simulando usuário")

	e.sendOnCreateTicker = time.NewTicker(e.sendOnCreateDelay)
	e.sendOnUpdateTicker = time.NewTicker(e.sendOnUpdateDelay)
	e.sendOnDeleteTicker = time.NewTicker(e.sendOnDeleteDelay)
	e.sendTestProcessTerminationTimer = time.NewTimer(e.sendTestProcessTerminationDelay)

	go func(e *DebeziumSimulation) {
		for {
			select {
			case <-e.sendTestProcessTerminationTimer.C:
				e.sendTestProcessTerminationTimer.Stop()
				e.sendOnCreateTicker.Stop()
				e.sendOnUpdateTicker.Stop()
				e.sendOnDeleteTicker.Stop()

				e.actionSimulationEnd()
				e.TerminationChan <- struct{}{}
				log.Printf("-------------------------------------------------------------------------------------------------------------------------")

				return

			case <-e.sendOnCreateTicker.C:
				go e.actionCreateData()

			case <-e.sendOnUpdateTicker.C:
				go e.actionUpdateData()

			case <-e.sendOnDeleteTicker.C:
				go e.actionDeleteData()
			}
		}
	}(e)

	return
}
