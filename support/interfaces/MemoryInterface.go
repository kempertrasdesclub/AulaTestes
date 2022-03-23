package interfaces

// MemoryInterface (português): Interface com as funções de memória usadas no projeto
// event.Event, para o CRUD de memória.
//
// O projeto event.Event foi feito para permitir o CDC, Change Data Capture, atualizar
// dados na memória cache, porém, nesse caso, set e update são feitos por set, get é
// desnecessário e delete é delete.
//
// Use está interface se a data informada do evento de mudança do dado não for crítica,
// caso contrário, use à interface, MemoryWithRemoteTimeClockInterface.
//
// Essa interface foi pensada para uma memória cache local onde todos os dados são
// atualizados apenas pelo CDC e enviados via sistema de mensagens.
type MemoryInterface interface {

	// Set (português): Salva o dado na memória
	//   Entrada:
	//     key: chave de identificação do dado salvo
	//     dataPointer: dado ou ponteiro do dado a ser salvo
	//   Saida:
	//     err: objeto padrão de erro
	Set(key, value interface{}) (err error)

	// Delete (português): Apaga a chave, caso exista.
	//   Entrada:
	//     Key: Uma ou mais chaves de identificação do dado, separadas por vírgula.
	//   Saida:
	//     err: Objecto padrão de erro.
	//
	//   Nota: - Se o sincronismo entre pods não estiver habilitado, o dado é apagado
	//           fisicamente, caso contrário, é apagado virtualmente.
	//         - A informação de hora do evento deve ser feita com
	//           `time.Now().Add(ntp.DiferenceBetweenClocks)` para aproximar ao máximo a hora
	//           dos eventos usando um servidor NTP, Network Time Protocol.
	//         - err é sempre nil e foi colocado por uma questão de compatibilidade com a
	//           interface{}
	Delete(key ...interface{}) (err error)
}
