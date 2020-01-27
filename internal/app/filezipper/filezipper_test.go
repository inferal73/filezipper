package filezipper_test

import (
	"bytes"
	"fmt"
	"github.com/Flaque/filet"
	"github.com/inferal73/filezipper/internal/app/filezipper"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestZipFile(t *testing.T) {
	defer filet.CleanUp(t)
	root := filet.TmpDir(t, "")
	entry := filet.TmpDir(t, root)
	out := filet.TmpDir(t, root)
	outBuf := new(bytes.Buffer)

	file := filet.TmpFile(t, entry, "some content")
	info, _ := file.Stat()
	outPath, err := filezipper.ZipFile(file.Name(), out, outBuf)

	// test func return
	actualPath := filepath.Join(out, info.Name() + ".zip")
	assert.Equal(t, outPath, actualPath)
	assert.NoError(t, err)

	// test out files in dir
	files, _ := ioutil.ReadDir(out)
	assert.Len(t, files, 1)
	outFile := files[0]
	assert.Equal(t, outFile.Name(), info.Name() + ".zip")

	// test out file content
	f, _ := os.Stat(outPath)
	assert.Equal(t, f.Size(), int64(176))

	// test out buffer
	s := fmt.Sprintf("Archiving %s.zip ...\n", info.Name())
	assert.Equal(t, outBuf.String(), s)
}

func TestZipFiles(t *testing.T) {
	defer filet.CleanUp(t)
	root := filet.TmpDir(t, "")
	entry := filet.TmpDir(t, root)
	out := filet.TmpDir(t, root)
	outBuf := new(bytes.Buffer)
	fileNames := make([]afero.File, 5)

	for i := 0; i < 5; i++ {
		fileNames[i] = filet.TmpFile(t, entry, "some content")
	}
	config := filezipper.NewConfig(entry, out)
	err := filezipper.ZipFiles(config, outBuf)
	assert.NoError(t, err)

	// test out files in dir
	files, _ := ioutil.ReadDir(out)
	assert.Len(t, files, 5)

	// test out buffer
	assert.NotEmpty(t, outBuf)
}