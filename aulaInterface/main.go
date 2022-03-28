package main

import (
	"fmt"
	builder "github.com/helmutkemper/iotmaker.docker.builder"
	"github.com/helmutkemper/util"
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/businessRules/system/managerDatasource"
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/gin/server"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	KColorTerminalBlack   = "\u001b[30m"
	KColorTerminalRed     = "\u001b[31m"
	KColorTerminalGreen   = "\u001b[32m"
	KColorTerminalYellow  = "\u001b[33m"
	KColorTerminalBlue    = "\u001b[34m"
	KColorTerminalMagenta = "\u001b[35m"
	KColorTerminalCyan    = "\u001b[36m"
	KColorTerminalWhite   = "\u001b[37m"
	KColorTerminalReset   = "\u001b[0m"

	KColorTerminalQuestion = KColorTerminalYellow
	KColorTerminalAnswer   = KColorTerminalGreen
	KColorTerminalWarning  = KColorTerminalRed
)

var querySelectDataSource string
var queryInstallMongoDB string
var queryUninstallMongoDB string
var queryRemoveImageMongo string

func main() {
	var err error
	var userDatasource managerDatasource.Name

	// formato do log
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Faz perguntas ao usuário
	userSetup()

	// captura de sinal para detectar o encerramento do programa e remover os containers
	var sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		removeDockerElements()
		os.Exit(0)
	}()

	// Pergunta ao usuário sobre a necessidade de instalar o MongoDB.
	if querySelectDataSource == "m" {
		// limpa elementos docker de uma instalação anterior
		builder.SaGarbageCollector()

		// instala o mongodb
		err = installEphemeralMongodb()
		if err != nil {
			util.TraceToLog()
			log.Panic(err)
		}
	}

	// Define a fonte de dados para o módulo `User`, usado na demonstração.
	switch querySelectDataSource {
	case "f":
		userDatasource = managerDatasource.KFakeData
	case "s":
		userDatasource = managerDatasource.KSQLite
	case "m":
		userDatasource = managerDatasource.KMongoDB
	}

	// inicializa a fonte de dados (SQLite ou MongoDB)
	err = managerDatasource.Linker.Init(userDatasource)
	if err != nil {
		panic(err)
	}

	// Abre a URL de demonstração.
	var cmd = exec.Command("open", "http://localhost:3000/datasource/user/helmut.kemper@gmail.com")
	_, err = cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	// inicializa o servidor
	err = server.ConfigAndStart()
	if err != nil {
		util.TraceToLog()
		log.Panic(err)
	}

	// Faça uma requisição GET para http://localhost:3000/datasource/user/helmut.kemper@gmail.com
	// Este é parte do módulo de usuários feito para funcionar independe da fonte de dados escolhida.
}

// Determina as configurações iniciais do sistema.
func userSetup() {
	userSetupSelectDatasource()

	if querySelectDataSource == "m" {
		userSetupInstallMongoDB()
	}
}

// Explica o código e faz perguntas ao usuário
func userSetupSelectDatasource() {
	log.Print("")
	log.Print(KColorTerminalWarning, "Atenção: Este código requer o docker instalado para instalar um banco de dados MongoDB temporário!", KColorTerminalReset)
	log.Print("")
	log.Print("O código foi feito para rodar com as fontes de dados, fake, SQLite ou MongoDB.")
	log.Print("Caso você escolha o MongoDB, este código tem a possibilidade de o instalar no docker de forma automática.")
	log.Print(KColorTerminalQuestion, "Qual fonte de dados você quer usar?", KColorTerminalReset)
	log.Print(KColorTerminalAnswer, "`f` para fake", KColorTerminalReset)
	log.Print(KColorTerminalAnswer, "`s` para SQLite", KColorTerminalReset)
	log.Print(KColorTerminalAnswer, "`m` para MongoDB", KColorTerminalReset)
	_, _ = fmt.Scan(&querySelectDataSource)

	querySelectDataSource = strings.ToLower(querySelectDataSource)
	if !(querySelectDataSource == "f" || querySelectDataSource == "s" || querySelectDataSource == "m") {
		log.Print("")
		log.Print(KColorTerminalWarning, "Por favor, digite `f` para fake, `s` para SQLite ou `m` para MongoDB.", KColorTerminalReset)
		userSetupSelectDatasource()
	}
}

// Pergunta ao usuário sobre a instalação do MongoDB
func userSetupInstallMongoDB() {
	log.Print(KColorTerminalQuestion, "Instalar o MongoDB no docker para teste?", KColorTerminalReset)
	log.Print(KColorTerminalAnswer, "`s` para sim", KColorTerminalReset)
	log.Print(KColorTerminalAnswer, "`n` para não", KColorTerminalReset)
	_, _ = fmt.Scan(&queryInstallMongoDB)

	queryInstallMongoDB = strings.ToLower(queryInstallMongoDB)
	if !(queryInstallMongoDB == "s" || queryInstallMongoDB == "n") {
		log.Print("")
		log.Print(KColorTerminalWarning, "Por favor, digite `s` para sim ou `n` para não.", KColorTerminalReset)
		userSetupInstallMongoDB()
	} else {
		log.Print("Por favor, aguarde...")
	}
}

// Pergunta ao usuário sobre a remoção do MongoDB ao fim da demonstração.
func userSetupRemoveContainerMongoDB() {
	log.Print(KColorTerminalQuestion, "Remover o container do MongoDB do docker?", KColorTerminalReset)
	log.Print(KColorTerminalAnswer, "`s` para sim", KColorTerminalReset)
	log.Print(KColorTerminalAnswer, "`n` para não", KColorTerminalReset)
	_, _ = fmt.Scan(&queryUninstallMongoDB)

	queryUninstallMongoDB = strings.ToLower(queryUninstallMongoDB)
	if !(queryUninstallMongoDB == "s" || queryUninstallMongoDB == "n") {
		log.Print("")
		log.Print(KColorTerminalWarning, "Por favor, digite `s` para sim ou `n` para não.", KColorTerminalReset)
		userSetupRemoveContainerMongoDB()
	} else {
		log.Print("Por favor, aguarde...")
	}
}

// Pergunta ao usuário sobre a remoção da imagem do MongoDB ao fim da demonstração.
func userSetupUninstallMongoDB() {
	log.Print(KColorTerminalQuestion, "Remover a imagem `mongo:latest` do docker?", KColorTerminalReset)
	log.Print(KColorTerminalAnswer, "`s` para sim", KColorTerminalReset)
	log.Print(KColorTerminalAnswer, "`n` para não", KColorTerminalReset)
	_, _ = fmt.Scan(&queryRemoveImageMongo)

	queryRemoveImageMongo = strings.ToLower(queryRemoveImageMongo)
	if !(queryRemoveImageMongo == "s" || queryRemoveImageMongo == "n") {
		log.Print("")
		log.Print(KColorTerminalWarning, "Por favor, digite `s` para sim ou `n` para não.", KColorTerminalReset)
		userSetupUninstallMongoDB()
	} else {
		log.Print("Por favor, aguarde...")
	}
}

// Concentra todas as etapas de remoção dos elementos docker.
func removeDockerElements() {
	// Se a fonte de dados não for o MongoDB, não há elementos docker instalados.
	if querySelectDataSource != "m" {
		return
	}

	userSetupRemoveContainerMongoDB()

	if queryUninstallMongoDB == "s" {
		builder.SaGarbageCollector()

		userSetupUninstallMongoDB()
		if queryRemoveImageMongo == "s" {
			builder.SaGarbageCollector("mongo:latest")
		}
	}
}

// Instala um MongoDb com dados efêmeros no docker.
func installEphemeralMongodb() (err error) {
	var mongoDocker = &builder.ContainerBuilder{}
	mongoDocker.SetImageName("mongo:latest")
	mongoDocker.SetContainerName("container_delete_mongo_after_test")
	mongoDocker.AddPortToExpose("27017")
	mongoDocker.SetPrintBuildOnStrOut()
	mongoDocker.SetEnvironmentVar(
		[]string{
			"--host 0.0.0.0",
		},
	)

	mongoDocker.SetWaitStringWithTimeout(`"msg":"Waiting for connections","attr":{"port":27017`, 20*time.Second)
	err = mongoDocker.Init()
	if err != nil {
		util.TraceToLog()
		return
	}

	err = mongoDocker.ImagePull()
	if err != nil {
		util.TraceToLog()
		return
	}

	err = mongoDocker.ContainerBuildAndStartFromImage()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
