package managerDatasource

const (
	// KSQLite
	//
	// Português:
	//
	// Define o datasource como sendo o SQLite3
	KSQLite Name = iota

	// KMongoDB
	//
	// Português:
	//
	// Define o datasource como sendo o MongoDB
	KMongoDB

	// KFakeData
	//
	// Português:
	//
	// Define o datasource como sendo dados aleatórios
	KFakeData

	// kSecretKey
	//
	// Português:
	//
	//
	kSecretKey = "axrvdsoimn653hgdcicbegrh"

	// kPluginUserFakeData
	//
	// Português:
	//
	// Nome do arquivo externo do plugin `User` para dados aleatórios.
	kPluginUserFakeData = "user.fake.so"

	// kPluginUserMongoDB
	//
	// Português:
	//
	// Nome do arquivo externo do plugin `User` para o MongoDB.
	kPluginUserMongoDB = "user.mongodb.so"

	// kPluginUserSQLite
	//
	// Português:
	//
	// Nome do arquivo externo do plugin `User` para o SQLite.
	kPluginUserSQLite = "user.sqlite.so"
)
