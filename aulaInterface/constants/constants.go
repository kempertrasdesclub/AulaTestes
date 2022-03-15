package constants

// main menu
const (
	KMainMenuTypeContentAsMenu     = 0
	KMainMenuTypeContentAsContent  = 1
	KmainMenuTypeClassRoomAsNormal = 0
	KmainMenuTypeClassRoomAsClass  = 1
	KmainMenuUserNormal            = 0
	KmainMenuUserAdmin             = 1
)

const (
	KMainMenuSiteNameID    = "04daa10f-1059-446a-a2b9-b48d536c8b23"
	KMainMenuSiteNameText  = "Kemper.com.br"
	KMainMenuSiteNameIcon  = "fas fa-code-branch"
	KMainMenuSiteNameURL   = ""
	KMainMenuSiteNameOrder = 0

	KMainMenuAboutMeID    = "d791c1a2-fc36-454a-a7f4-038149e30e4a"
	KMainMenuAboutMeText  = "Sobre mim"
	KMainMenuAboutMeIcon  = "fas fa-info-circle"
	KMainMenuAboutMeURL   = ""
	KMainMenuAboutMeOrder = 0

	KMainMenuGithubID    = "d40894ba-8015-402f-9c10-5a4836413b1d"
	KMainMenuGithubText  = "Github"
	KMainMenuGithubIcon  = "fab fa-github"
	KMainMenuGithubURL   = "https://github.com/helmutkemper"
	KMainMenuGithubOrder = 1

	KMainMenuLinkedinID    = "873da901-8285-4846-85df-59f16850429e"
	KMainMenuLinkedinText  = "LinkedIn"
	KMainMenuLinkedinIcon  = "fab fa-linkedin"
	KMainMenuLinkedinURL   = "https://www.linkedin.com/in/helmut-kemper-93a5441b/"
	KMainMenuLinkedinOrder = 2

	KMainMenuCodeID    = "0409921a-600b-49db-b8cc-d906e3837082"
	KMainMenuCodeText  = "Códigos"
	KMainMenuCodeIcon  = "fas fa-code"
	KMainMenuCodeURL   = ""
	KMainMenuCodeOrder = 1

	KMainMenuTalkingWithDevID    = "207b66f1-eef9-4572-a57d-5152b511e37d"
	KMainMenuTalkingWithDevText  = "Conversando com o sênior"
	KMainMenuTalkingWithDevIcon  = "fas fa-fire-extinguisher"
	KMainMenuTalkingWithDevURL   = ""
	KMainMenuTalkingWithDevOrder = 0

	KMainMenuMigratingToGoID    = "ae116cdf-1941-4ca7-9129-0bad2793ca42"
	KMainMenuMigratingToGoText  = "Migrando para o Go"
	KMainMenuMigratingToGoIcon  = "fas fa-fire"
	KMainMenuMigratingToGoURL   = ""
	KMainMenuMigratingToGoOrder = 1

	KMainMenuLoginID    = "08c7e064-8225-4ee4-9cf4-19dcdd546c84"
	KMainMenuLoginText  = "Login"
	KMainMenuLoginIcon  = ""
	KMainMenuLoginURL   = ""
	KMainMenuLoginOrder = 1

	KMainMenuNewContentID    = "24fda64b-0af2-4da6-9d95-38664614d3bd"
	KMainMenuNewContentText  = "New content"
	KMainMenuNewContentIcon  = ""
	KMainMenuNewContentURL   = ""
	KMainMenuNewContentOrder = 0

	KMainMenuAdminID    = "1726ce36-7338-4932-858b-6863700dbe51"
	KMainMenuAdminText  = "Admin"
	KMainMenuAdminIcon  = "fas fa-cogs"
	KMainMenuAdminURL   = ""
	KMainMenuAdminOrder = 2

	KMainMenuDonationID    = "d6667be4-6794-4c6f-9eb2-0d3cac1d7801"
	KMainMenuDonationText  = "Donation"
	KMainMenuDonationIcon  = "fas fa-donate"
	KMainMenuDonationURL   = ""
	KMainMenuDonationOrder = 3
)

// main user
const (
	KMainUserID       = "24867707-3a21-4368-8058-3d7b1ddb8c06"
	KMainUserName     = "Helmut Kemper"
	KMainUserNickName = "Kemper"
	KMainUserMail     = "helmut.kemper@gmail.com"
	KMainUserPassword = "admin"
)

const (
	KInstallLanguageID   = "f9c71382-40e8-4027-aa7d-b31faacc8111"
	KInstallLanguageName = "GoLang"
)

const (
	KErrorInicializeDataSourceFirst    = "please, inicialize data source first"
	KErrorPasswordOneSpecialChar       = "the password must be one special char"
	KErrorPasswordOneLowerCaseChar     = "the password must be one lower case char"
	KErrorPasswordMustBe8LettersOrMore = "the password must be 8 letters or more"
	KErrorUserNotFound                 = "user not found"
	KErrorLanguageTableEmpty           = "language table is empty"
	KErrorEmailValidSintax             = "e-mail must be a valid sintax"
)

const (
	KMongoDBDatabase           = "site"
	KMongoDBCollectionLanguage = "language"
	KMongoDBCollectionUser     = "user"
	KMongoDBCollectionMenu     = "menu"
)

const (
	//KSQLiteConnectionString (Português): Caminho e nome do arquivo de dados do SQLite criado para os testes
	KSQLiteConnectionString = "./database.sqlite"

	KMongoDBConnectionString = "mongodb://127.0.0.1:27017/"
)
