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
) (
	err error,
) {

	var indexAsString = strconv.FormatInt(index, 10)

	log.Printf("Instalação do container de simulação, índice %v: início", indexAsString)

	simulation.SetImageCacheName("cachechaostest:latest")
	simulation.SetImageExpirationTime(30 * time.Minute) //fixme: definir no arquivo de teste
	// Imprime a saída padrão do container na saída padrão do golang
	simulation.SetPrintBuildOnStrOut()
	// Habilita o uso da imagem cache:latest
	simulation.SetCacheEnable(true)
	// Aponta o gerenciador de rede
	simulation.SetNetworkDocker(netDocker)
	// Determina o nome da imagem a ser usada
	simulation.SetImageName("memcache_delete_after_test:latest")
	// Define o nome do container
	simulation.SetContainerName("memcache_container_delete_after_test_" + indexAsString)
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

	err = simulation.AddFileOrFolderToLinkBetweenConputerHostAndContainer(memoryInputPath, "/memory_container")
	if err != nil {
		util.TraceToLog()
		log.Printf("Error: %v", err.Error())
		return
	}

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

	err = simulation.ContainerStartAfterBuild() //fixme: apagar esta linha
	if err != nil {
		util.TraceToLog()
		log.Printf("Error: %v", err.Error())
		return
	}

	log.Printf("Instalação do container de simulação, índice %v: fim", indexAsString)

	return
}
