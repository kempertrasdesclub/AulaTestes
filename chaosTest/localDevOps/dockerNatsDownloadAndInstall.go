package localDevOps

import (
	dockerBuilder "github.com/helmutkemper/iotmaker.docker.builder"
	dockerBuilderNetwork "github.com/helmutkemper/iotmaker.docker.builder.network"
	"github.com/helmutkemper/util"
	"log"
	"time"
)

// dockerNatsDownloadAndInstall (português): Instala o container com o nats.io
//   Entrada:
//     netDocker: Ponteiro para o objeto gerenciador de rede docker
//   Saída:
//     err: Objeto padrão de erro
func dockerNatsDownloadAndInstall(
	netDocker *dockerBuilderNetwork.ContainerBuilderNetwork,
) (
	err error,
) {

	log.Print("instalação do container nats: início")

	// Prepara a instalação do container para a imagem nats:latest
	var natsDocker = dockerBuilder.ContainerBuilder{}
	natsDocker.SetPrintBuildOnStrOut()
	// Aponta o gerenciador de rede [opcional]
	// Como o gateway é 10.0.0.1, o primeiro container gerado fica no endereço 10.0.0.2
	natsDocker.SetNetworkDocker(netDocker)
	// Determina o nome da imagem a ser usada
	natsDocker.SetImageName("nats:latest")
	// Determina o nome do container a ser criado
	natsDocker.SetContainerName("nats_delete_after_test")

	//natsDocker.SetImageBuildOptionsMemory(4 * dockerBuilder.KMegaByte)

	// Você pode expor a porta 4222 e 6222 para o fora da rede
	natsDocker.AddPortToExpose("4222")
	natsDocker.AddPortToExpose("6222")

	// Você pode trocar uma porta 4222 para 4200 e a expor para fora da rede
	//natsDocker.AddPortToChange("4222", "4200")

	// Espera pelo texto abaixo no log do container antes de prosseguir
	natsDocker.SetWaitStringWithTimeout(
		"Listening for route connections on 0.0.0.0:6222",
		40*time.Second,
	)

	// Inicializa o objeto depois de todas as configurações feitas
	err = natsDocker.Init()
	if err != nil {
		util.TraceToLog()
		log.Printf("Error: %v", err.Error())
		return
	}

	// Baixa a imagem caso a mesma não exista e deve ser usado apenas em caso de imagens públicas
	err = natsDocker.ImagePull()
	if err != nil {
		util.TraceToLog()
		log.Printf("Error: %v", err.Error())
		return
	}

	// Monta o container a partir da imagem baixada por ImagePull() e definida em SetImageName()
	err = natsDocker.ContainerBuildAndStartFromImage()
	if err != nil {
		util.TraceToLog()
		log.Printf("Error: %v", err.Error())
		return
	}

	time.Sleep(5 * time.Second)

	log.Print("instalação do container nats: fim")
	return
}
