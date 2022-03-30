package messagingSystemNats

// Disconnect (Português): desconecta do servidor de fila Nats
func (e *MessagingSystemNats) Disconnect() {
	e.tickerStop <- true
	e.conn.Close()
}
