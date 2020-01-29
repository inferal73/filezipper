package filezipper

import (
	"bytes"
	"fmt"
	"github.com/Flaque/filet"
	"github.com/inferal73/filezipper/internal/app/logger"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestZipFile(t *testing.T) {
	defer filet.CleanUp(t)
	root := filet.TmpDir(t, "")
	entry := filet.TmpDir(t, root)
	out := filet.TmpDir(t, root)

	outBuf := new(bytes.Buffer)
	l := logger.GetLogger()
	l.SetWriter(outBuf)

	file := filet.TmpFile(t, entry, "some content")
	info, _ := file.Stat()
	err := zipFile(file.Name(), out)

	// test func return
	assert.NoError(t, err)

	// test out files in dir
	files, _ := ioutil.ReadDir(out)
	assert.Len(t, files, 1)
	outFile := files[0]
	assert.Equal(t, outFile.Name(), info.Name() + ".zip")

	// test out buffer
	s := fmt.Sprintf("Archiving %s.zip ...\n", info.Name())
	assert.Equal(t, outBuf.String(), s)
}

func TestZip(t *testing.T) {
	defer filet.CleanUp(t)
	root := filet.TmpDir(t, "")
	entry := filet.TmpDir(t, root)
	out := filet.TmpDir(t, root)

	outBuf := new(bytes.Buffer)
	l := logger.GetLogger()
	l.SetWriter(outBuf)

	fileNames := make([]afero.File, 5)

	for i := 0; i < 5; i++ {
		fileNames[i] = filet.TmpFile(t, entry, "some content")
	}
	config := NewConfig(entry, out)
	err := Zip(config)
	assert.NoError(t, err)

	// test out files in dir
	files, _ := ioutil.ReadDir(out)
	assert.Len(t, files, 5)

	// test out buffer
	assert.NotEmpty(t, outBuf)
}