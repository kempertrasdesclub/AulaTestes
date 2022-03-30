package localDevOps

import (
	"os"
	"path"
	"strings"
	"testing"
)

func TestCreateFolderMemoryContainer(t *testing.T) {
	var err error
	var dirPath string
	dirPath, err = createFolderMemoryContainer("chaosTest.folder", 0)
	if err != nil {
		t.FailNow()
	}

	_ = os.Remove(dirPath)

	dirPath, err = createFolderMemoryContainer("chaosTest.folder", 1)
	if err != nil {
		t.FailNow()
	}

	_ = os.Remove(dirPath)

	dirPath, err = createFolderMemoryContainer("chaosTest.folder", 2)
	if err != nil {
		t.FailNow()
	}

	_ = os.Remove(dirPath)

	var remove string
	remove = path.Base(dirPath)
	remove = strings.Replace(dirPath, remove, "", 1)
	remove = path.Base(remove)

	_ = os.Remove(remove)
}
