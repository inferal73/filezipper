package filezipper

import (
	"archive/zip"
	"fmt"
	"github.com/inferal73/filezipper/internal/app/logger"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Zip(config *Config) (err error) {
	logger.Log("Start processing\n")

	defer func() {
		if err == nil {
			logger.Log("Finished\n")
		}
	}()

	info, err := os.Stat(config.Entry)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return zipDir(config.Entry, config.Out)
	}
	return zipFile(config.Entry, config.Out)
}

func zipDir(entry, out string) error {
	return filepath.Walk(entry, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		return zipFile(path, out)
	})
}

func zipFile(entry, out string) (err error) {
	// get entry file info
	entryFile, err := os.Stat(entry)
	if err != nil {
		return fmt.Errorf("get info entry file error: %s", err)
	}

	zipFile, err := generateFile(out, entryFile.Name())
	if err != nil {
		return fmt.Errorf("generate out file error: %s", err)
	}
	defer func() {
		err := zipFile.Close()
		logger.Log("out file closing incorrectly: %s", err)
	}()

	// init zip file writer
	archive := zip.NewWriter(zipFile)
	defer func() {
		err := archive.Close()
		logger.Log("archive file closing incorrectly: %s", err)
	}()

	// create zip file headers
	header, err := zip.FileInfoHeader(entryFile)
	if err != nil {
		return fmt.Errorf("create zip file headers error: %s", err)
	}
	header.Method = zip.Deflate

	// open entry file
	zFile, err := os.Open(entry)
	if err != nil {
		return fmt.Errorf("create zip file headers error: %s", err)
	}
	defer func() {
		err = zFile.Close()
		logger.Log("open entry file error: %s", err)
	}()
	writer, err := archive.CreateHeader(header)
	if err != nil {
		return err
	}

	// compress file
	_, err = io.Copy(writer, zFile)
	return err
}

func generateFile(out, fileName string) (*os.File, error) {
	zipFileName := changeFileExt(fileName, "zip")
	logger.Log("Archiving %s ...\n", zipFileName)

	zipFile, _, err := createFile(out, zipFileName)
	if err != nil {
		return nil, err
	}
	return zipFile, nil
}

func changeFileExt(fileName string, ext string) string {
	fileExt := filepath.Ext(fileName)
	withOutExt := strings.TrimSuffix(fileName, fileExt)
	return fmt.Sprintf("%s.%s", withOutExt, ext)
}

func createFile(path string, fileName string) (file *os.File, pathWithFile string, err error) {
	_, err = os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
		err = os.MkdirAll(path, 0750)
		if err != nil {
			return
		}
	}
	pathWithFile = filepath.Join(path, fileName)
	file, err = os.Create(pathWithFile)
	return
}