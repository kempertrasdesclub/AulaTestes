package messagingSystemNats

// tickerRun (Português): Roda a intervalos regulares para reportar problemas de conexão
func (e *MessagingSystemNats) tickerRun() {
	go func(e *MessagingSystemNats) {
		for {
			select {
			case <-e.tickerStop:
				return

			case <-e.ticker.C:
				e.report()
			}
		}
	}(e)
}
