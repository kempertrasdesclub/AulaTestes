package managerDatasource

import (
	"errors"
	"github.com/helmutkemper/util"
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/interfaces"
	"plugin"
)

// installUserByPlugin
//
// Português:
//
// Carrega o arquivo binário externo contendo o pacote de dados `User`.
//   Entrada:
//     pluginPlath: caminho do arquivo contendo o pacote de dados `User`
//   Saída:
//     err: objeto padrão de erro do go
func (e *RefList) installUserByPlugin(pluginPlath string) (err error) {
	var ok bool
	var user *plugin.Plugin
	var userSymbol plugin.Symbol

	user, err = plugin.Open(pluginPlath)
	if err != nil {
		util.TraceToLog()
		return
	}

	userSymbol, err = user.Lookup("User")
	if err != nil {
		util.TraceToLog()
		return
	}

	e.User, ok = userSymbol.(interfaces.InterfaceUser)
	if ok == false {
		err = errors.New("plugin user conversion into interface user has an error")
		util.TraceToLog()
		return
	}

	_, err = e.User.New()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
