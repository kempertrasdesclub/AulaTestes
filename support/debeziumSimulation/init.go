package debeziumSimulation

import (
	"github.com/helmutkemper/util"
	"log"
	"time"
)

// Init (português): Inicializa a simulação
//   Entrada:
//     length: quantidade de dados simulados
//   Saída:
//     err: Objeto padrão de erro
func (e *DebeziumSimulation) Init() (err error) {
	e.ErrChan = make(chan error)

	// Dados fantasia para MySQL
	e.Source.Version = "1.5.0.Final"
	e.Source.Connector = "mysql"
	e.Source.Name = "mysql"
	e.Source.TsMs = 1622826301925
	e.Source.Snapshot = true
	e.Source.Db = "tradersclub"
	e.Source.Sequence = nil
	e.Source.Table = "Idea"
	e.Source.ServerId = 0
	e.Source.Gtid = nil
	e.Source.File = "mysql-bin.000008"
	e.Source.Pos = 82625
	e.Source.Row = 0
	e.Source.Thread = nil
	e.Source.Query = nil

	log.Print("início: populando dados")
	start := time.Now()

	// fixme: os dados são criados todos de uma vez nessa abordagem e isto é um erro
	if e.sendOnPopulateData == true {
		for _, after := range e.create {
			err = e.SendOnPopulateData(after)
			if err != nil {
				util.TraceToLog()
				return
			}

			if e.sendOnStartDelay != 0 {
				time.Sleep(e.sendOnStartDelay)
			}
		}
	}

	log.Printf("fim: a cache foi populada sem erros")
	log.Printf("tempo total: %v", time.Since(start))
	log.Printf("inicio: simulando usuário")

	e.sendOnCreateTicker = time.NewTicker(e.sendOnCreateDelay)
	e.sendOnUpdateTicker = time.NewTicker(e.sendOnUpdateDelay)
	e.sendOnDeleteTicker = time.NewTicker(e.sendOnDeleteDelay)

	go func(e *DebeziumSimulation) {
		for {
			select {
			case <-e.sendOnCreateTicker.C:
				//fixme: isto é um bug
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
