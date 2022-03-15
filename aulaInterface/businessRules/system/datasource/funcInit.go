package datasource

import (
	"errors"
	"github.com/helmutkemper/util"
	"plugin"
	"test/aulaInterface/constants"
	"test/aulaInterface/interfaces"
	jwtverify "test/aulaInterface/toModule/JWT"
	"test/aulaInterface/toModule/passwordHash"
	"test/aulaInterface/toModule/uID"
)

// Init (Português): Inicializa o datasource escolhido
//   name: tyme Name
//     KSQLite: Inicializa o banco de dados como sendo o SQLite
func (e *RefList) Init(name Name) (err error) {

	var path string

	err = errors.New(constants.KErrorInicializeDataSourceFirst)

	// Inicializa o objeto Password
	e.Password = &passwordHash.Password{}

	// Inicializa o objeto UID
	e.UniqueID = &uID.UID{}

	// Inicializa o gerador/verificador de JWT
	e.Jwt = &jwtverify.JwtVerify{}
	err = e.Jwt.NewAlgorithm([]byte("colocar em constants")) //fixme
	if err != nil {
		util.TraceToLog()
		return
	}

	// Inicializa o banco de dados
	switch name {

	case KMongoDB:

		// user
		path, err = util.FileFindInTree("user.mongodb.so")
		if err != nil {
			util.TraceToLog()
			return
		}

		err = e.installUserByPlugin(path)
		if err != nil {
			util.TraceToLog()
			return
		}

	case KSQLite:
		path, err = util.FileFindInTree("user.sqlite.so")
		if err != nil {
			util.TraceToLog()
			return
		}

		err = e.installUserByPlugin(path)
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	return
}

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
