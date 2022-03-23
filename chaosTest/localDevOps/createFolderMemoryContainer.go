package localDevOps

import (
	"errors"
	"github.com/helmutkemper/util"
	"io/fs"
	"os"
	"path"
	"strconv"
)

// createFolderMemoryContainer
//
// Cria um diretório "/memory/container_"+index no diretório onde está contido o arquivo de
// referência.
//
//   Entrada:
//     reference: nome do arquivo de referência a ser procurado na árvore de diretórios;
//     index: número do container.
//   Saída:
//     dirPath: caminho relativo do diretório criado;
//     err: objeto de erro padrão do go.
func createFolderMemoryContainer(reference string, index int64) (dirPath string, err error) {
	var indexAsString = strconv.FormatInt(index, 10)
	dirPath, err = util.FileFindInTree(reference)
	if err != nil {
		return
	}

	dirPath = path.Dir(dirPath)
	_ = os.Mkdir(dirPath+"/memory", fs.ModePerm)
	_ = os.Mkdir(dirPath+"/memory/container_"+indexAsString, fs.ModePerm)

	var info fs.FileInfo
	info, err = os.Stat(dirPath + "/memory/container_" + indexAsString)
	if err != nil {
		return
	}

	dirPath = dirPath + "/memory/container_" + indexAsString

	if info.IsDir() == false {
		err = errors.New("create container dir error")
	}

	return
}
