package managerDatasource

import (
	"errors"
	"github.com/helmutkemper/util"
	"test/aulaInterface/constants"
	jwtverify "test/aulaInterface/toModule/JWT"
	"test/aulaInterface/toModule/passwordHash"
	"test/aulaInterface/toModule/uID"
)

// Init Inicializa o datasource escolhido
//   name: type Name
//     KFakeData: Inicializa o banco de dados como sendo o objeto com dados aleatórios.
//     KSQLite:   Inicializa o banco de dados como sendo o SQLite
//     KMongoDB:  Inicializa o banco de dados como sendo o MongoDB
func (e *RefList) Init(name Name) (err error) {

	var userPluginPath string

	err = errors.New(constants.KErrorInicializeDataSourceFirst)

	// Inicializa o objeto Password
	e.Password = &passwordHash.Password{}

	// Inicializa o objeto UID
	e.UniqueID = &uID.UID{}

	// Inicializa o gerador/verificador de JWT
	e.Jwt = &jwtverify.JwtVerify{}
	err = e.Jwt.NewAlgorithm([]byte(kSecretKey))
	if err != nil {
		util.TraceToLog()
		return
	}

	// Inicializa o banco de dados
	switch name {

	case KFakeData:
		userPluginPath, err = util.FileFindInTree(kPluginUserFakeData)
		if err != nil {
			util.TraceToLog()
			return
		}

	case KMongoDB:
		userPluginPath, err = util.FileFindInTree(kPluginUserMongoDB)
		if err != nil {
			util.TraceToLog()
			return
		}

	case KSQLite:
		userPluginPath, err = util.FileFindInTree(kPluginUserSQLite)
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	err = e.installUserByPlugin(userPluginPath)
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
