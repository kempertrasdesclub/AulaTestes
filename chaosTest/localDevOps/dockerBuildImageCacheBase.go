package localDevOps

import (
	dockerBuilder "github.com/helmutkemper/iotmaker.docker.builder"
	"github.com/helmutkemper/util"
	"log"
	"path"
)

func dockerBuildImageCacheBase() (err error) {

	log.Print("verificando a necessidade de criação da imagem cache:latest: início")
	defer log.Print("verificando a necessidade de criação da imagem cache:latest: fim")

	// Este é o nome padrão da imagem cache usada pelo projeto docker.build
	// Embora esta cache seja grande, ela é usada apenas na primeira etapa do build e não
	// interfere no tamanho total do container final. Ela apenas contém os elementos para
	// o processo de build.
	var imageCacheName = "cachebase:latest"
	var imageId string
	var cachePath string
	var container = &dockerBuilder.ContainerBuilder{}

	cachePath, err = util.FileFindInTree("cachebase")
	cachePath = path.Dir(cachePath)

	// Caso a imagem exista, ignora o resto do código
	imageId, err = container.ImageFindIdByName(imageCacheName)
	if err != nil && err.Error() != "image name not found" {
		return
	}

	if imageId != "" {
		return
	}

	// Define o nome da imagem
	container.SetImageName(imageCacheName)
	// Imprime a saída padrão do container na saída padrão do golang
	container.SetPrintBuildOnStrOut()
	// fixme: desnecessário?
	//container.SetContainerName(imageCacheName)
	// Caminho relativo da pasta usada na criação da imagem
	container.SetBuildFolderPath(cachePath)
	// Inicializa o objeto dockerBuilder.ContainerBuilder
	err = container.Init()
	if err != nil {
		return
	}
	// Monta a imagem baseada no conteúdo da pasta
	_, err = container.ImageBuildFromFolder()
	if err != nil {
		return
	}

	return
}
