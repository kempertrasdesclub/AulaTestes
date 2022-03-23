package interfaces

import (
	"time"
)

// MemoryWithRemoteTimeClockInterface (português): Interface com as funções de memória
// usadas no projeto event.Event, para o CRUD de memória.
//
// O projeto event.Event foi feito para permitir o CDC, Change Data Capture, atualizar
// dados na memória cache, porém, nesse caso, set e update são feitos por set, get é
// desnecessário e delete é delete.
//
// Use está interface para definir a data do evento com o dado como sendo a data do dado
// na memória. Isto evita problemas de sincronismo.
//
// Esta interface foi pensada para ser usada com sincronismo de dados entre pods, onde os
// dados atualizados em um pod são sincronizados com os demais pods.
// Como o sincronismo usa a data em que o dado foi gravado para determinar o dado mais
// novo, é importante que a data informada no CDC seja usada como data de gravação do
// dado, poupando processamento desnecessário no sincronismo do dado.
type MemoryWithRemoteTimeClockInterface interface {

	// SetWithRemoteTimeClock (português): Salva o dado na memória
	//   Entrada:
	//     key: chave de identificação do dado salvo
	//     dataPointer: dado ou ponteiro do dado a ser salvo
	//     eventTime: objeto padrão time.Time com a data do evento no servidor remoto.
	//   Saida:
	//     err: objeto padrão de erro
	SetWithRemoteTimeClock(key, dataPointer interface{}, eventTime time.Time) (err error)

	// DeleteWithRemoteTimeClock (português): Apaga a chave, caso exista.
	//   Entrada:
	//     eventTime: objeto padrão time.Time com a data do evento no servidor remoto.
	//     key: n chaves de identificação do dado.
	//
	//   Nota: - Se o sincronismo entre pods não estiver habilitado, o dado é apagado
	//           fisicamente, caso contrário, é apagado virtualmente.
	//         - A informação de hora do evento deve ser feita com
	//           `time.Now().Add(ntp.DiferenceBetweenClocks)` para aproximar ao máximo a hora
	//           dos eventos usando um servidor NTP, Network Time Protocol.
	//         - Esta função impede que dados mais antigos sejam apagados e só funciona com o
	//           servidor de sincronismo ativo
	DeleteWithRemoteTimeClock(eventTime time.Time, key ...interface{}) (err error)
}
