package localDevOps

import (
	dockerBuilder "github.com/helmutkemper/iotmaker.docker.builder"
	dockerBuilderNetwork "github.com/helmutkemper/iotmaker.docker.builder.network"
	"github.com/helmutkemper/util"
	"log"
	"strconv"
	"time"
)

func dockerSimulationInstall(
	netDocker *dockerBuilderNetwork.ContainerBuilderNetwork,
	simulation *dockerBuilder.ContainerBuilder,
	index int64,
	memoryInputPath string,
	enableChaos bool,
) (
	err error,
) {

	var indexAsString = strconv.FormatInt(index, 10)

	log.Printf("Instalação do container de simulação, índice %v: início", indexAsString)

	simulation.SetImageCacheName("cachechaostest:latest")
	simulation.SetImageExpirationTime(10 * time.Minute) //fixme: definir no arquivo de teste
	// Imprime a saída padrão do container na saída padrão do golang
	simulation.SetPrintBuildOnStrOut()
	// Habilita o uso da imagem cache:latest
	simulation.SetCacheEnable(true)
	// Aponta o gerenciador de rede
	simulation.SetNetworkDocker(netDocker)
	// Determina o nome da imagem a ser usada
	simulation.SetImageName("delete_after_test:latest")
	// Define o nome do container
	simulation.SetContainerName("container_delete_after_test_" + indexAsString)
	// Define o caminho da pasta contando o projeto
	simulation.SetBuildFolderPath("../chaosTest/toContainer")
	// Gera o Dockerfile de forma automática
	simulation.MakeDefaultDockerfileForMe()
	// Define os repositórios da TC como sendo privados
	simulation.SetGitPathPrivateRepository("github.com/tradersclub")
	// Copias as credenciais do usuário para o container
	err = simulation.SetPrivateRepositoryAutoConfig()
	if err != nil {
		util.TraceToLog()
		log.Printf("Error: %v", err.Error())
		return
	}
	log.Printf("memoryInputPath: %v", memoryInputPath)
	err = simulation.AddFileOrFolderToLinkBetweenConputerHostAndContainer(memoryInputPath, "/memory_container")
	if err != nil {
		util.TraceToLog()
		log.Printf("Error: %v", err.Error())
		return
	}

	// Adiciona um indicador de falha com gravação de arquivo em log ao projeto.
	// Indicador de falha é um texto procurado na saída padrão do container e indica algo que não deveria ter acontecido durante o teste.
	// Algumas falhas críticas podem ser monitoradas e quando elas acontecem, a saída padrão do container é arquivada em um arquivo `log.N.log`, onde N é um número incrementado automaticamente.
	err = simulation.AddFailMatchFlagToFileLog(
		"bug:",
		"../chaosTest/bug",
	)
	if err != nil {
		util.TraceToLog()
		log.Printf("Error: %v", err.Error())
		return
	}

	// Adiciona um indicador de falha com gravação de arquivo em log ao projeto.
	// Indicador de falha é um texto procurado na saída padrão do container e indica algo que não deveria ter acontecido durante o teste.
	// Algumas falhas críticas podem ser monitoradas e quando elas acontecem, a saída padrão do container é arquivada em um arquivo `log.N.log`, onde N é um número incrementado automaticamente.
	err = simulation.AddFailMatchFlagToFileLog(
		"panic:",
		"../chaosTest/bug",
	)
	if err != nil {
		util.TraceToLog()
		log.Printf("Error: %v", err.Error())
		return
	}

	// English: Adds a filter to look for a value in the container's standard output releasing the chaos test to be started
	//
	// Português: Adiciona um filtro para procurar um valor na saída padrão do container liberando o início do teste de caos
	simulation.AddFilterToStartChaos(
		"chaos enable",
		"chaos enable",
		"",
		"",
	)

	// English: Defines the probability of the container restarting and changing the IP address in the process.
	//
	// Português: Define a probalidade do container reiniciar e mudar o endereço IP no processo.
	simulation.SetRestartProbability(1.0, 1.0, 1)

	// English: Defines a time window used to start chaos testing after container initialized
	//
	// Português: Define uma janela de tempo usada para começar o teste de caos depois do container inicializado
	simulation.SetTimeToStartChaosOnChaosScene(2*time.Second, 5*time.Second)

	// English: Sets a time window used to release container restart after the container has been initialized
	//
	// Português: Define uma janela de tempo usada para liberar o reinício do container depois do container ter sido inicializado
	simulation.SetTimeBeforeStartChaosInThisContainerOnChaosScene(2*time.Second, 5*time.Second)

	// English: Defines a time window used to pause the container
	//
	// Português: Define uma janela de tempo usada para pausar o container
	simulation.SetTimeOnContainerPausedStateOnChaosScene(2*time.Second, 5*time.Second)

	// English: Defines a time window used to unpause the container
	//
	// Português: Define uma janela de tempo usada para remover a pausa do container
	simulation.SetTimeOnContainerUnpausedStateOnChaosScene(2*time.Second, 5*time.Second)

	// English: Sets a time window used to restart the container after stopping
	//
	// Português: Define uma janela de tempo usada para reiniciar o container depois de parado
	simulation.SetTimeToRestartThisContainerAfterStopEventOnChaosScene(2*time.Second, 5*time.Second)

	// English: Enable chaos test
	//
	// Português: Habilita o teste de caos
	simulation.EnableChaosScene(enableChaos)

	// Inicializa o objeto dockerBuilder.ContainerBuilder
	err = simulation.Init()
	if err != nil {
		util.TraceToLog()
		log.Printf("Error: %v", err.Error())
		return
	}
	// Gera a imagem baseada no conteúdo da pasta
	_, err = simulation.ImageBuildFromFolder()
	if err != nil {
		util.TraceToLog()
		log.Printf("Error: %v", err.Error())
		return
	}
	// Gera o container
	err = simulation.ContainerBuildWithoutStartingItFromImage()
	if err != nil {
		util.TraceToLog()
		log.Printf("Error: %v", err.Error())
		return
	}

	log.Printf("Instalação do container de simulação, índice %v: fim", indexAsString)

	return
}
